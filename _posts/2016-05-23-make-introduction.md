---
layout: post
title:  "Makefile 简介"
date:   2016-05-23 13:52:10 +0800
categories: make
keywords: make
---

[https://www.gnu.org/software/make/manual/html_node/Introduction.html](https://www.gnu.org/software/make/manual/html_node/Introduction.html)

make需要一个Makefile, Makefile告诉make如何编译和链接.

这章会讨论一个简单的makefile, 描述了如何编译和链接一个文本编辑器程序,此程序由8个C文件和3个头文件组成.

make重新编译程序的时候, 会保证每一个有新变化的C文件被重新编译, 如果头文件有变动, 每个包含这个头文件的C文件也需要重新编译. 每个编译会生成一个对象文件, 然后再把所有对象文件链接生成新的可执行程序.

- [规则简介]()          规则长什么样子
- [简单的makefile文件]()
- [make如何工作]()      make是如何解析和执行makefile的
- [使用变量]()          变量让makefile更简单
- [make中的推断]()      让make推断recipe(怎么翻译?)
- [先决条件的组合]()    另一个风格的makefile
- [清理]()              清理目录的规则

