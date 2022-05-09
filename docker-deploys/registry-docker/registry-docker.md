# 仓库：是集中存放镜像的地方。注册服务器是管理仓库的具体服务器，每个服务器上可以有多个仓库，而每个仓库下面有多个镜像。

# docker login 登陆dockerhub

# docker logout 退出登陆

# docker search 查找镜像

# docker tag ubnutu:18.04 username/ubuntu:18.04打标签

# docker push username/ubuntu:18.04 推送标签数据到dockerhub

# 使用官方registry镜像运行       $ docker run -d -p 5000:5000 --restart=always --name registry registry

# docker tag ubuntu:latest 127.0.0.1:5000/ubuntu:latest 打标签到本地registry

# docker push 127.0.0.1:5000/ubuntu:latest 将镜像推送到本地registry

# 