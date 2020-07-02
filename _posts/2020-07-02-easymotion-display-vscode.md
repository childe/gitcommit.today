---

date: 2020-07-02T21:11:14+0800

---

[https://github.com/VSCodeVim/Vim/issues/2938](https://github.com/VSCodeVim/Vim/issues/2938)

> I was having the same problem, but I found a workaround. Mine had to do with my line-height setting. "editor.lineHeight": 30,. Updating the "vim.easymotionMarkerHeight": 30, to match made it better for me. You're using custom fonts and things and might be different, but this may help a little.
