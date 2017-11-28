---

layout: post
title: changetime和modifytime的区别
data: 2017-11-28 10:51:55 +0800

---

[https://stackoverflow.com/questions/3385203/regarding-access-time-unix](https://stackoverflow.com/questions/3385203/regarding-access-time-unix)

The stat(2) structure keeps track of all the file date/times:

       struct stat {
           dev_t     st_dev;     /* ID of device containing file */
           ino_t     st_ino;     /* inode number */
           mode_t    st_mode;    /* protection */
           nlink_t   st_nlink;   /* number of hard links */
           uid_t     st_uid;     /* user ID of owner */
           gid_t     st_gid;     /* group ID of owner */
           dev_t     st_rdev;    /* device ID (if special file) */
           off_t     st_size;    /* total size, in bytes */
           blksize_t st_blksize; /* blocksize for file system I/O */
           blkcnt_t  st_blocks;  /* number of 512B blocks allocated */
           time_t    st_atime;   /* time of last access */
           time_t    st_mtime;   /* time of last modification */
           time_t    st_ctime;   /* time of last status change */
       };
st_atime is the access time, updated on read(2) calls (and probably also when open(2) opens a file for reading) -- it is NOT updated when files are read via mmap(2). (Which is why I assume open(2) will mark the access time.)

st_mtime is the data modification time, either via write(2) or truncate(2) or open(2) for writing. (Again, it is NOT updated when files are written via mmap(2).)

st_ctime is the metadata modification time: when any of the other data in the struct stat gets modified.

You can change the timestamps on files with utime(2):

       struct utimbuf {
           time_t actime;       /* access time */
           time_t modtime;      /* modification time */
       };
Note you can only change the access time and (data) modification time. You can set either of these to arbitrary times, but the ctime is going to be set to the current time -- because you changed the metadata for the file.

There is no create time in this structure, so it's not possible to find out when a file was created directly from the system.

If you really need to know the create time, you can narrow it down to a range by looking at your backups -- assuming the file you're interested in has been backed up, along with its metadata.
