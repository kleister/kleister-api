FROM --platform=$BUILDPLATFORM golang:1.26.5-alpine@sha256:99e12cfb19b753915f9b9fdc5a99f1869a24a69d3a0955832d5702e7fa68f1be AS builder

RUN apk add --no-cache -U nodejs npm git curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /go/src/kleister
COPY . /go/src/kleister/

RUN --mount=type=cache,target=/go/pkg \
    go mod download -x

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    task fe:install fe:generate fe:build && \
    task be:generate build:server GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:3.24@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b

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

COPY --from=builder /go/src/kleister/bin/kleister-api /usr/bin/kleister-api
COPY --from=builder --chown=1337:1337 /go/src/kleister/pkg/frontend/files /usr/share/kleister
WORKDIR /var/lib/kleister
USER kleister
