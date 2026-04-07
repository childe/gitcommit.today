---

date: 2025-02-10T15:05:00+0800
title: 从零构建LLM学习笔记一：token到向量嵌入
layout: post

---

1. 一句话，像`"I am a student."`，am 的 tokenID 是一个[xx,xx,xx]这样的“三维”向量。第书的第二章讲的encode token，是变成一个数字，怎么到了第三章，突然成了一个三维向量？中间发生了哪些操作？

2. 在第三章的“Implementing self-attention with trainable weights”这一节，说到“我们成功将6个输入标记从三维投影到二维嵌入空间上”。为什么要把三维转成二维？这里说的“embedding space”又是指什么？
