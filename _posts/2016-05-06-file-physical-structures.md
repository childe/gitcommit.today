---
layout: post
title:  "问题 文件的物理结构是什么样的"
date:   2016-05-06 12:06:19 +0800
categories: os
keywords: os file structure
---
文件系统中, 文件的物理结构是什么样的呢? 文件系统中, 怎么定义的文件的存储呢? 先记下这个问题, 以后看.

```c
#include <fcntl.h>
#include <unistd.h>
#include <stdio.h>

int main(){
    int f = open("blank",O_WRONLY);
    if (f<0){
        perror("open error");
    }
    printf("%d\n",f);
    write(f,"abcdefg",7);
    lseek(f,10000000,SEEK_CUR);
    write(f,"abcdefg\n",8);
    close(f);
}
```

写一个空洞文件.

```
root@2d3967f5850a:/tmp/test# ll blank
-rw-r--r-- 1 root root 10000015 May  6 03:56 blank
root@2d3967f5850a:/tmp/test# du -sh blank
104K    blank 
```

文件在大概10M, 但实际占用的磁盘空间只有104K

但读取文件时候, 是会读取10M字节的内容的, cat blank > blank2, 然后再du也可以佐证.

那文件系统中, 文件的物理结构是什么样的呢? 文件系统中, 怎么定义的文件的存储呢? 先记下这个问题, 以后看.
