---

date: 2025-02-25T18:24:57+0800
title: podman-compose启动报错missing networks default
layout: post

---

Q: macos上，podman-compose启动失败，报错 "RuntimeError: missing networks: default"

A: 根据搜索结果，Podman Compose 在某些情况下无法自动创建默认网络，需要在 docker-compose.yml 文件中显式定义默认网络。例如：

```yaml
version: '3.8'
services:
  your_service:
    image: your_image
    networks:
      - default

networks:
  default:
    driver: bridge
```
