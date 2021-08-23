---

date: 2021-08-23T17:05:06+0800

---

有时候 Docker login/pull/push 的时候，出现 `http: server gave HTTP response to HTTPS client` 的报错，是因为 Docker 向这个服务器发出HTTPS 请求，但对方返回了 HTTP 响应。

一般来说呢，可能是自己的服务没有配置 HTTPS，这样的话，需要在 /etc/docker/deamon.json 里面配置一下。

记录一下，如果有如下配置，Docker 会是什么行为。

[https://docs.docker.com/registry/insecure/](https://docs.docker.com/registry/insecure/)

>   {
>     "insecure-registries" : ["myregistrydomain.com:5000"]
>   }

> Substitute the address of your insecure registry for the one in the example.
>
> With insecure registries enabled, Docker goes through the following steps:
>
> First, try using HTTPS.
> If HTTPS is available but the certificate is invalid, ignore the error about the certificate.
> If HTTPS is not available, fall back to HTTP.
