# ARG 构建参数
- 格式： ARG <参数名>[=<默认值>]   

- ARG 指令有生效范围，如果在 FROM 指令之前指定，那么只能用于 FROM 指令中。
- 对于在各个阶段中使用的变量都必须在每个阶段分别指定

```bash
    ARG DOCKER_USERNAME=library
    FROM ${DOCKER_USERNAME}/alpine
    # 在FROM 之后使用变量，必须在每个阶段分别指定
    ARG DOCKER_USERNAME=library
    RUN set -x ; echo ${DOCKER_USERNAME}
    FROM ${DOCKER_USERNAME}/alpine
    # 在FROM 之后使用变量，必须在每个阶段分别指定
    ARG DOCKER_USERNAME=library
    RUN set -x ; echo ${DOCKER_USERNAME}

```