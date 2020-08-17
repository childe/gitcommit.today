---

date: 2020-08-17T17:49:44+0800
title: 'zsh: no matches found'

---

经常见这个报错吧.

zsh 里面有个 NOMATCH 选项, 意思是: 如果 pattern of filename 如果没有匹配任何文件, 会打印上面这个错误.

如果把这个项目关掉, 就不会打印错误, 而是原封不动的继续执行.

`setopt +o onmatch`
