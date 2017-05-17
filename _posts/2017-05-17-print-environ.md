---

layout: post
title:  "打印环境变量"
date:  2017-05-17 14:17:53 +0800

---

```
#include <stdio.h>
#include <stdlib.h>

extern char **environ;

int main(int argc, char const* argv[])
{
    char **e = environ;
    while ( *e != NULL ){
        printf("%s\n", (*e));
        e += 1;
    }
    return 1;
}
```
