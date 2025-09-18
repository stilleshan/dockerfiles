const express = require('express');
const { exec } = require('child_process');
const path = require('path');
const app = express();

const ALLOWED_INTERFACES = ['eth0', 'wlan0', 'docker0', 'tailscale0'];

const FRONTEND_DIR = process.env.FRONTEND_DIR || 'frontend-build';
app.use(express.static(path.join(__dirname, '..', FRONTEND_DIR)));


app.get('/api/vnstat/:iface', (req, res) => {
  const iface = req.params.iface;
  if (!ALLOWED_INTERFACES.includes(iface)) {
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

// Serve React frontend for all non-API routes
app.get('*', (req, res) => {
  res.sendFile(path.join(__dirname, '..', FRONTEND_DIR, 'index.html'));
});

const PORT = process.env.PORT || 3001;
app.listen(PORT, () => console.log(`Server running on port ${PORT}`));
