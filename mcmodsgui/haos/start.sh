#!/bin/sh
set -eu

if [ ! -S /run/docker.sock ]; then
  echo "Docker-Socket fehlt. Deaktiviere den Schutzmodus nur für MCPanelHAOS und starte das Add-on neu." >&2
  exit 1
fi
if ! docker info >/dev/null; then
  echo "Die Docker-Engine ist aus diesem Add-on nicht erreichbar." >&2
  exit 1
fi

mkdir -p /app/servers /app/data /app/data/haos-web
PANEL_HOST="$(node /opt/minepanel/haos/options.mjs host)"
PANEL_PORT="$(node /opt/minepanel/haos/options.mjs port)"
JWT_SECRET="$(node /opt/minepanel/haos/options.mjs secret)"
FRONTEND_URL="https://${PANEL_HOST}:${PANEL_PORT}"
export PANEL_HOST PANEL_PORT JWT_SECRET FRONTEND_URL
node /opt/minepanel/haos/options.mjs landing
exec /usr/bin/supervisord -c /opt/minepanel/haos/supervisord.conf
