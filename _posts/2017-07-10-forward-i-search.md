---

layout: post
title:  "forward-i-search in bash"
date:   2017-07-10 17:22:07 +0800

---

[https://stackoverflow.com/questions/549810/control-r-reverse-i-search-in-bash-how-do-you-reset-the-search-in-cygwin](https://stackoverflow.com/questions/549810/control-r-reverse-i-search-in-bash-how-do-you-reset-the-search-in-cygwin)

[https://superuser.com/questions/159106/reverse-i-search-in-bash](https://superuser.com/questions/159106/reverse-i-search-in-bash)

[https://superuser.com/questions/472846/how-to-reverse-i-search-back-and-forth](https://superuser.com/questions/472846/how-to-reverse-i-search-back-and-forth)

Ctrl-R可以用来反向搜索曾经搜索行的命令, 用过的同学一定爱不释手, 可是怎么样正向搜索呢? 万一多按了几下Ctrl-R就要从头再来了, 真的很麻烦.

Ctrl-S是可以正向搜索的, 但是如果你按了它没反应, 是因为Ctrl-S被另外一个功能覆盖了.

Ctrl-S可以停止输出, 比如在`ping www.ctrip.com`的时候, Ctrl-S就可以把输出暂时屏蔽了. Ctrl-Q可以再次输出来.

stty -ixon可以把这个功能停掉, stty ixon可以再次启用这个功能.

反正我的感觉是 stty -ixon更好一些吧, 'START/STOP output control'用到并不多, 但forward-i-search却是会经常用.
