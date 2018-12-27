---

title: golang中ticker会阻塞累积吗?
date: 2018-12-27T15:25:25+0800
layout: post

---

不会. 始终是{}里面执行完开始计算下一个duration.
