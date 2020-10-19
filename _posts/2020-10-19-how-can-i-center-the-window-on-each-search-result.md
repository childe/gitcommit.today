---

date: 2020-10-19T15:10:07+0800

---

[https://vi.stackexchange.com/questions/13641/how-can-i-center-the-window-on-each-search-result](https://vi.stackexchange.com/questions/13641/how-can-i-center-the-window-on-each-search-result)

[https://stackoverflow.com/questions/45866451/put-the-search-results-at-the-top-of-the-screen-in-vi](https://stackoverflow.com/questions/45866451/put-the-search-results-at-the-top-of-the-screen-in-vi)

If you :set scrolloff=999 (or other arbitrarily high number) you will get the effect you are looking for but it doesn't turn off. In other words, the cursor line will be midway between the top and bottom of the window whether you're searching, editing, moving around in Normal mode, etc.

From :help scrolloff:

> Minimal number of screen lines to keep above and below the cursor. This will make some context visible around where you are working. If you set it to a very large value (999) the cursor line will always be in the middle of the window (except at the start or end of the file or when long lines wrap).

Try it out. Some people keep this on all the time. Maybe you'll like it, too. We might also be able to figure out a mapping/script that enables it just while searching, too, but if you don't want to mess around with a mapping with zz I imagine you don't want to mess around with a mapping for this either.
