# HEALTHCHECK 健康检查:
- HEALTHCHECK [选项] CMD <命令>：设置检查容器健康状况的命令 
- 当在一个镜像指定了 HEALTHCHECK 指令后，用其启动容器，初始状态会为 starting，在 HEALTHCHECK 指令检查成功后变为 healthy，如果连续一定次数失败，则会变为 unhealthy

- 支持选项：
    - --interval=<间隔>：两次健康检查的间隔，默认为 30 秒；
    - --timeout=<时长>：健康检查命令运行超时时间，如果超过这个时间，本次健康检查就被视为失败，默认 30 秒；
    - --retries=<次数>：当连续失败指定次数后，则将容器状态视为 unhealthy，默认 3 次。

- 在 HEALTHCHECK [选项] CMD 后面的命令，格式和 ENTRYPOINT 一样，分为 shell 格式，和 exec 格式。命令的返回值决定了该次健康检查的成功与否：0：成功；1：失败；2：保留，不要使用这个值。

- 示例： 
```bash
    FROM nginx
    RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/*
    HEALTHCHECK --interval=5s --timeout=3s \
      CMD curl -fs http://localhost/ || exit 1

```