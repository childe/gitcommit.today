---

date: 2022-11-26T18:30:54+0800
title: MACOS 更新证书
layout: post

---


Add	Use command:

```
sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain ~/new-root-certificate.crt
```

Remove	Use command:

```
sudo security delete-certificate -c "<name of existing certificate>"
```
