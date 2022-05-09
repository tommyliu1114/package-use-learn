# 数据卷： 是一个可供一个或多个容器使用的特殊目录，它绕过 UFS，可以提供很多有用的特性
- 数据卷 可以在容器之间共享和重用

- 对 数据卷 的修改会立马生效

- 对 数据卷 的更新，不会影响镜像

- 数据卷 默认会一直存在，即使容器被删除 

# 创建一个数据卷：
- docker volume create my-vol

# 查看所有的数据卷：
- docker volume ls 

# 查看数据卷信息：
- docker volume inspect mms17-vol-1 

# 启动一个挂载数据卷的容器：使用 --mount 标记来将 数据卷 挂载到容器里。在一次 docker run 中可以挂载多个 数据卷
- docker run -d -P --name web --mount source=mms17-vol-1,target=/usr/share/nginx/html nginx:v3-mms17 

# 删除数据卷：
- docker volume rm mms17-vol-1

# 删除所有无主的数据卷：
- docker volume prune 


