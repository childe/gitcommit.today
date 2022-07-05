---

date: 2022-07-05T10:55:18+0800
title: jsonpath 里面的转义
layout: post

---

如果一个字段名字带点，可以这样转义

```
kd get pod -n redis 808098-2 -o jsonpath="{['metadata']['labels']['statefulset\.kubernetes\.io/pod-name']}"
```

如果要输出换行要：

```
kd get pod -n redis 808098-2 -o jsonpath="{['metadata']['labels']['statefulset\.kubernetes\.io/pod-name']} {'\n'}"
```
