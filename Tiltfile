load("ext://restart_process", "docker_build_with_restart")

docker_build_with_restart(
    "scaling-ws/ws-server",
    context=".",
    dockerfile="dockerfiles/ws_server.Dockerfile",
    entrypoint="/usr/local/bin/run-ws-server",
    only=["./apps/ws-server"],
    live_update=[
        sync("./apps/ws-server", "/app"),
        run("go build -mod=vendor -o /usr/local/bin/run-ws-server /app/cmd/run-server"),
    ],
)

docker_build(
    "scaling-ws/web",
    context=".",
    dockerfile="./dockerfiles/web.Dockerfile",
    only=["./apps/web"],
    ignore=["./web/dist/"],
    live_update=[
        fall_back_on('./apps/web/vite.config.ts'),
        sync('./web/', '/app/'),
        run(
            'pnpm install',
            trigger=['./apps/web/package.json', './apps/web/pnpm-lock.yaml']
        )
    ]
)

k8s_yaml(
    ["k8s/ws-server/deployment.yml", "k8s/ws-server/service.yml"],
)
k8s_resource(
    "ws-server",
    port_forwards=8080,
    labels=["ws-server"]
)



k8s_yaml(["k8s/web/deployment.yml", "k8s/web/service.yml"])
k8s_resource("web", port_forwards=["3000:3000"], labels=["web"])