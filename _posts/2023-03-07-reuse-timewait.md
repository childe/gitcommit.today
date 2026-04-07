---

date: 2023-03-07T11:57:06+0800
title: TCP TIMEWAIT连接复用行为测试
layout: post

---

myip:port -> hub.corp.com:80

hub.corp.com 上面大量 timewait

新的 myip:port -> hub.corp.com:80 进来的时候，hub.corp.com 上面这个连接还是 timewait

会发生什么情况。

测试下来，服务端会丢弃这个 timewait 端口，然后正常建连。
