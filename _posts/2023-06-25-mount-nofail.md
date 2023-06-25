---

date: 2023-06-25T15:26:21+0800
layout: post

---

External devices that are to be mounted when present but ignored if absent may require the nofail option. This prevents errors being reported at boot. For example:

```
/etc/fstab
/dev/sdg1        /media/backup    jfs    nofail,x-systemd.device-timeout=1ms    0  2
```

The nofail option is best combined with the x-systemd.device-timeout option. This is because the default device timeout is 90 seconds, so a disconnected external device with only nofail will make your boot take 90 seconds longer, unless you reconfigure the timeout as shown. Make sure not to set the timeout to 0, as this translates to infinite timeout.

[https://wiki.archlinux.org/title/fstab](https://wiki.archlinux.org/title/fstab)
