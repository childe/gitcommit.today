---

date: 2024-03-14T18:51:54+0800
layout: post

---

export DEBUGINFOD_URLS=""

在使用 bcc profile 工具做火焰图的时候，卡在输出时。

通过 strace 和 dcpdump 看到他在请求 https://debuginfod.centos.org.
