import { randomBytes } from 'node:crypto';
import { readFileSync, writeFileSync } from 'node:fs';

const mode = process.argv[2];
const options = JSON.parse(readFileSync('/app/options.json', 'utf8'));
const hostname = options.panel_hostname ?? 'homeassistant.local';
const port = options.panel_port ?? 8443;
if (typeof hostname !== 'string'
    || hostname.length > 253
    || !hostname.split('.').every((part) => /^[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?$/i.test(part))) {
  throw new Error('panel_hostname muss der DNS-Name oder die IPv4-Adresse deines Home Assistant sein.');
}
if (!Number.isInteger(port) || port < 1 || port > 65535) {
  throw new Error('panel_port muss zwischen 1 und 65535 liegen.');
}

if (mode === 'host') {
  process.stdout.write(hostname);
} else if (mode === 'port') {
  process.stdout.write(String(port));
} else if (mode === 'secret') {
  const key = '/app/data/session.key';
  try {
    writeFileSync(key, randomBytes(32).toString('hex'), { mode: 0o600, flag: 'wx' });
  } catch (error) {
    if (error.code !== 'EEXIST') throw error;
  }
  const secret = readFileSync(key, 'utf8').trim();
  if (!/^[a-f0-9]{64}$/.test(secret)) throw new Error('Der gespeicherte Sitzungsschlüssel ist ungültig.');
  process.stdout.write(secret);
} else if (mode === 'landing') {
  const url = `https://${hostname}:${port}/`;
  writeFileSync('/app/data/haos-web/index.html', `<!doctype html>
<html lang="de"><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>MCPanelHAOS</title>
<style>body{font:16px system-ui;margin:2rem;max-width:44rem;line-height:1.6;color-scheme:light dark}</style>
<h1>MCPanelHAOS</h1>
<p>Panel, Minecraft und Welten laufen auf diesem HAOS-Gerät.</p>
<p><a href="${url}" target="_blank" rel="noopener noreferrer">Minecraft-Verwaltung öffnen</a></p>
<p>Beim ersten Aufruf kannst du das lokale <a href="./ca.crt" download="MCPanelHAOS-CA.crt">HTTPS-Stammzertifikat</a> über Home Assistant herunterladen und auf dem Verwaltungsgerät als vertrauenswürdig hinterlegen.</p></html>`);
}
