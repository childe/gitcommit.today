---

date: 2025-02-06T16:20:40+0800
layout: post

---

将 GitHub Pages 的自定义域名迁到 cloudflare 之后，一直死循环 301。

原因是，cloudflare 默认的 SSL 配置是 "灵活"：仅在访问者与 Cloudflare 之间启用加密。这可以避免浏览器发出安全警告，但 Cloudflare 与您的源服务器之间的所有连接均通过 HTTP 建立。

而我的 github pages 配置了 enforce https，所以就会出现死循环。

但是，还有一个github pages 项目使用不能配置 enforce https，配置页面显示"Enforce HTTPS — Unavailable for your site because your domain is not properly configured to support HTTPS"，还不明白怎么回事，怎么解决。目前是在 cloudflare 配置了始终重定向到 https
