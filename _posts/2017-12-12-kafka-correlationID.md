---

layout: post
title: kafka中Server返回的correlationID
date: 2017-12-12 14:12:33 +0800

---

原来kafka server返回的correlationID并不是对应的request中的, 而是依次增加的?

哦, 搞错了. 是对应request里面的.
