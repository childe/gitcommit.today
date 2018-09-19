---

layout: post
title: mmap中的read ahead
date: 2018-08-15T10:38:25+0800

---

[https://github.com/elastic/elasticsearch/issues/27748](https://github.com/elastic/elasticsearch/issues/27748)

上面这个issue提到 

  The bottleneck came from the memory management read ahead which is different from the block device read ahead.

  The memory management read ahead is only used for mapped files and the block device read ahead for all read operations on storage devices whether using niofs or mmapfs.

  In fact, the 2MB read ahead limit I observed in my tests seems to be hard coded in the kernel:
  https://github.com/torvalds/linux/blob/master/mm/readahead.c#L236

mmap有他自己的readahead, (感觉有些多此一举, 我知道一定是我理解不深刻), 可以用 `madvise` 避免
