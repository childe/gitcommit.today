---

date: 2020-12-29T16:17:11+0800

---

kibana6.8 中使用的还是 vage-lite-2, 所以要看对应版本的文档, [https://vega.github.io/vega-lite-v2/docs/](https://vega.github.io/vega-lite-v2/docs/)

主要是使用 [calculate](https://vega.github.io/vega-lite-v2/docs/calculate.html) 

```
{
  $schema: https://vega.github.io/schema/vega-lite/v2.json
  title: CodeReview 比例

  // Define the data source
  data: {
    url: {
      index: XXX
      body: {
        aggs: {
          organization_name: {
            terms: {size: 100, field: "organization_name", min_doc_count: 0}
            aggs: {
              cr: {
                filters: {
                  filters: {
                    CR: {
                      query_string: {query: "is_code_review:true"}
                    }
                    noCR: {
                      query_string: {query: "-is_code_review:true"}
                    }
                  }
                }
              }
            }
          }
        }
        size: 0
      }
    }
    format: {property: "aggregations.organization_name.buckets"}
  }
  mark: bar

 "transform": [
    {"calculate": "round(datum.cr.buckets.CR.doc_count / datum.doc_count*100)/100", "as": "ratio"},
    {"calculate": "datum.key", "as": "bu"}
  ],
  "encoding": {
    "x": {
      "field": "bu",
      "type": "ordinal",
      "title": "BU",
      "sort": {"op": "mean", "field": "ratio", "order":"descending"}
      },
    "y": {
      "field": "ratio", "type": "quantitative", "title": "CR比例"
    },
    "tooltip": [
      {"field": "bu", "type": "ordinal"},
      {"field": "ratio", "type": "quantitative"}
    ]
  }
}
```
