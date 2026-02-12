package server

const dashboardHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Wallet</title>
<style>
  *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
  body {
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
    background: #0f1117;
    color: #e4e4e7;
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  header {
    width: 100%;
    padding: 1.5rem 2rem;
    background: #16181d;
    border-bottom: 1px solid #27272a;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  header h1 { font-size: 1.25rem; font-weight: 600; }
  .header-right { display: flex; align-items: center; gap: 1rem; }
  .header-right .version { color: #71717a; font-size: 0.875rem; }
  main {
    width: 100%;
    max-width: 72rem;
    padding: 2rem;
    flex: 1;
  }

  /* Wallet identity bar */
  .wallet-bar {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-bottom: 2rem;
    padding: 1rem 1.25rem;
    background: #16181d;
    border: 1px solid #27272a;
    border-radius: 0.5rem;
  }
  .wallet-bar .label { color: #71717a; font-size: 0.875rem; white-space: nowrap; }
  .wallet-bar .address {
    font-family: monospace;
    font-size: 0.875rem;
    color: #a1a1aa;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    flex: 1;
  }
  .wallet-bar .no-wallet { color: #71717a; font-size: 0.875rem; font-style: italic; }

  /* Endpoint cards grid */
  .endpoints {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(20rem, 1fr));
    gap: 1rem;
  }
  .ep-card {
    background: #16181d;
    border: 1px solid #27272a;
    border-radius: 0.5rem;
    overflow: hidden;
    display: flex;
    flex-direction: column;
  }
  .ep-card-header {
    padding: 1rem 1.25rem;
    border-bottom: 1px solid #1e1e22;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  .ep-card-header h3 { font-size: 1rem; font-weight: 600; }
  .ep-card-body {
    padding: 1rem 1.25rem;
    display: flex;
    flex-direction: column;
    gap: 0.625rem;
    flex: 1;
  }
  .ep-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 0.8125rem;
  }
  .ep-row .label { color: #71717a; }
  .ep-row .value { font-family: monospace; font-size: 0.8rem; color: #a1a1aa; }
  .ep-row .value.balance { color: #e4e4e7; font-weight: 600; font-size: 0.9rem; }

  /* Status dot */
  .status-dot {
    display: inline-block;
    width: 8px;
    height: 8px;
    border-radius: 50%;
    margin-right: 0.375rem;
  }
  .status-online .status-dot { background: #4ade80; }
  .status-offline .status-dot { background: #f87171; }
  .status-checking .status-dot { background: #facc15; animation: pulse 1.5s infinite; }
  @keyframes pulse { 0%, 100% { opacity: 1; } 50% { opacity: 0.4; } }

  .status-text { font-size: 0.75rem; }
  .status-online .status-text { color: #4ade80; }
  .status-offline .status-text { color: #f87171; }
  .status-checking .status-text { color: #facc15; }

  /* URL display */
  .url-display {
    font-family: monospace;
    font-size: 0.75rem;
    color: #52525b;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    max-width: 14rem;
  }

  /* Buttons */
  .btn {
    padding: 0.35rem 0.75rem;
    border: 1px solid #27272a;
    border-radius: 0.25rem;
    background: #27272a;
    color: #e4e4e7;
    font-size: 0.75rem;
    cursor: pointer;
    transition: background 0.15s;
  }
  .btn:hover { background: #3f3f46; }
  .btn-primary {
    background: #1d4ed8;
    border-color: #1d4ed8;
  }
  .btn-primary:hover { background: #2563eb; }

  /* Empty state */
  .empty-state {
    text-align: center;
    padding: 4rem 2rem;
    color: #71717a;
  }
  .empty-state h2 { font-size: 1.125rem; margin-bottom: 0.5rem; color: #a1a1aa; }
  .empty-state p { font-size: 0.875rem; margin-bottom: 1.5rem; }

  /* Modal */
  .modal-overlay {
    display: none;
    position: fixed;
    inset: 0;
    background: rgba(0,0,0,0.6);
    z-index: 100;
    justify-content: center;
    align-items: center;
  }
  .modal-overlay.active { display: flex; }
  .modal {
    background: #16181d;
    border: 1px solid #27272a;
    border-radius: 0.5rem;
    padding: 1.5rem;
    width: 26rem;
    max-width: 90vw;
  }
  .modal h3 { margin-bottom: 1rem; font-size: 1rem; }
  .modal label {
    display: block;
    font-size: 0.8125rem;
    color: #71717a;
    margin-bottom: 0.25rem;
    margin-top: 0.75rem;
  }
  .modal label:first-of-type { margin-top: 0; }
  .modal input, .modal select {
    width: 100%;
    padding: 0.5rem 0.75rem;
    background: #0f1117;
    border: 1px solid #27272a;
    border-radius: 0.25rem;
    color: #e4e4e7;
    font-size: 0.875rem;
    font-family: inherit;
  }
  .modal input:focus, .modal select:focus {
    outline: none;
    border-color: #1d4ed8;
  }
  .modal-footer {
    display: flex;
    justify-content: flex-end;
    gap: 0.5rem;
    margin-top: 1.25rem;
  }
  .modal-error {
    color: #f87171;
    font-size: 0.8125rem;
    margin-top: 0.5rem;
    display: none;
  }

  /* Hex block number formatting */
  .mono { font-family: monospace; font-size: 0.8rem; }

  /* Latency */
  .latency { font-size: 0.75rem; color: #52525b; }
  .latency.fast { color: #4ade80; }
  .latency.medium { color: #facc15; }
  .latency.slow { color: #fb923c; }
</style>
</head>
<body>

<header>
  <h1>Wallet</h1>
  <div class="header-right">
    <span class="version">v{{VERSION}}</span>
  </div>
</header>

<main>
  <div class="wallet-bar">
    <span class="label">Address</span>
    <span id="wallet-address" class="no-wallet">No wallet loaded</span>
    <button class="btn btn-primary" id="btn-load-key" onclick="showModal('key-modal')">Load Key</button>
  </div>

  <div id="endpoints-container">
    <div class="empty-state status-checking">
      <span class="status-dot"></span>
      <span class="status-text">Loading endpoints...</span>
    </div>
  </div>
</main>

<!-- Load Private Key Modal -->
<div class="modal-overlay" id="key-modal">
  <div class="modal">
    <h3>Load Private Key</h3>
    <label for="key-input">Private Key (hex)</label>
    <input type="password" id="key-input" placeholder="0x..." autocomplete="off" spellcheck="false">
    <div class="modal-error" id="key-error"></div>
    <div class="modal-footer">
      <button class="btn" onclick="hideModal('key-modal')">Cancel</button>
      <button class="btn btn-primary" onclick="loadKey()">Load</button>
    </div>
  </div>
</div>

<script>
// ── State ──────────────────────────────────────────────
let endpoints = [];
let walletAddress = sessionStorage.getItem('wallet_address') || '';
let privateKey = sessionStorage.getItem('wallet_key') || '';

// ── Init ───────────────────────────────────────────────
if (walletAddress) {
  document.getElementById('wallet-address').className = 'address';
  document.getElementById('wallet-address').textContent = walletAddress;
  document.getElementById('btn-load-key').textContent = 'Change Key';
}

refresh();
setInterval(refresh, 10000);

// ── Refresh ────────────────────────────────────────────
async function refresh() {
  try {
    const resp = await fetch('/api/status');
    const data = await resp.json();
    endpoints = data.endpoints || [];
    renderEndpoints();
  } catch (err) {
    console.error('status poll failed:', err);
  }
}

// ── Render ─────────────────────────────────────────────
function renderEndpoints() {
  const container = document.getElementById('endpoints-container');

  if (endpoints.length === 0) {
    container.innerHTML =
      '<div class="empty-state">' +
        '<h2>No Endpoints Configured</h2>' +
        '<p>Add RPC endpoints to endpoints.json to get started.</p>' +
      '</div>';
    return;
  }

  let html = '<div class="endpoints">';
  for (const ep of endpoints) {
    const statusClass = ep.online ? 'status-online' : 'status-offline';
    const statusLabel = ep.online ? 'Online' : 'Offline';
    const chainId = ep.chain_id ? hexToDecimal(ep.chain_id) : '—';
    const blockNum = ep.block_number ? hexToDecimal(ep.block_number) : '—';
    const latencyClass = ep.latency_ms < 200 ? 'fast' : ep.latency_ms < 1000 ? 'medium' : 'slow';
    const urlAbbrev = abbreviateURL(ep.url);

    html += '<div class="ep-card">';
    html +=   '<div class="ep-card-header">';
    html +=     '<h3>' + esc(ep.name) + '</h3>';
    html +=     '<span class="' + statusClass + '">';
    html +=       '<span class="status-dot"></span>';
    html +=       '<span class="status-text">' + statusLabel + '</span>';
    html +=     '</span>';
    html +=   '</div>';
    html +=   '<div class="ep-card-body">';
    html +=     '<div class="ep-row">';
    html +=       '<span class="label">RPC</span>';
    html +=       '<span class="url-display" title="' + esc(ep.url) + '">' + esc(urlAbbrev) + '</span>';
    html +=     '</div>';
    html +=     '<div class="ep-row">';
    html +=       '<span class="label">Chain ID</span>';
    html +=       '<span class="value">' + chainId + '</span>';
    html +=     '</div>';
    html +=     '<div class="ep-row">';
    html +=       '<span class="label">Block</span>';
    html +=       '<span class="value">' + formatNumber(blockNum) + '</span>';
    html +=     '</div>';
    html +=     '<div class="ep-row">';
    html +=       '<span class="label">Latency</span>';
    html +=       '<span class="latency ' + latencyClass + '">' + ep.latency_ms + ' ms</span>';
    html +=     '</div>';

    // Balance row — only if wallet is loaded and endpoint is online
    if (walletAddress && ep.online) {
      html +=   '<div class="ep-row" id="balance-' + esc(ep.id) + '">';
      html +=     '<span class="label">Balance</span>';
      html +=     '<span class="value balance" data-ep="' + esc(ep.id) + '">...</span>';
      html +=   '</div>';
    }

    html +=   '</div>'; // ep-card-body
    html += '</div>'; // ep-card
  }
  html += '</div>';
  container.innerHTML = html;

  // Fetch balances after rendering.
  if (walletAddress) {
    fetchBalances();
  }
}

// ── Balances ───────────────────────────────────────────
async function fetchBalances() {
  for (const ep of endpoints) {
    if (!ep.online) continue;
    try {
      const resp = await fetch('/api/rpc/' + ep.id, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ method: 'eth_getBalance', params: [walletAddress, 'latest'] })
      });
      const data = await resp.json();
      if (data.result) {
        const el = document.querySelector('[data-ep="' + ep.id + '"]');
        if (el) {
          el.textContent = formatBalance(data.result) + ' ' + (ep.symbol || 'ETH');
        }
      }
    } catch (err) {
      console.error('balance fetch failed for ' + ep.name + ':', err);
    }
  }
}

// ── Key Management ─────────────────────────────────────
async function loadKey() {
  const input = document.getElementById('key-input');
  const errEl = document.getElementById('key-error');
  let key = input.value.trim();

  if (!key) {
    errEl.textContent = 'Please enter a private key.';
    errEl.style.display = 'block';
    return;
  }

  // Normalize: add 0x prefix if missing.
  if (!key.startsWith('0x')) key = '0x' + key;

  // Basic hex validation: should be 0x + 64 hex chars.
  if (!/^0x[0-9a-fA-F]{64}$/.test(key)) {
    errEl.textContent = 'Invalid key format. Expected 64 hex characters.';
    errEl.style.display = 'block';
    return;
  }

  try {
    // Derive address client-side using ethers.js (loaded lazily).
    await ensureEthers();
    const wallet = new ethers.Wallet(key);
    walletAddress = wallet.address;
    privateKey = key;

    sessionStorage.setItem('wallet_address', walletAddress);
    sessionStorage.setItem('wallet_key', privateKey);

    document.getElementById('wallet-address').className = 'address';
    document.getElementById('wallet-address').textContent = walletAddress;
    document.getElementById('btn-load-key').textContent = 'Change Key';

    input.value = '';
    errEl.style.display = 'none';
    hideModal('key-modal');
    refresh();
  } catch (err) {
    errEl.textContent = 'Failed to derive address: ' + err.message;
    errEl.style.display = 'block';
  }
}

// ── Ethers.js Lazy Load ────────────────────────────────
let ethersLoaded = false;
function ensureEthers() {
  if (ethersLoaded) return Promise.resolve();
  return new Promise((resolve, reject) => {
    const script = document.createElement('script');
    script.src = 'https://cdnjs.cloudflare.com/ajax/libs/ethers/6.13.4/ethers.umd.min.js';
    script.onload = () => { ethersLoaded = true; resolve(); };
    script.onerror = () => reject(new Error('Failed to load ethers.js'));
    document.head.appendChild(script);
  });
}

// ── Helpers ────────────────────────────────────────────
function hexToDecimal(hex) {
  if (!hex || hex === '0x') return '0';
  return parseInt(hex, 16).toString();
}

function formatNumber(n) {
  if (n === '—') return n;
  return Number(n).toLocaleString();
}

function formatBalance(hexWei) {
  // Convert hex wei to ether (18 decimals).
  const wei = BigInt(hexWei);
  const ether = Number(wei) / 1e18;
  if (ether === 0) return '0';
  if (ether < 0.0001) return '< 0.0001';
  return ether.toFixed(4);
}

function abbreviateURL(url) {
  try {
    const u = new URL(url);
    let display = u.hostname;
    if (u.port) display += ':' + u.port;
    if (u.pathname !== '/') display += u.pathname;
    return display;
  } catch {
    return url;
  }
}

function esc(s) {
  const d = document.createElement('div');
  d.textContent = s || '';
  return d.innerHTML;
}

function showModal(id) {
  document.getElementById(id).classList.add('active');
}
function hideModal(id) {
  document.getElementById(id).classList.remove('active');
}

// Close modals on overlay click.
document.querySelectorAll('.modal-overlay').forEach(overlay => {
  overlay.addEventListener('click', (e) => {
    if (e.target === overlay) overlay.classList.remove('active');
  });
});

// Close modals on Escape key.
document.addEventListener('keydown', (e) => {
  if (e.key === 'Escape') {
    document.querySelectorAll('.modal-overlay.active').forEach(m => m.classList.remove('active'));
  }
});
</script>
</body>
</html>`
