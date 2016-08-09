---
layout: post
title:  "mesos scheduler与master通讯"
date:   2016-08-09 10:24:51 +0800
---

用的是http, 但是需要长链接. 服务器返回“200 OK” status code with Transfer-Encoding: chunked.  
所以如果一个Http库只能在链接close之后才解释, 这样的库**不能用**.

因为接下来的通讯需要用**另外**一个链接,我们叫做第二个链接. master收到请求后, 对第二个链接回复202. 然后把真正的响应回复到第一个长链接中.
