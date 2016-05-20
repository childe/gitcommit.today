---
layout: post
title:  "how to write go code"
date:   2016-05-20 14:30:40 +0800
categories: golang
keywords: golang build install
---

看了[https://golang.org/doc/code.html](https://golang.org/doc/code.html), 觉得主要是讲了关于golang的Lib组织形式, 简单记录一下.

## 简介
本文描述了获取,构造和安装Go包以及Go工具的标准方法. go tool(就是go这个命令, 下面包含build run等命令)需要把你的代码按一定的格式组织起来.

## 代码组织

### 概述

- 一般来说, 所有的Go代码都放在一个单独的**workspace**
- 一个workspace包含多个版本控制仓库
- 每个仓库里面包括一个或者多个**packages**
- 一个package包含一个或者多个Go源代码, 这些源代码需要在一个单独的目录下面
- package的路径决定了它的**import path**

### Workspaces

一个workspace是这样一个目录结构, 包括下面三个子目录:

- src 包含Go的源码
- pkg 包含packagecfqj
- bin 包含可执行命令

go tool会build源码, 然后把结构安装到pkg和bin目录下.

### GOPATH 环境变量

GOPATH环境变量定义的你的workspace位置. 它是你需要的惟一一个环境变量. 它**一定不能**Go的安装路径使用同一个目录.

    $ mkdir $HOME/work
    $ export GOPATH=$HOME/work

为了方便, 可以把bin子目录放在PATH变量下

    $ export PATH=$PATH:$GOPATH/bin

### import path

*import path*是能惟一确认package路径的字符串. 

标准库的路径都是一个短路径, 像fmt, net/http. 你自己的代码, 需要选一个基础路径, 尽量不要和以后的标准库或者其他的第三方库有冲突.

比如说可以用 github.com/user做为基础路径, 在workspace下面新建一个目录存放你的源码.

### 你的第一个程序

在workspace目录下, 新建一个package目录做为package path.

    $ mkdir $GOPATH/src/github.com/user/hello

下一步, 在这个目录下创建一个hello.go:

```go
package main

import "fmt"

func main() {
        fmt.Printf("Hello, world.\n")
}
```

下面就可以build和install你的程序.

    $ go install github.com/user/hello

注意, 你可以在任意路径下运行这个命令, go会自己找到的. 如果你已经在包目录下面了, 直接跑 go install 也可以.

    $ cd $GOPATH/src/github.com/user/hello
    $ go install

会在GOPATH/bin下面生成一个hello可执行文件.

### 你的第一个库

让我们来写一个库, 并在hello中调用它

同样的, 先选一个package路径, 在这里我们用github.com/user/stringutil.

    $ mkdir $GOPATH/src/github.com/user/stringutil

在这个路径下创建go文件, reverse.go

```go
// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

现在可以build了(好像没什么用, 就是测试一下是不是能编译)

    $ go build github.com/user/stringutil

不会生成任何文件. 如果来`go install`, 可以在pkg目录下生成package对象.

我们在hello中引用它, 修改$GOPATH/src/github.com/user/hello/hello.go

```go
package main

import (
    "fmt"

    "github.com/user/stringutil"
)

func main() {
    fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
```

安装hello, 同时它的依赖(stringutil)也会被安装.

    $ go install github.com/user/hello

### Package names

go源码的第一行必须是
    
    package name

在这里, name就是这个包被引用用的名字,一个包里面的所有文件都必须用同一个名字.

包名需要是引用路径的最后一个元素, 如果包被引用为srypto/rto13, 包名则应该是rto13.

可执行文件的名字一定要是main.

一个二进制文件中引用的包, 并不要求它们的名字不一样, 但是完整路径肯定不能相同.
