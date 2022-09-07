---

layout: post
title:  "vim中自动补全的颜色调整"
date:   2022-09-07T14:42:03 +0800

---


exec "hi PMenu " . " guifg=#" . a:fg . " ctermfg=" . a:fg

exec "hi PMenu" . " guibg=#" . a:bg . " ctermbg=" . a:bg

exec "hi PMenu" . " gui=" . a:attr . " cterm=" . a:attr
