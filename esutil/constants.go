package esutil

// Mapping is the Elasticsearch index mapping
const Mapping = `
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

// QueryAll for fetching all indexed data
const QueryAll = `
{
    "match_all": {}
}`

// SearchQuery accepts string to search
const SearchQuery = `{
	"multi_match" : {
		"query" : %q,
		"fields" : ["author^100", "content^100"],
		"operator" : "and"
	}
}`
