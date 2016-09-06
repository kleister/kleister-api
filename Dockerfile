FROM alpine:edge

RUN apk update && \
  apk add \
    ca-certificates \
    sqlite && \
  rm -rf \
    /var/cache/apk/*

ADD bin/kleister-api /usr/bin/
ENTRYPOINT ["/usr/bin/kleister-api"]
CMD ["server"]
