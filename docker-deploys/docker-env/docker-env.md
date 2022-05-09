# ENV设置环境变量

- ENV <key> <value>
- ENV <key1>=<value1> <key2>=<value2>...
- 换行，以及对含有空格的值用双引号括起来的办法，这和 Shell 下的行为是一致的
```bash

    ENV VERSION=1.0 DEBUG=on \
        NAME="Happy Feet"

```