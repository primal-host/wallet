FROM golang:1.25-alpine AS build
WORKDIR /src
ENV GOPROXY=http://host.docker.internal:3000|direct
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /wallet ./cmd/wallet

FROM alpine:3.21
RUN apk add --no-cache ca-certificates
COPY --from=build /wallet /usr/local/bin/wallet
COPY endpoints.json /etc/wallet/endpoints.json
ENV ENDPOINTS_FILE=/etc/wallet/endpoints.json
ENTRYPOINT ["wallet"]
