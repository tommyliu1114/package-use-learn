# 需要注意的：
- 没有使用很多个 RUN 一一对应不同的命令，而是仅仅使用一个 RUN 指令，并使用 && 将各个所需命令串联起来。

- Dockerfile 支持 Shell 类的行尾添加 \ 的命令换行方式，以及行首 # 进行注释的格式

- sudo docker build -t nginx:v3-mms17 .

# 上下文的概念：
- 一般来说，应该会将 Dockerfile 置于一个空目录下，或者项目根目录下。如果该目录下没有所需文件，那么应该把所需文件复制一份过来。如果目录下有些东西确实不希望构建时传给 Docker 引擎，那么可以用 .gitignore 一样的语法写一个 .dockerignore，该文件是用于剔除不需要作为上下文传递给 Docker 引擎的。

# Run命令：
- shell 格式： RUN <命令>
    - 示例：     RUN echo '<h1>Hello, Docker!</h1>' > /usr/share/nginx/html/index.html
- exec 格式： RUN ["可执行文件", "参数1", "参数2"]