---
layout: post
title:  "每次read时的buffer size如何影响性能"
date:   2016-04-20 13:53:49 +0800
---

磁盘IO是比较费时的操作. 操作系统在每次读取磁盘的时候, 会认为接下来的文件内容可能在接下来很快再次读取,所以会预先就先读到内存中.

这样可以查看当前系统的read ahead的值, 单位是512 bytes.

    blockdev --getra /dev/sda1

这样可以配置read ahead的值, 单位是512 bytes.

    blockdev --setra 1024 /dev/sda1


简单的脚本测试一下:

```sh
echo 3 > /proc/sys/vm/drop_caches

read1=$(cat /proc/diskstats | grep sda1 | awk '{print $4}')
dd if=test of=/dev/null ibs=$@
read2=$(cat /proc/diskstats | grep sda1 | awk '{print $4}')
echo $((read2-read1))
```

运行结果:

```sh
# blockdev --setra 256 /dev/sda1 
# ./go.sh 4096
4194
#
# blockdev --setra 128 /dev/sda1 
# bash go.sh 4096
8233
```

[关于/proc/diskstats的解释](https://www.kernel.org/doc/Documentation/iostats.txt)

# 疑问1

我是计算的 'reads completed' 的差值.

    This is the total number of reads completed successfully.

这个到底是指的什么?

系统调用的次数? 看上面的测试结果应该不是.  
实际磁盘IO的次数? 也会有矛盾.


# 疑问2

如果readahead是16(512 bytes), 16*512=8192

read(16384)的时候, 会有几次磁盘IO? 看reads completed的值, 8192和16384没有区别.

# 非疑问
st_blksize 是4096, 所以readahead小于8是没效果的.
