# CMD命令格式：
- shell格式： CMD <命令>

- exec格式：  CMD ["可执行文件","参数1","参数2",...]
    - 示例：      CMD echo $HOME
    - 解析为：     CMD [ "sh", "-c", "echo $HOME" ] 

- 参数列表方式： CMD ["参数1","参数2",...] ,在指定了 ENTRYPOINT 指令后，用 CMD 指定具体的参数

