---

date: 2021-12-11T00:58:45+0800

---

问题1，一个打到 Docker Registry Mirror 的请求，为什么透传了两次到后端呢？

问题2，Proxy 里面的 inflight 又能起什么作用呢？

问题3，为什么有的机器把同一时间到来的所有请求都打到了后端，而有的机器只透传了非常少的几个请求，后面就走缓存了呢？


问题 1 回答，因为 pbs.copyContent(ctx, dgst, w) 里面请求了一次，pbs.storeLocal(ctx, dgst) 里面又请求了一次。

问题 2 回答，inflight 是为了避免两个同样的请求都 storeLocal 。

问题 3 Debug 方向：

1. inflight 打印一下，看看哪哪请求进到里面？
2. 本地调试，断点看一下 filesystem.GetContent 是哪里调用来的。
3. pbs.serveLocal(ctx, w, r, dgst) 之前先 sleep random，拿一个小 Layer 来测试的时候，sleep 短的那个请求肯定已经完成了，所以只能会一次请求到后面。
4. 拿大 Layer 来测，两个请求应该会一起透传到后面吧。


2. filesystem.GetContent 是 Stat 的时候，需要先读取 link 内容，然后再去 stat link 文件
