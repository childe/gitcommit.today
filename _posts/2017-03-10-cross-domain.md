---

layout: post
title:  ajax跨域请求
date:   2017-03-10 10:46:52 +0800

---

正常情况, ajax跨域调用的时候, 浏览器会先发送options看服务端是不是可以允许域名.

但是如果有basic auth的话, options不会把auth信息发过去, 然后就401了.

难道还要nginx单独配置options方法?

