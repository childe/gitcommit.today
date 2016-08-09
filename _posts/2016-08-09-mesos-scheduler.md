---
layout: post
title:  "mesos scheduler与master通讯"
date:   2016-08-09 10:24:51 +0800
---

用的是http, 但是需要长链接. 服务器返回“200 OK” status code with Transfer-Encoding: chunked.  
所以如果一个Http库只能在链接close之后才解释, 这样的库**不能用**.
