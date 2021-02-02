---

date: 2021-02-01T20:50:38+0800

---

[https://unix.stackexchange.com/questions/156897/why-does-du-and-quota-results-not-match](https://unix.stackexchange.com/questions/156897/why-does-du-and-quota-results-not-match)

I believe there may be some files still held open by some processes. You can try to list them using,

lsof | grep username | grep deleted
A more better version would be to use,

lsof +L1 | grep username

