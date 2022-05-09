# 启动docker容器的两种方式：
- 基于镜像新建一个容器并启动
```bash
    docker run -t -i ubuntu:18.04 /bin/bash

```

- 将在终止状态（exited）的容器重新启动
```bash
    docker container start
```

# 守护态运行容器：-d 
- docker run -d ubuntu:18.04 /bin/sh -c "while true; do echo hello world; sleep 1; done"

# docker container ls 查看容器信息

# docker container logs 来查看容器输出内容

# docker container stop 来终止运行中的容器 

# 进入容器：
- attach 命令：
- exec 命令：
    -i -t 参数

# 导出容器：
- docker export 7691a814370e > ubuntu.tar
```bash
    $ docker container ls -a
    CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                    PORTS               NAMES
    7691a814370e        ubuntu:18.04        "/bin/bash"         36 hours ago        Exited (0) 21 hours ago                       test
    $ docker export 7691a814370e > ubuntu.tar

```

# 导入容器快照：可以使用 docker import 从容器快照文件中再导入为镜像

```bash
    $ cat ubuntu.tar | docker import - test/ubuntu:v1.0
    $ docker image ls
    REPOSITORY          TAG                 IMAGE ID            CREATED              VIRTUAL SIZE
    test/ubuntu         v1.0                9d37a6082e97        About a minute ago   171.3 MB

```

# 删除容器：
- docker container rm nginx:v3-mms17 

# 清理所有处于终止状态的容器：
- docker container prune 