---

layout: post
title: elasticsearch script中的value values区别
date: 2017-10-27 17:12:51 +0800

---

[https://www.elastic.co/guide/en/elasticsearch/reference/5.6/modules-scripting-fields.html](https://www.elastic.co/guide/en/elasticsearch/reference/5.6/modules-scripting-fields.html)
这里说明了script中如何取文档中的值, 但是并没有说明value和values的差别.

没有找到官方的说明, 测了一下, 对于list值, values会返回所有数据, value只返回第一个值. 所以在contains时, 用value可能不会得到正确的结果. 如果不是list, 返回结果都是一样的.

```
POST lesson_v20170823/_search?q=lessonID:44
{
  "took": 5,
  "timed_out": false,
  "_shards": {
    "total": 5,
    "successful": 5,
    "failed": 0
  },
  "hits": {
    "total": 1,
    "max_score": 1,
    "hits": [
      {
        "_index": "lesson_v20170823",
        "_type": "article",
        "_id": "44",
        "_score": 1,
        "_source": {
          "lessonID": 44,
          "name": "邮轮上所有餐饮都免费吗？",
          "author": "JZY",
          "labelNames": [
            "美食",
            "第一次上船",
            "船上体验",
            "船上消费"
          ],
          "description": "除付费餐厅外，每艘邮轮上都有一个或数个主餐厅和自助餐厅，基本都是免费的，无需支付任何额外费用即可享用美食（另点的收费饮料、酒水除外）。歌诗达邮轮旗下大西洋号、赛琳娜号和维多利亚号均有1-2个主餐厅，一般晚餐分2批进行，需根据所分配到的用餐批次在指定时间前往就餐。主餐厅规模很大，且各具特色。如在歌诗达大西洋号上，跨越两层甲板的主餐厅“提香餐厅(Tiziano Restaurant)”最多可容纳1218位客人同时用餐，以意大利上世纪初的建筑风格为设计灵感，装饰有水晶吊灯和穹顶，无论偏爱中国、意大利或者其他美餐，在这里都能轻松找到最适合自己口味的地道美食。公主邮轮蓝宝石公主号有5个主餐厅，其中国际餐厅(International Dining Room)最多可容纳500余人，其余4个主餐厅均各自可容纳200余人。5个主餐厅可以更好地分散客流，用餐没有指定的批次与时间，客人可以随到随吃（少数情况需排队等候片刻），提高就餐服务体验。这5个主餐厅虽然风格各异，但菜单是相同的，餐厅门口会贴有图文并茂的当日菜单，以便客人提前知晓可以享用到的美味佳肴。如果喜欢更自由的用餐方式，可以前往自助餐厅，相对主餐厅而言，自助餐厅开放时间更长，可供选择的食品种类更多。帆船自助餐厅(Windjammer Marketplace)位于皇家加勒比海洋量子号的14层，这里就好像是一个全球美食馆，每天24小时提供不同的各国美食、新鲜果汁、现场烘培甜点供客人自由选择。",
          "labelIDs": [
            "3",
            "20",
            "22",
            "26"
          ],
          "readCount": 2145,
          "enabled": false,
          "lessonDateCreated": "2015-07-20 17:20:45",
          "dataChangeLastTime": "2017-03-31 15:00:46",
          "isRecommended": false
        }
      }
    ]
  }
}

POST lesson_v20170823/_search?q=lessonID:44
{
  "script_fields": {
    "1": {
      "script": {
        "lang": "painless",
        "inline": "doc['labelNames.keyword'].values"
      }
    }
  }
}
{
  "took": 4,
  "timed_out": false,
  "_shards": {
    "total": 5,
    "successful": 5,
    "failed": 0
  },
  "hits": {
    "total": 1,
    "max_score": 1,
    "hits": [
      {
        "_index": "lesson_v20170823",
        "_type": "article",
        "_id": "44",
        "_score": 1,
        "fields": {
          "1": [
            "第一次上船",
            "美食",
            "船上体验",
            "船上消费"
          ]
        }
      }
    ]
  }
}

POST lesson_v20170823/_search?q=lessonID:44
{
  "script_fields": {
    "1": {
      "script": {
        "lang": "painless",
        "inline":  "doc['labelNames.keyword'].value.contains('第一次')"
      }
    }
  }
}
{
  "took": 5,
  "timed_out": false,
  "_shards": {
    "total": 5,
    "successful": 5,
    "failed": 0
  },
  "hits": {
    "total": 1,
    "max_score": 1,
    "hits": [
      {
        "_index": "lesson_v20170823",
        "_type": "article",
        "_id": "44",
        "_score": 1,
        "fields": {
          "1": [
            true
          ]
        }
      }
    ]
  }
}

POST lesson_v20170823/_search?q=lessonID:44
{
  "script_fields": {
    "1": {
      "script": {
        "lang": "painless",
        "inline":  "doc['labelNames.keyword'].values.contains('第一次')"
      }
    }
  }
}
{
  "took": 9,
  "timed_out": false,
  "_shards": {
    "total": 5,
    "successful": 5,
    "failed": 0
  },
  "hits": {
    "total": 1,
    "max_score": 1,
    "hits": [
      {
        "_index": "lesson_v20170823",
        "_type": "article",
        "_id": "44",
        "_score": 1,
        "fields": {
          "1": [
            false
          ]
        }
      }
    ]
  }
}
```

