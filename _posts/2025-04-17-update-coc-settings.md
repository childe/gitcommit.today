---

date: 2025-04-17T11:35:46+0800
title: 运行时更新coc配置无需重启vim
layout: post

---

更新coc-setting.json 配置之后，其实 vim 里面会自动更新配置，不需要关闭 vim 再打开。

如果不想持久化这个配置，只是临时应用，可以像下面这样在vim 里面执行：

```
:call coc#config('preferences.formatOnSave', v:true)
```
