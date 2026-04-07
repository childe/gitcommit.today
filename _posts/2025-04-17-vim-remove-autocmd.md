---

date: 2025-04-17T11:59:31+0800
title: vim移除或动态控制autocmd
layout: post

---


```
:autocmd! BufWritePre *.py
```


方法二：通过变量动态控制（更灵活）
如果你希望未来能快速切换启用/禁用状态，可以修改原配置为 条件触发：

修改原有自动命令（未来长期方案）：
在配置文件中将原有命令改为：

```vim
let g:enable_organize_import = 1  " 默认启用
autocmd BufWritePre *.py if g:enable_organize_import | silent call CocAction('runCommand', 'editor.action.organizeImport') | endif
```

临时禁用（当前会话中）：

```vim
:let g:enable_organize_import = 0
```

重新启用（需要时）：

```vim
:let g:enable_organize_import = 1
```
