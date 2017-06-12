---

layout: post
title:  "nginx中location的匹配规则"
date:   2016-07-01 14:36:45 +0800

---

官方资料见 [http://nginx.org/en/docs/http/ngx_http_core_module.html#location](http://nginx.org/en/docs/http/ngx_http_core_module.html#location)

location语法是这样的 `location [ = | ~ | ~* | ^~ ] uri { ... }` , []里面的东西代表可有可无. 举一个例子

```
location = / {
    [ configuration A ]
}

location / {
    [ configuration B ]
}

location /documents/ {
    [ configuration C ]
}

location ^~ /images/ {
    [ configuration D ]
}

location ~* \.(gif|jpg|jpeg)$ {
    [ configuration E ]
}
```

nginx在匹配路径之前, 会首先把uri中%XX这样格式的内容先解码, 然后把. ..转成绝对路径, 以及把多个//合并成一个/

再说明匹配规则前, 还是先解释一下上面这个语法什么意思. ~ ~*这两个代表uri是一个正则表达式, ~是大小写敏感, ~*是大小写不敏感.  **^~可不是正则的意思,后面会有解释**. 其它情况下, uri代表一个路径前缀.

匹配路径按以下规则进行:

1. 先序遍历所有路径前缀, 记录下来. 然后按顺序遍历所有正则,选用第一个正则匹配成功的结果. 如果没有正则可以成功匹配, 选用最长长度的前缀路径.

2. 如果**最长的**前缀路径有 ^~ 修饰符, 就不再进行正则匹配这一步了.

3. 如果前缀路径有 = 修饰符, 不再继续寻找

4. uri前面加一个@, 叫做 named location, 普通的请求不管他, 只用在一些内部的跳转上, 比如下面这样

        location / {
            error_page 404 = @fallback;
        }

        location @fallback {
            proxy_pass http://backend;
        }

5. 如果一个前缀路径最后以/结尾, 而且请求被proxy_pass, fastcgi_pass, uwsgi_pass, scgi_pass, or memcached_pass中的一个处理, 这个情况比较特殊: 请求如果不以/结尾, 会返回一个301的重定向响应, 再最后加上/. 如果不想这样, 就要用 = 修饰符做精确匹配, 像下面这样

        location /user/ {
            proxy_pass http://user.example.com;
        }
        location = /user {
            proxy_pass http://login.example.com;
        }
