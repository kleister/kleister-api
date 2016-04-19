FROM alpine:edge

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf \
    /var/cache/apk/*

ADD bin/solder-api /usr/bin/
ENTRYPOINT ["/usr/bin/solder-api"]
CMD ["server"]
