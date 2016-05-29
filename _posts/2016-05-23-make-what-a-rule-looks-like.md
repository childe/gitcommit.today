---
layout: post
title:  "规则长什么样子"
date:   2016-05-23 13:52:10 +0800
categories: make
keywords: make
---

[https://www.gnu.org/software/make/manual/html_node/Rule-Introduction.html](https://www.gnu.org/software/make/manual/html_node/Rule-Introduction.html)

一个简单的makefile包含下面这样的一个或者多个规则

    target … : prerequisites …
        recipe
        …
        …

target(目标)经常是一个文件的名字, 这个文件由一个程序生成. 也可以是一个动作, 比如说clean. (参考[Phony Targets](https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html#Phony-Targets))

一个prerequisite(先决条件)就是目标的一个输入. 经常来说, 一个目标依赖多个先决条件. (译者注: 但实际上, 先决条件不仅仅可以是一个文件, 也可以是另外的target, 是链式依赖的)
