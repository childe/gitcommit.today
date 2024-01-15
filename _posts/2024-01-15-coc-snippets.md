---

date: 2024-01-15T17:51:31+0800
layout: post

---

今天又折腾了一下 vim 里面的snippets，了解了几个概念。

原来 snippets 现在有两种格式，一种是 UltiSnips，一种是snipmate。 像 honza/vim-snippets 里面就包含两种格式的（但内容有所不同）

coc-snippets 好像对 UltiSnips 的支持更好，配置更丰富，相比之下，snipmate 的配置就比较少。

coc-snippets 加载 snipmate 格式的文件时，好像不能多目录加载，这个后面需要翻代码确认一下。我更新 runtimepath 也没能成。最后还是把自己 custom 的内容添加到 honza/vim-snippets/snippets/python.snippets 后面了。
