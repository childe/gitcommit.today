---

date: 2024-12-28T22:30:31+0800
title: Golang时间解析格式毫秒需要点或逗号分隔
layout: post

---

> case '.', ',': // ,000, or .000, or ,999, or .999 - repeated digits for fractional seconds.

看 golang time lib 代码，000 前面需要有, 或者是.  ， 才是合法的

2006-01-02 150405000 像这种是解析不出来的，不管你的 format 怎么配置。
