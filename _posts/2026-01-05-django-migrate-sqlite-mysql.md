---

date: 2026-01-05T14:20:54+0800
title: 'django migrate：同时兼容 SQLite 和 MySQL 的表重命名'
layout: post

---

```python
def rename_tables_forward(apps, schema_editor):
    """根据数据库类型执行不同的重命名表操作"""
    vendor = schema_editor.connection.vendor
    with schema_editor.connection.cursor() as cursor:
        if vendor == "mysql":
            cursor.execute("RENAME TABLE foo TO bar")
        elif vendor == "sqlite":
            cursor.execute("ALTER TABLE foo RENAME TO bar")
        else:
            raise ValueError(f"Unsupported database vendor: {vendor}")


def rename_tables_reverse(apps, schema_editor):
    """反向操作：恢复表名"""
    vendor = schema_editor.connection.vendor
    with schema_editor.connection.cursor() as cursor:
        if vendor == "mysql":
            cursor.execute("RENAME TABLE bar TO foo")
        elif vendor == "sqlite":
            cursor.execute("ALTER TABLE foo bar TO foo")
        else:
            raise ValueError(f"Unsupported database vendor: {vendor}")


class Migration(migrations.Migration):
    operations = [
        # 使用 RunPython 重命名数据库表（支持 MySQL 和 SQLite）
        migrations.RunPython(
            code=rename_tables_forward,
            reverse_code=rename_tables_reverse,
        )
    ...
```
