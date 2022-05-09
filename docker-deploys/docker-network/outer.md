# 当使用 -P 标记时，Docker 会随机映射一个端口到内部容器开放的网络端口

# -p 则可以指定要映射的端口，并且，在一个指定端口上只可以绑定一个容器

- 映射所有接口地址, hostPort:containerPort

docker run -d -p 80:80 nginx:alpine , 本地的 80 端口映射到容器的 80 端口

- 映射到指定地址的指定端口,ip:hostPort:containerPort 格式指定映射使用一个特定地址

docker run -d -p 127.0.0.1:80:80 nginx:alpine 

- 映射到指定地址的任意端口, 使用 ip::containerPort 绑定 localhost 的任意端口到容器的 80 端口，本地主机会自动分配一个端口   

- 查看映射端口配置, docker port 来查看当前映射的端口配置，也可以查看到绑定,

- -p 标记可以多次使用来绑定多个端口,

docker run -d \
    -p 80:80 \
    -p 443:443 \
    nginx:alpine

    
