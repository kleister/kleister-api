FROM --platform=$BUILDPLATFORM golang:1.26.2-alpine@sha256:27f829349da645e287cb195a9921c106fc224eeebbdc33aeb0f4fca2382befa6 AS builder

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

FROM alpine:3.23@sha256:5b10f432ef3da1b8d4c7eb6c487f2f5a8f096bc91145e68878dd4a5019afde11

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
