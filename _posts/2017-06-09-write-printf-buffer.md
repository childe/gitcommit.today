---

layout: post
title:  "缓冲"
date:   2017-06-09 10:57:07 +0800

---

apue里面的8.1示例代码 

```
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>

int glob = 6;
char buf[] = "a write to stdout\n";

int main(int argc, char const* argv[])
{
    int var;
    pid_t pid;
    var = 88;
    if (write(STDOUT_FILENO, buf, sizeof(buf)-1 ) != sizeof (buf)-1){
        printf("write error");
        exit(1);
    }
    printf("before fork\n");

    if ((pid = fork()) <0){
        printf("fork error");
        exit(1);
    }else if (pid == 0){
        glob +=1 ;
        var += 1;
    }else {
        sleep( 2 );
    }
    printf("pid=%d, glob=%d, var=%d\n", pid, glob, var);
    return 0;
}
```

write是不带缓冲的, printf有. 标准输出重定向到文件时, 会有两行 before fork

他们的文件指针指向的文件表以及文件表中的文件偏移量是一样的, 这样才能保证写的数据不会覆盖.
