---

date: 2024-03-26T15:20:37+0800
layout: post

---

https://github.com/moby/moby/issues/39869

目前看起来，只能重启才能加载系统证书。

如果不想重启，有两个办法

1. 配置 insecure-registries
2. 创建 /etc/docker/certs.d/{domain} 文件夹，并将证书放在这下面
