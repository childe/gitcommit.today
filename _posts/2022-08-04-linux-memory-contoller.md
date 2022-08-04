---

date: 2022-08-04T16:17:53+0800
layout: post

---

[Memory Resource Controller - The Linux Kernel  documentation](https://www.kernel.org/doc/html/latest/admin-guide/cgroup-v1/memory.html)

Each cgroup maintains a per cgroup LRU which has the same structure as global VM. When a cgroup goes over its limit, we first try to reclaim memory from the cgroup so as to make space for the new pages that the cgroup has touched. If the reclaim is unsuccessful, an OOM routine is invoked to select and kill the bulkiest task in the cgroup. (See 10. OOM Control below.)

## OOM Control¶

memory.oom_control file is for OOM notification and other controls.

Memory cgroup implements OOM notifier using the cgroup notification API (See Control Groups). It allows to register multiple OOM notification delivery and gets notification when OOM happens.

To register a notifier, an application must:

create an eventfd using eventfd(2)

open memory.oom_control file

write string like “<event_fd> <fd of memory.oom_control>” to cgroup.event_control

The application will be notified through eventfd when OOM happens. OOM notification doesn’t work for the root cgroup.

You can disable the OOM-killer by writing “1” to memory.oom_control file, as:

#echo 1 > memory.oom_control

If OOM-killer is disabled, tasks under cgroup will hang/sleep in memory cgroup’s OOM-waitqueue when they request accountable memory.

For running them, you have to relax the memory cgroup’s OOM status by

enlarge limit or reduce usage.

To reduce usage,

kill some tasks.

move some tasks to other group with account migration.

remove some files (on tmpfs?)

Then, stopped tasks will work again.

At reading, current status of OOM is shown.

oom_kill_disable 0 or 1 (if 1, oom-killer is disabled)

under_oom 0 or 1 (if 1, the memory cgroup is under OOM, tasks may be stopped.)

oom_kill integer counter The number of processes belonging to this cgroup killed by any kind of OOM killer.
