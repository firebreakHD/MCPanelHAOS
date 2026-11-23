ARG APP_VERSION=2.0.15

# Generate protobuf code using buf
FROM bufbuild/buf:1.57.2 AS proto-builder
WORKDIR /app
COPY buf.yaml buf.gen.yaml buf.lock ./
COPY proto/ ./proto/
RUN mkdir -p web/discopanel/static pkg/proto web/discopanel/src/lib/proto && buf generate

FROM node:22-alpine AS frontend-builder

ARG APP_VERSION
ENV APP_VERSION=${APP_VERSION}

WORKDIR /app/web/discopanel

COPY web/discopanel/package*.json ./
RUN npm ci

COPY web/discopanel/ ./
COPY --from=proto-builder /app/web/discopanel/src/lib/proto ./src/lib/proto/
COPY --from=proto-builder /app/web/discopanel/static/schemav1.yaml ./static/

RUN npm run build

FROM golang:1.25-alpine AS backend-builder

ARG APP_VERSION
ENV APP_VERSION=${APP_VERSION}

RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY --from=proto-builder /app/pkg/proto ./pkg/proto/
COPY --from=frontend-builder /app/web/discopanel/build ./web/discopanel/build

RUN CGO_ENABLED=1 GOOS=linux go build -o discopanel ./cmd/discopanel

FROM alpine:3.23
ARG APP_VERSION=2.0.15
ARG BUILD_ARCH=amd64
ARG BUILD_VERSION
LABEL io.hass.type="app" io.hass.arch="$BUILD_ARCH" io.hass.version="$BUILD_VERSION"
ENV APP_VERSION=$APP_VERSION MCPANEL_HAOS=true
RUN apk add --no-cache ca-certificates sqlite-libs caddy docker-cli jq supervisor
WORKDIR /opt/discopanel
COPY --from=backend-builder /app/discopanel ./discopanel
COPY haos/ ./haos/
RUN chmod 755 /opt/discopanel/haos/start.sh
EXPOSE 8099 8443
CMD ["/opt/discopanel/haos/start.sh"]
