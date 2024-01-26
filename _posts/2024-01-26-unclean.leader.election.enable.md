---

date: 2024-01-26T15:43:46+0800
title: unclean.leader.election.enable
layout: post

---

提前设置好 unclean.leader.election.enable ，然后 replica 变成 offline。这个时候leader 不会选举出来。

需要等 leader 变成 -1 之后，再设置 unclean.leader.election.enable 才能选举出来
