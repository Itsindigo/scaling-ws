FROM golang:1.22-alpine

WORKDIR /app

ENV GOMODCACHE=/cache/gomod
ENV GOCACHE=/cache/gobuild
COPY ./apps/ws-server/go.mod ./apps/ws-server/go.sum ./

RUN --mount=type=cache,target=/cache/gomod \
    go mod download

COPY apps/ws-server /app/

RUN --mount=type=cache,target=/cache/gomod \
    --mount=type=cache,target=/cache/gobuild,sharing=locked \
    go mod vendor && \
    go build -mod=vendor -o /usr/local/bin/run-ws-server /app/cmd/run-server

CMD ["/usr/local/bin/run-ws-server"]
