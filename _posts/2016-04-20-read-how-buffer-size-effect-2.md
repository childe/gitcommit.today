---
layout: post
title:  "每次read时的buffer size如何影响性能2"
date:   2016-04-21 11:50:26 +0800
---

昨天的时候有个疑问, /proc/diskstats里面的第四列, 也就是reads completed是什么. 和iostat显示的tps对比之后, 确认了, read tps取的值和reads completed就是一个.

那应该就是说,  read(fd, 1024000)的时候, 未必是一次"I/O request to the device". 具体系统怎么处理read的, 继续学习.
