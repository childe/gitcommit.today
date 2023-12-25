---

date: 2023-12-25T18:25:33+0800
layout: post

---

[https://www.bentasker.co.uk/posts/documentation/linux/480-disk-automatically-unmounts-immediately-after-mounting.html](https://www.bentasker.co.uk/posts/documentation/linux/480-disk-automatically-unmounts-immediately-after-mounting.html)


When it happens, it's incredibly frustrating - you've had a disk replaced on a linux box, the disk has shown up with a different name in /dev, so you edit /etc/fstab and then try to mount the disk.

The command runs, without error, but the disk isn't mounted, and doesn't appear in df

This documentation details the likely cause, and how to resolve it

If you look in dmesg, you might see something like the following

```
[  462.754500] XFS (sdc): Mounting V5 Filesystem
[  462.857216] XFS (sdc): Ending clean mount
[  462.871119] XFS (sdc): Unmounting Filesystem
```

Which, whilst it shows the disk is getting unmounted almost immediately, isn't otherwise very helpful. It doesn't tell us why.

However, if you look in syslog (e.g. /var/log/messages, journalctl or /var/log/syslog) you may well see this logged again with a couple of additional relevant lines

```
kernel: XFS (sde): Mounting V5 Filesystem
kernel: XFS (sde): Ending clean mount
systemd: Unit cache2.mount is bound to inactive unit dev-sdc.device. Stopping, too.
systemd: Unmounting /cache2...
kernel: XFS (sde): Unmounting Filesystem
systemd: Unmounted /cache2.
```

We can now see that the erstwhile init system - systemd - decided to unmount the filesystem

systemd: Unit cache2.mount is bound to inactive unit dev-sdc.device. Stopping, too.
The reason for this is that at boot time systemd-fstab-generator generates, in effect, a bunch of dynamic unit files for each mount. From the output above we can tell the disk used to be sdc but is now sde. Despite fstab saying

```
/dev/sde      /mnt/cache2   xfs   defaults,nofail  0 0
```

When we issue the command

```
mount /cache2
```

SystemD picks up on the fact that it has an inactive unit file (inactive because the block device has gone away) which should be mounted to that path, decides there's a conflict, and that it knows better, and unmounts your mount again
If you're in this position, then, you should be able to briefly resolve with a single command

```
systemctl daemon-reload
```

Keep in mind that if your disk moves back following a reboot, you'll be back to this point where SystemD decides you can't have wanted to mount your disk after all.

 

SystemD have a bug for this, raised in 2015 and seemingly still unresolved (it's certainly still attracting complaints at time of writing). Rather worryingly, it suggests that the above will not always resolve the issue, and instead suggests the following "workaround"
