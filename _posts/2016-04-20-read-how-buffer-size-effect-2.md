---
layout: post
title:  "每次read时的buffer size如何影响性能2"
date:   2016-04-21 11:50:26 +0800
---

昨天的时候有个疑问, /proc/diskstats里面的第四列, 也就是reads completed是什么. 和iostat显示的tps对比之后, 确认了, read tps取的值和reads completed就是一个.

那应该就是说,  read(fd, 1024000)的时候, 未必是一次"I/O request to the device". 具体系统怎么处理read的, 继续学习.

下午更新:

在高手的指点下, 想通了. 至少在读文件(而不是裸设备)的时候, read(fd, 1024000)不是一次IO请求. 下面具体说一下:

如果文件系统的blocksize也是4096. readahead设置成16, 也就是16*512=8192字节, 2个block.

我们假设文件是顺序的.

### 我原来的想法
因为文件在磁盘中顺序存放的, 所以read(fd, 1024000)的时候, 一次性取1024000字节, 可以极大的减少磁盘IO次数. 但和测试下来的结果不符合, 所以很困惑.

### 现在的理解
OS在执行read file的时候,可不知道文件在磁盘中是顺序存在的,它需要在分析文件当前的block之后才知道下一个block在磁盘的哪里.

如果一次性读取1024000这么多字节, 非常有可能后面很多数据都浪费了. 所以每次IO操作时, OS还是按照readahead配置来读.

在`read(fd, 1024000)`的时候, 是250个block. 因为有readahead, 所以一次会读8192字节. 如果磁盘是顺序的, 那么就是需要1024000/8192=125次IO requests
