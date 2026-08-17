---

date: 2026-08-17T17:01:44+0800
title: clientA-review-clientB场景下-如何明确把问题告知对方
layout: post

---

codex写的代码让claude去review，如果有问题交给codex排查，如此反复，直到两者都认为MR可以Merge了。
期间，对于“问题是什么”的沟通，可以使用测试代码的方式。claude如果认为有问题，可以添加test code 确认问题的确存在。然后让codex 运行测试脚本明白是什么问题。
