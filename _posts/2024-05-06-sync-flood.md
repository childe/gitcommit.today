---

date: 2024-05-06T14:40:40+0800
title: tcp sync flood
layout: post

---

Apr 28 13:05:22 HOSTNAME01 kernel: TCP: request_sock_TCP: Possible SYN flooding on port 9092. Sending cookies.  Check SNMP counters.

收到告警，kafka broker 不通，到机器上看到这样的报错。原来是这个引起的。

哦，不对不对。这个是一周前的日志。
