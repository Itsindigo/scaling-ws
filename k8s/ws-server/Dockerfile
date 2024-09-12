FROM scaling-ws/go-base

WORKDIR /app

COPY apps/ws-server /app/

RUN --mount=type=cache,target=/cache/gomod \
    --mount=type=cache,target=/cache/gobuild,sharing=locked \
    go mod vendor && \
    go build -mod=vendor -o /usr/local/bin/run-ws-server /app

CMD ["/usr/local/bin/run-ws-server"]
