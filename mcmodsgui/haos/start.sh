#!/bin/sh
set -eu

export DOCKER_HOST=unix:///run/docker.sock
if [ ! -S /run/docker.sock ] || ! docker version --format '{{.Server.Version}}' >/dev/null; then
  echo 'Docker ist nicht erreichbar. Deaktiviere den Schutzmodus für MCPanelHAOS und starte das Add-on.' >&2
  exit 1
fi

# The container hostname may be a Supervisor DNS alias. Use Docker's mounts
# for the current container to resolve the private data directory on the daemon.
self="$(sed -nE 's@.*/containers/([0-9a-f]{64})/(hostname|hosts|resolv\.conf) .*@\1@p' /proc/self/mountinfo | head -n 1)"
if [ -z "$self" ]; then
  self="$(hostname)"
fi
host_app="$(docker inspect --type container "$self" | jq -er '.[0].Mounts[] | select(.Destination == "/app" and .Type == "bind" and .RW) | .Source | select(startswith("/"))')"
if [ -z "$host_app" ] || [ "$host_app" = / ]; then
  echo 'Der private Supervisor-Datenordner konnte nicht ermittelt werden.' >&2
  exit 1
fi

export DISCOPANEL_SELF_CONTAINER_ID="$self"
MCPANEL_HAOS_INSTALLATION="$(printf '%s' "$host_app" | sha256sum | cut -d ' ' -f 1)"
export MCPANEL_HAOS_INSTALLATION
export DISCOPANEL_DATA_DIR=/app/discopanel/data
export DISCOPANEL_HOST_DATA_PATH="$host_app/discopanel/data"
export XDG_DATA_HOME=/app/discopanel/caddy
export XDG_CONFIG_HOME=/app/discopanel/caddy-config
mkdir -p /app/discopanel/data /app/discopanel/backups /app/discopanel/tmp /app/discopanel/landing
chmod 700 /app/discopanel
PANEL_HOST="$(jq -er '.panel_hostname // "homeassistant.local" | select(type == "string" and length <= 253 and (split(".") | all(test("^[A-Za-z0-9]([A-Za-z0-9-]{0,61}[A-Za-z0-9])?$"))))' /app/options.json)"
PANEL_PORT="$(jq -er '.panel_port // 18443 | select(type == "number" and floor == . and . >= 1 and . <= 65535)' /app/options.json)"
export PANEL_HOST
cat >/app/discopanel/landing/index.html <<EOF
<!doctype html>
<html lang="de"><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1">
<title>MCPanelHAOS</title>
<style>body{font:16px system-ui;margin:2rem;max-width:44rem;line-height:1.6;color-scheme:light dark}</style>
<h1>MCPanelHAOS · DiscoPanel</h1>
<p><a href="https://$PANEL_HOST:$PANEL_PORT/" target="_blank" rel="noopener noreferrer">DiscoPanel öffnen</a></p>
<p>Für HTTPS kannst du das <a href="./ca.crt" download="MCPanelHAOS-CA.crt">lokale Stammzertifikat</a> herunterladen und auf deinem Verwaltungsgerät als vertrauenswürdig hinterlegen.</p>
</html>
EOF
exec supervisord -c /opt/discopanel/haos/supervisord.conf
