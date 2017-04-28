---

layout: post
title:  在doker中跑crontab
date:   2017-04-28 15:26:46 +0800

---

echo "0 * * * * /usr/local/bin/python /opt/bulk-reject/bulk-reject.py -l /opt/bulk-reject/bulk-reject.log --level debug --eshost http://10.0.0.100:9200 --index-pattern 'abcd-\%Y.\%m.\%d' --from-when 1h --to-when 0h --mail-user username0123 --mail-password \"fafdsaflb'\\\$baaaa#1&2\"" > /tmp/mycron && crontab /tmp/mycron && touch /opt/bulk-reject/bulk-reject.log && cron && tail -f /opt/bulk-reject/bulk-reject.log

- crontab里面%要转义
- echo 里面的引号转义
- $转义

等等吧

