---

date: 2025-05-19T11:44:21+0800
layout: post

---

```
#!/usr/bin/env bpftrace

tracepoint:block:block_rq_issue
/ args->dev == (uint32)((8<<20)+144) /  // 8:144 对应 /dev/sdj 的主/次设备号
{
    printf("%-6d %-16s %-6d\n",pid,comm,args->dev);
}
```
