---

date: 2022-10-25T18:04:40+0800
title: OBJC_DISABLE_INITIALIZE_FORK_SAFETY
layout: post

---

在运行一些 Python 应用时，遇到下面这个错误。

```
objc[12297]: +[__NSCFConstantString initialize] may have been in progress in another thread when fork() was called. We cannot safely call it or ignore it in the fork() child process. Crashing instead. Set a breakpoint on objc_initializeAfterForkError to debug.
```

解决：

```sh
export OBJC_DISABLE_INITIALIZE_FORK_SAFETY=YES
```
