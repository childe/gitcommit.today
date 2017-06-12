---

layout: post
title:  "nginx做websocket代理的idle timeout问题"
date:   2017-06-12 21:11:11 +0800

---

需要在proxy_pass里面加上proxy_read_timeout 60;

Defines a timeout for reading a response from the proxied server. The timeout is set only between two successive read operations, not for the transmission of the whole response. If the proxied server does not transmit anything within this time, the connection is closed.
