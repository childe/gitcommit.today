---

date: 2021-12-11T00:58:45+0800

---

一个打到 Docker Registry Mirror 的请求，为什么透传了两次到后端呢？

Proxy 里面的 inflight 又能起什么作用呢？

为什么有的机器把同一时间到来的所有请求都打到了后端，而有的机器只透传了非常少的几个请求，后面就走缓存了呢？


问题 1 回答，因为 pbs.copyContent(ctx, dgst, w) 里面请求了一次，pbs.storeLocal(ctx, dgst) 里面又请求了一次。

问题 2 回答，inflight 是为了避免两个同样的请求都 storeLocal 。

问题 3 Debug 方向：

1. inflight 打印一下，看看哪哪请求进到里面？
2. 本地调试，断点看一下 filesystem.GetContent 是哪里调用来的。
