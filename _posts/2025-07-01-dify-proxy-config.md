---

date: 2025-07-01T15:03:37+0800
title: dify部署在内网时的代理配置
layout: post

---

dify 版本 v1.5.0

添加以下配置到你的 Squid 配置文件中，以允许访问 Dify 的 Marketplace：

```
acl allowed_domains dstdomain .marketplace.dify.ai

cache_peer proxy.corp.com parent 8080 0 no-query no-digest name=company_proxy
cache_peer_access company_proxy allow allowed_domains
cache_peer_access company_proxy deny all
never_direct allow all
```
