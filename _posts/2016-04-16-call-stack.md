---
layout: post
title:  "栈帧中的一些概念 1"
date:   2016-04-16 11:41:43 +0800
categories: os
keywords: os call_stack
---

前不久想了解一下python中协程的概念, 看了greenlet源码. 看到slp_switch的时候, 里面是汇编语言, 搜索之后好像是需要了解栈帧的原理.

基础太差, 看了好多文章, 才明白了一点点.

1. 一个栈帧是一个被调用函数的一段内存空间
2. ebp指向栈帧的底部, esp指向顶部, 即当前指向的内存地址.
3. 入栈操作: push eax; 等价于 esp=esp-4,eax->[esp] 
4. 出栈操作：pop eax; 等价于 [esp]->eax,esp=esp+4

一些文章:

- [栈帧 yxysdcl的专栏ACM](http://blog.csdn.net/yxysdcl/article/details/5569351)
- [函数调用过程栈帧变化详解](http://www.cnblogs.com/fxplove/articles/2574451.html)
- [C函数调用机制及栈帧指针](http://blog.csdn.netss318/article/details/7185802)
- [Call stack From Wikipedia](https://en.wikipedia.org/wiki/Call_stack)
- [x86 Assembly Guide](http://www.cs.virginia.edu/~evans/cs216/guides/x86.html)
- [The Program Stack](http://www.c-jump.com/CIS77/ASM/Stack/index.html)
