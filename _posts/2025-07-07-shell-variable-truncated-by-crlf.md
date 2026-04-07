---

date: 2025-07-07T14:39:30+0800
title: shell变量末尾有换行符导致echo输出截断
layout: post

---

linux shell

通过一个 shell 脚本获取了一个 token 变量。然后echo "header: $token http://xxx.corp.com" ，发现被 token 截断了。

终于发现 token 变量最后有一个`\n\r`导致。
