---

date: 2023-03-23T14:01:26+0800
title: 用jq过滤嵌套数组中的字段
layout: post

---

```
jq '.[] | select (.csetname == "online") | .routes | .[] | .id?'
```
