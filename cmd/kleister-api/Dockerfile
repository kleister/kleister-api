FROM --platform=$BUILDPLATFORM debian:13-slim@sha256:89c1ea91e897f0fe0a3731595330bc45b8bdef4047464a07116f891a94054b31 AS builder
WORKDIR /src

RUN apt-get update && \
    apt-get -y --no-install-recommends install ca-certificates curl git build-essential && \
    rm -rf /var/lib/apt/lists/*

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
ENV PATH="/root/.local/bin:/mise/shims:$PATH" \
    MISE_DATA_DIR="/mise" \
    MISE_CONFIG_DIR="/mise" \
    MISE_CACHE_DIR="/mise/cache" \
    MISE_INSTALL_PATH="/usr/local/bin/mise"

RUN curl --proto '=https' --proto-redir '=https' \
    --fail --show-error --silent --location https://mise.run | sh

COPY mise.toml ./
RUN --mount=type=cache,target=/mise/cache \
   mise trust && mise install

COPY . /src/
RUN --mount=type=cache,target=/root/go/pkg \
    go mod download -x

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/root/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    task fe:install fe:generate fe:build && \
    task be:generate build:server GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:3.24@sha256:be2656061b536c1883c4034790e21d0a09c5aaaf51968b89368a3c0371ea2c9a

RUN apk add --no-cache ca-certificates mailcap && \
    addgroup -g 1337 kleister && \
    adduser -D -u 1337 -h /var/lib/kleister -G kleister kleister

EXPOSE 8000 8080
VOLUME ["/var/lib/kleister"]
ENTRYPOINT ["/usr/bin/kleister-api"]
CMD ["server"]
HEALTHCHECK CMD ["/usr/bin/kleister-api", "health"]

ENV KLEISTER_API_DATABASE_DRIVER="sqlite3"
ENV KLEISTER_API_DATABASE_NAME="/var/lib/kleister/database.sqlite3"
ENV KLEISTER_API_UPLOAD_DRIVER="file"
ENV KLEISTER_API_UPLOAD_PATH="/var/lib/kleister/uploads"

COPY --from=builder /src/bin/kleister-api /usr/bin/kleister-api
COPY --from=builder --chown=1337:1337 /src/pkg/frontend/files /usr/share/kleister
WORKDIR /var/lib/kleister
USER kleister
