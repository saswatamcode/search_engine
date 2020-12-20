package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"search_engine/esutil"

	log "github.com/sirupsen/logrus"
)

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type searchReq struct {
	Query string `json:"query"`
}

func search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, ctx := esutil.CreateClient()
	indexName := os.Getenv("INDEX_NAME")
	switch r.Method {
	case "GET":
		searchRes, err := esutil.SearchIndex(ctx, client, indexName, esutil.QueryAll)
		if err != nil {
			log.WithFields(log.Fields{
				"Error": err,
			}).Error("Failed to get search results")
			resp := response{Status: "error", Message: "Failed to get search results"}
			json.NewEncoder(w).Encode(resp)
		}
		log.Info("Search successful")
		json.NewEncoder(w).Encode(searchRes)

	case "POST":
		var req searchReq
		_ = json.NewDecoder(r.Body).Decode(&req)

		searchRes, err := esutil.SearchIndex(ctx, client, indexName, fmt.Sprintf(esutil.SearchQuery, req.Query))
		log.Info(fmt.Sprintf(esutil.SearchQuery, req.Query))
		if err != nil {
			log.WithFields(log.Fields{
				"Error": err,
			}).Error("Failed to get search results")
			resp := response{Status: "error", Message: "Failed to get search results"}
			json.NewEncoder(w).Encode(resp)
		}
		log.Info("Search successful")
		json.NewEncoder(w).Encode(searchRes)
	}
}

//StartServer starts the search server
func StartServer() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	port := os.Getenv("PORT")
	http.HandleFunc("/search", search)

	log.Info("Server listening on " + port)
	error := http.ListenAndServe(":"+port, nil)
	if error != nil {
		log.WithFields(log.Fields{
			"Error": error,
		}).Error("Failed to start server")
		return
	}
}
