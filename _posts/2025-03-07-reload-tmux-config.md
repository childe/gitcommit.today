---

date: 2025-03-07T11:09:21+0800
title: 重新加载tmux配置
layout: post

---

## 1

按下 Prefix 键（默认是 Ctrl+b），然后按下 : 键，进入命令输入模式。

在命令行中输入以下内容并回车：

set -g mouse off

## 2

按下 Prefix 键（默认是 Ctrl+b），然后按下 : 键，进入命令输入模式。

在命令行中输入以下内容并回车：

source-file ~/.tmux.conf

您可以通过观察 tmux 的行为或运行以下命令来确认配置是否已生效：

tmux show-options -g
