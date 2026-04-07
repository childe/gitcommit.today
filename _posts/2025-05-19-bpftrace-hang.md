---

date: 2025-05-19T11:22:36+0800
title: bpftrace卡住因DEBUGINFOD_URLS连接外网
layout: post

---

解决方法：

```bash
export DEBUGINFOD_URLS=""
```

通过 strace 看到 connect(106, {sa_family=AF_INET, sin_port=htons(443), sin_addr=inet_addr("18.219.164.54")}, 16) = -1 EINPROGRESS (Operation now in progress)

浏览器里面访问 18.219.164.54，看到报错 

This server couldn't prove that it's 18.219.164.54; its security certificate is from debuginfod.centos.org. This may be caused by a misconfiguration or an attacker intercepting your connection.
