---

date: 2023-01-17T17:06:33+0800
layout: post

---

[https://pkg.go.dev/runtime#MemStats](https://pkg.go.dev/runtime#MemStats)

一个例子

```
# runtime.MemStats
# Alloc = 0.46 GB
# TotalAlloc = 427382 GB
# Sys = 6.25 GB
# Lookups = 0
# Mallocs = 1864
# Frees = 1864
# HeapAlloc = 0.46 GB
# HeapSys = 5.93 GB
# HeapIdle = 5.21 GB
# HeapInuse = 0.72 GB
# HeapReleased = 4.03 GB
# HeapObjects = 4048368
# Stack = 6225920 / 6225920
# MSpan = 9607448 / 73842688
# MCache = 83328 / 98304
# BuckHashSys = 2.13 MB
# GCSys = 0.23 GB
# OtherSys = 13 MB
# NextGC = 0.83 GB
# LastGC = 1653288332185470574
```

```
 MEM USAGE / LIMIT
918.7MiB / 125.5GiB
```

可见，RSS 大概是 HeapInuse GCSys 的总和。

注意，HeapInuse 是占用的内存，比 HeapAlloc 大。HeapAlloc 是已经分配给对象的，剩下的是已经占用的内存，可以随时分给差不多大小的对象。

HeapInuse is bytes in in-use spans. In-use spans have at least one object in them. These spans can only be used for other objects of roughly the same size.

HeapInuse minus HeapAlloc estimates the amount of memory that has been dedicated to particular size classes, but is not currently being used. This is an upper bound on fragmentation, but in general this memory can be reused efficiently.
