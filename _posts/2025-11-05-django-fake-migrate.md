---

date: 2025-11-05T17:34:02+0800
title: djang-fake-migrate
layout: post

---


django项目中，一个表已经在数据库中数据了，然后我又写了Model。makemigrations之后，migrate出错，报错说表已经有了。有什么好办法解决这个问题：我希望全部都走migrations，以便后续更改表结构方便。只能删了表走migrate重建吗？

`python manage.py migrate [your_app_name] --fake`
