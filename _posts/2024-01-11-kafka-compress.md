---

date: 2024-01-11T01:01:35+0800
title: kafka 压缩对比
layout: post

---

原始数据26MB

gzip: 6.2M, 3.00s user 0.67s system 65% cpu 5.597 total

lz4: 11M 2.9s user 0.70s system 66% cpu 5.712 total

snappy: 13M 2.81s user 0.69s system 59% cpu 5.874 total

但是 gzip 的解压速度应该比较慢，未测试。

综合来看，lz4 不错。
