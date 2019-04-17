---

date: 2019-04-17T13:50:17+0800
title: ES中已经删除文档会影响打分
layout: post

---

被删除的文档如果没有merge, 没有被彻底删除, 还是会影响打分. 因为被计算在docCount里面.
