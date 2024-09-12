FROM golang:1.22-alpine

WORKDIR /app

ENV GOMODCACHE=/cache/gomod
ENV GOCACHE=/cache/gobuild

COPY ./apps/ws-server/go.mod ./
RUN --mount=type=cache,target=/cache/gomod \
    go mod download
