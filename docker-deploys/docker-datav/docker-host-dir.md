# 挂载本地主机目录为数据卷：

- -mount 标记可以指定挂载一个本地主机的目录到容器中去


- docker run -d -P --name web --mount type=bind,source=/home/sen0324/docker-dirs,target=/usr/share/nginx/html nginx:v3-mms17