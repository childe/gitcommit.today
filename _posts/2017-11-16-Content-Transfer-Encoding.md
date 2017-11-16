---

layout: post
title: mail中的Content-Transfer-Encoding定义和用法
date: 2017-11-16 15:16:13 +0800

---

python发送邮件时出现乱码, 搜索了一下可以通过更改Content-Transfer-Encoding定义解决.

```
msg = MIMEText(None, "plain")
msg.replace_header('content-transfer-encoding', 'quoted-printable')
msg.set_payload(content, 'utf-8')
```

[https://stackoverflow.com/questions/25710599/content-transfer-encoding-7bit-or-8-bit](https://stackoverflow.com/questions/25710599/content-transfer-encoding-7bit-or-8-bit)

[https://www.w3.org/Protocols/rfc1341/5_Content-Transfer-Encoding.html](https://www.w3.org/Protocols/rfc1341/5_Content-Transfer-Encoding.html)
