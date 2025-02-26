---

date: 2025-02-26T19:15:24+0800
layout: post

---

dockeer/docker-compose 切换到 podman/podman-compose 之后，compose 起不来了。

中间各种查日志和资料不说，一个可能的原因是这样的（带猜测性质）：

是docker-compose 里面，如果不配置network，会使用default这个network。 但换到podman之后，没有default了（有一个叫podman的)，所以就找不到network了。

解决方法是在 docker-compose.yml 里面，每个 service 都加上network的配置。

networks:
  - podman

如果之前配置的 default networkd，换成 podman
