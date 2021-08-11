---

date: 2021-08-11T11:00:46+0800

---

[https://stackoverflow.com/questions/8591600/nginx-proxy-pass-based-on-whether-request-method-is-post-put-or-delete](https://stackoverflow.com/questions/8591600/nginx-proxy-pass-based-on-whether-request-method-is-post-put-or-delete)

```
location ~* "(ngx.HTTP_POST|ngx.HTTP_DELETE|ngx.HTTP_PUT)" {
    proxy_pass http://127.0.0.1:8080;
```

```
    server {
  location / {
    # This proxy_pass is used for requests that don't
    # match the limit_except
    proxy_pass http://127.0.0.1:8080;

    limit_except PUT POST DELETE {
      # For requests that *aren't* a PUT, POST, or DELETE,
      # pass to :9080
      proxy_pass http://127.0.0.1:9080;
    }
  }
}
```

[If is Evil... when used in location context](https://www.nginx.com/resources/wiki/start/topics/depth/ifisevil/)

```
location / {
    error_page 418 = @other;
    recursive_error_pages on;

    if ($something) {
        return 418;
    }

    # some configuration
    ...
}

location @other {
    # some other configuration
    ...
}
```
