# 新建网络： docker network create -d bridge my-net 

# 运行容器连接到网络：
docker run -it --rm --name busybox2 --network my-net busybox sh