---
layout: post
title:  "docker内存限制不能用free来查看"
date:   2016-04-19 00:16:10 +0800
---

这样可以限制容器中所有进程使用的内存数量

```sh
docker run -it -m 300M ubuntu:14.04 /bin/bash
```

但它是用cgroup来实现的, 在窗口中用free看不出来.

可以写一个程序, 循环申请一定量的内存, 看看申请到多少的时候出错.
