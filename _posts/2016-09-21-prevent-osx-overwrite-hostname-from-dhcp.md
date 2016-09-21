---

layout: post
title:  "阻止osx的hostname在DHCP时被更改"
date:   2016-09-21 11:41:06 +0800

---

osx设置Hostname的顺序, 参考[https://excitedcuriosity.wordpress.com/2007/08/24/mac-os-x-hostname-determination/](https://excitedcuriosity.wordpress.com/2007/08/24/mac-os-x-hostname-determination/)
07年的文章, 不知道还准不准

1. The name provided by the DHCP or BootP server for the primary IP address

2. The first name returned by a reverse DNS (address-to-name) query for the primary IP address

3. The local hostname (set in the Sharing pane of System Preferences)

4. The name localhost

阻止被更改

    sudo scutil --set HostName myhostname
