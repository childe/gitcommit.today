---

date: 2025-11-11T16:53:56+0800
title: Django应用中NO_PROXY通配符格式错误导致代理绕过失效
layout: post

---


我一个django应用中配置了环境变量：`environment=NO_PROXY="localhost,127.0.0.1,*.ctripcorp.com",https_proxy="http://proxygate2.ctripcorp.com:8080",http_proxy="http://proxygate2.ctripcorp.com:8080"`，但是在SSO认证ticket的时候，请求 https://cas.fat.qa.nt.ctripcorp.com 的时候，还是使用了代理。问题在哪里？

这是一个常见的**NO_PROXY配置陷阱**。

您的问题很可能出在以下两点：

1. 域名格式问题（最可能的原因）
NO_PROXY 变量的值通常是不区分大小写的，但它有一个重要的格式要求：对于要排除整个域名的请求（包括子域名），您需要在域名前面加上一个点号 (.)。

您的配置是：

NO_PROXY="localhost,127.0.0.1,*.ctripcorp.com"
虽然有些系统或库可能支持 *.ctripcorp.com 这样的通配符，但更通用且推荐的格式（尤其是在 Python/Django 应用中常用的 requests 库或底层网络库中）是：

对于匹配 ctripcorp.com 及其所有子域名（如 cas.fat.qa.nt.ctripcorp.com）的请求，应使用：

.ctripcorp.com
修改建议：

将您的环境变量修改为：

environment=NO_PROXY="localhost,127.0.0.1,.ctripcorp.com",https_proxy="http://proxygate2.ctripcorp.com:8080",http_proxy="http://proxygate2.ctripcorp.com:8080"
（注意 .ctripcorp.com 前面的点）

2. 环境变量名称的大小写问题
虽然您在描述中使用了 NO_PROXY（全大写），但在某些系统或网络库中，**no_proxy（全小写）**可能也是需要配置的，甚至优先级更高。

为了保险起见，建议您同时设置两个版本（如果您的运行环境允许）：

更保险的修改建议：

environment=NO_PROXY="localhost,127.0.0.1,.ctripcorp.com",no_proxy="localhost,127.0.0.1,.ctripcorp.com",https_proxy="http://proxygate2.ctripcorp.com:8080",http_proxy="http://proxygate2.ctripcorp.com:8080"
总结
请您尝试使用 .ctripcorp.com 格式并重启您的 Django 应用容器或服务，以确保新的环境变量生效。

如果修改后仍然不生效，您可能需要：

在 Django 代码中确认变量是否加载： 在应用启动时或进行 SSO 请求前，打印出 os.environ.get('NO_PROXY') 和 os.environ.get('no_proxy') 的值，确认它们确实是您配置的带有点号的版本。

确认发起 SSO 请求的库： 确定 Django 应用中使用哪个库来发起对 https://cas.fat.qa.nt.ctripcorp.com 的请求（例如 requests、urllib、或某个 SSO 客户端库），然后查阅该库的文档，确认它是否支持读取标准的 NO_PROXY/HTTPS_PROXY/HTTP_PROXY 环境变量。

您希望我帮您搜索一下关于 Django/Python 中处理代理环境变量的最佳实践吗？
