---
layout: post
title:  "数据类型对es bulk的影响"
date:   2016-06-21 14:26:45 +0800
categories: es
keywords: elastcisearch mapping bulk
---


共14912997 数据,就一个字段, 平均31bytes, 里面有平均7个token

## 把字符串按mapping转成integer ; 1 shard
merge后291.4m
es cpu time 12m
hangout user time 8m
search 98ms

## string ; 1 shars
merge后325.2m
es cpu time 13m10s
hangout user time 8m
search 120ms

## string not analyzed; 1 shars
merge后358.5m
search 400ms

## string ; 8 shars
es cpu time  12m19s
hangout user time  4m
