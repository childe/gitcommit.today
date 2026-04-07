---

date: 2023-03-19T21:13:24+0800
title: docker buildx构建多架构镜像
layout: post

---

docker buildx build --platform linux/amd64,linux/arm64 --push -t membermatters/membermatters .
