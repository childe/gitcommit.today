---

date: 2023-01-12T13:19:58+0800
title: 'harbor devcenter-api 测试报 csrf 错误'
layout: post

---

harbor 2.0 版本

harbor devcenter-api 页面测试某些 API（不确定是不是全部 POST 请求） 的时候，报错 'CSRF token invalid'

直接 curl 没问题。

查下来，是因为请求的 Cookie 里面带 sid  字段，Harbor 认为这是一个带 session 的请求，就不会跳过 csrf 验证。
