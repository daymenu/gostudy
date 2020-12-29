# Wire GO中的自动初始化

Wire是一个代码生成工具，它使用依赖注入自动连接组件。组件之间的依赖关系在Wire中表示为函数参数，鼓励显式初始化，而不是全局变量。由于Wire没有运行时状态或反射，因此编写用于Wire的代码即使对于手写的初始化也很有用。

## 安装

```bash
go get github.com/google/wire/cmd/wire
```

> $GOPATH/bin 添加到你的$PATH中

## 文档

[Tutorial](https://github.com/google/wire/blob/master/_tutorial/README.md)
[User Guide](https://github.com/google/wire/blob/master/docs/guide.md)
[Best Practices](https://github.com/google/wire/blob/master/docs/best-practices.md)
[FQA](https://github.com/google/wire/blob/master/docs/faq.md)

-----------我也是有底线的-----------
