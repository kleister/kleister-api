FROM alpine:edge
MAINTAINER Thomas Boerger <thomas@webhippie.de>

EXPOSE 8080 80 443
VOLUME ["/var/lib/kleister"]

RUN apk update && \
  apk add \
    ca-certificates \
    bash \
    sqlite && \
  rm -rf \
    /var/cache/apk/* && \
  addgroup \
    -g 1000 \
    kleister && \
  adduser -D \
    -h /var/lib/kleister \
    -s /bin/bash \
    -G kleister \
    -u 1000 \
    kleister

COPY kleister-api /usr/bin/

ENV KLEISTER_SERVER_STORAGE /var/lib/kleister

USER kleister
ENTRYPOINT ["/usr/bin/kleister-api"]
CMD ["server"]
