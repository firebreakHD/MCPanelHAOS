# Fullstack image (backend + frontend under supervisord). Build context is the repo root.
FROM node:24-alpine AS base

ENV HUSKY=0 CI=true NEXT_TELEMETRY_DISABLED=1
RUN corepack enable

WORKDIR /app

FROM base AS builder

COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./
COPY backend/package.json backend/
COPY frontend/package.json frontend/
RUN --mount=type=cache,target=/root/.local/share/pnpm/store \
    pnpm install --frozen-lockfile --filter ./backend --filter ./frontend

COPY backend/ backend/
COPY frontend/ frontend/

RUN pnpm --filter ./backend build && pnpm --filter ./frontend build:turbopack
RUN pnpm --filter ./backend --prod --legacy deploy /out/backend

FROM base AS runner

RUN apk add --no-cache caddy docker-cli docker-cli-compose supervisor
ARG BUILD_ARCH=amd64
ARG BUILD_VERSION
LABEL io.hass.type="app" io.hass.arch="$BUILD_ARCH" io.hass.version="$BUILD_VERSION"
ENV APP_VERSION=$BUILD_VERSION \
    NODE_ENV=production \
    DOCKER_HOST=unix:///run/docker.sock \
    MINEPANEL_HAOS=true \
    MINEPANEL_EXCLUSIVE_JAVA=true \
    COMPOSE_PROJECT=mcmodsgui-haos \
    NEXT_PUBLIC_BACKEND_URL=/backend \
    NEXT_PUBLIC_DEFAULT_LANGUAGE=de \
    XDG_DATA_HOME=/app/data/caddy \
    XDG_CONFIG_HOME=/app/data/caddy-config

RUN addgroup --system --gid 1001 nodejs && adduser --system --uid 1001 nextjs
WORKDIR /opt/minepanel
COPY --from=builder /out/backend ./backend
COPY --from=builder --chown=nextjs:nodejs /app/frontend/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/frontend/.next/static ./frontend/.next/static
COPY --from=builder --chown=nextjs:nodejs /app/frontend/public ./frontend/public
COPY haos/ ./haos/
RUN chmod 755 /opt/minepanel/haos/start.sh
EXPOSE 8099 8443
CMD ["/opt/minepanel/haos/start.sh"]
