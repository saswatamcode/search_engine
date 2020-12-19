package main

import (
	"search_engine/crawler"
	"search_engine/esutil"
)

const mapping = `
{
	"settings": {
		"number_of_shards": 1,
		"number_of_replicas": 0,
		"analysis": {
			"analyzer": {
				"my_analyzer": { 
					"type":"custom",
					"tokenizer":"standard",
					"filter": [
						"lowercase"
					]
				},
            	"my_stop_analyzer": { 
               		"type":"custom",
               		"tokenizer":"standard",
               		"filter": [
                  		"lowercase",
                  		"english_stop"
               		]
            	}
         	},
         	"filter": {
            	"english_stop": {
               		"type":"stop",
               		"stopwords":"_english_"
            	}
         	}
      	}
   },
	"mappings": {
		"properties": {
			"quote": {
				"properties": {
					"author": {
						"type":"keyword"
					},
					"content": {
						"type":"text",
						"analyzer":"my_analyzer", 
            			"search_analyzer":"my_stop_analyzer", 
             			"search_quote_analyzer":"my_analyzer" 
					},
					"created": {
						"type":"date"
					},
					"suggest_field": {
						"type":"completion"
					}
				}
			}
		}
	}
}`

const name = "quotes"

func main() {
	quotes := crawler.Run(50)

	client, ctx := esutil.CreateClient()
	esutil.CreateIndex(ctx, client, name, mapping)

	esutil.BulkIndexQuotes(ctx, client, "quotes_index", quotes)
}
