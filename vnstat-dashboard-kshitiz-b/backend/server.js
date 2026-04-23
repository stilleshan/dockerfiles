const express = require('express');
const { exec, execSync } = require('child_process');
const path = require('path');
const app = express();

// ----- Dynamic interface allowlist -----
const ALLOWED_PREFIXES = (process.env.ALLOWED_PREFIXES || 'eth,enp,wlan,wlp,tailscale,docker')
  .split(',')
  .map(s => s.trim())
  .filter(Boolean);

const ENV_IFACES = (process.env.ALLOWED_INTERFACES || '')
  .split(',')
  .map(s => s.trim())
  .filter(Boolean);

const DENY_PATTERNS = [
  /^veth/i, /^br[-:]/i, /^lo$/i, /^tap/i, /^tun/i, /^wg/i,
];
const isDenied = name => DENY_PATTERNS.some(rx => rx.test(name));
const isAllowedByPrefix = name => ALLOWED_PREFIXES
  .some(p => name.toLowerCase().startsWith(p.toLowerCase()));

let ALLOWED_INTERFACES = [];
if (ENV_IFACES.length) {
  // explicit list via env has priority
  ALLOWED_INTERFACES = ENV_IFACES;
} else {
  try {
    const out = execSync('vnstat --iflist', { encoding: 'utf8' });
    // work with either "Available interfaces:" or "interfaces:"
    const line = out.split('\n').find(l => /interfaces:/i.test(l));
    const listPart = line ? line.replace(/^.*interfaces:\s*/i, '') : '';
    if (listPart) {
      ALLOWED_INTERFACES = listPart
        .split(/\s+/)
        .filter(tok => /^[a-z0-9:_-]+$/i.test(tok))  // drop "(100" "Mbit)" etc.
        .filter(name => !isDenied(name))
        .filter(name => isAllowedByPrefix(name));
    }
    console.log('✅ Detected interfaces:', ALLOWED_INTERFACES.join(', ') || '(none)');
  } catch (e) {
    console.warn('⚠️ Could not auto-detect vnstat interfaces:', e.message);
  }
}
if (ALLOWED_INTERFACES.length === 0) {
  ALLOWED_INTERFACES = ['eth0', 'wlan0', 'tailscale0', 'docker0'];
  console.log('ℹ️ Using fallback interfaces:', ALLOWED_INTERFACES.join(', '));
}
const ALLOWED_SET = new Set(ALLOWED_INTERFACES);

const FRONTEND_DIR = process.env.FRONTEND_DIR || 'frontend-build';
app.use(express.static(path.join(__dirname, '..', FRONTEND_DIR)));

app.get('/api/interfaces', (_req, res) => {
  res.json({ interfaces: ALLOWED_INTERFACES });
});

app.get('/api/vnstat/:iface', (req, res) => {
  const iface = req.params.iface;
  if (!ALLOWED_SET.has(iface)) {
    return res.status(400).json({ error: 'Invalid interface' });
  }
  exec(`vnstat -i ${iface} --json`, (error, stdout, stderr) => {
    if (error) {
      return res.status(500).json({ error: stderr || error.message });
    }
    try {
      const data = JSON.parse(stdout);
      res.json(data);
    } catch (e) {
      res.status(500).json({ error: 'Failed to parse vnstat output.' });
    }
  });
});

// Heath Check
app.get('/healthz', (_req, res) => res.json({ ok: true }));

// Serve React frontend for all non-API routes
app.get('*', (req, res) => {
  res.sendFile(path.join(__dirname, '..', FRONTEND_DIR, 'index.html'));
});

const PORT = process.env.PORT || 3001;
app.listen(PORT, () => console.log(`Server running on port ${PORT}`));