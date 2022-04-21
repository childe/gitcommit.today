---

date: 2022-04-21T14:47:16+0800
layout: post

---

今天安装了 https://github.com/dandavison/delta ，并选了一个和 Iterms 搭配的主题。但是 zsh powerlevel10k theme 里面 prompt 显示不出来 git branch 了。

是因为 delta 项目里面自带的 theme 配置文件可能和我的 git 版本(2.28.0)不兼容。把 zebra-* 里面的换行的配置放到一行就好了。
