---
layout: post
title:  "brew安装应用时出错"
date:   2016-08-26 09:32:04 +0800
---

brew安装Python3 和 vim时, 都是报错

    ==> ./configure --prefix=/usr/local/Cellar/apr/1.5.2/libexec
    checking for suffix of executables...
    checking whether we are cross compiling... configure: error: in `/private/tmp/vim-7.4.2235/':
    configure: error: cannot run C compiled programs.
    If you meant to cross compile, use `--host'.
    See `config.log' for more details

看github上面说是因为osx 10.11与老版本的xcode不兼容了

解决:

    xcode-select --install
    sudo xcode-select -switch /Applications/Xcode.app/Contents/Developer/
