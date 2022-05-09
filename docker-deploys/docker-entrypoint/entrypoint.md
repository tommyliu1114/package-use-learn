# entrypoint 指令：
- 当指定了 ENTRYPOINT 后，CMD 的含义就发生了改变，不再是直接的运行其命令，而是将 CMD 的内容作为参数传给 ENTRYPOINT 指令，换句话说实际执行时，将变为：     <ENTRYPOINT> "<CMD>"

# entrypoint 使用场景举例：

- 让镜像变成像命令一样使用：
```bash
# dockerfile：
    FROM ubuntu:18.04
    RUN apt-get update \
        && apt-get install -y curl \
        && rm -rf /var/lib/apt/lists/*
    ENTRYPOINT [ "curl", "-s", "http://myip.ipip.net" ]


 # 下面执行bash命令：
# 生成镜像：  docker build -t myip .
# 执行镜像+参数：  docker run myip -i

```

- 应用运行前准备： 启动容器就是启动主进程，但有些时候，启动主进程前，需要一些准备工作。比如 mysql 类的数据库，可能需要一些数据库配置、初始化的工作，这些工作要在最终的 mysql 服务器运行之前解决。此外，可能希望避免使用 root 用户去启动服务，从而提高安全性，而在启动服务前还需要以 root 身份执行一些必要的准备工作，最后切换到服务用户身份启动服务。或者除了服务外，其它命令依旧可以使用 root 身份执行，方便调试等。这些准备工作是和容器 CMD 无关的，无论 CMD 为什么，都需要事先进行一个预处理的工作。这种情况下，可以写一个脚本，然后放入 ENTRYPOINT 中去执行，而这个脚本会将接到的参数（也就是 <CMD>）作为命令，在脚本最后执行。

```bash

分析官方的redis镜像

```