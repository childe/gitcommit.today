---

date: 2021-05-07T14:10:54+0800

---

[https://forum.osdev.org/viewtopic.php?t=22481](https://forum.osdev.org/viewtopic.php?t=22481)

```
/* The magic field should contain this.  */
#define MULTIBOOT2_HEADER_MAGIC         0xe85250d6

  ...

struct multiboot_header
{
  /* Must be MULTIBOOT_MAGIC - see above.  */
  multiboot_uint32_t magic;

  /* Feature flags.  */
  multiboot_uint32_t flags;

  /* The above fields plus this one must equal 0 mod 2^32. */
  multiboot_uint32_t checksum;

  ...
};
```
