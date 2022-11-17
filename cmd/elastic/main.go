package main

import (
	"log"
	"net/http"

	"github.com/ramadoiranedar/go_elasticsearch/internal/pkg/storage/elasticsearch"
	"github.com/ramadoiranedar/go_elasticsearch/internal/post"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Bootstrap elasticsearch.
	elastic, err := elasticsearch.New([]string{"http://0.0.0.0:9200"})
	if err != nil {
		log.Fatalln(err)
	}
	if err := elastic.CreateIndex("post"); err != nil {
		log.Fatalln(err)
	}

	// Bootstrap storage.
	storage, err := elasticsearch.NewPostStorage(*elastic)
	if err != nil {
		log.Fatalln(err)
	}

	// Bootstrap API.
	postAPI := post.New(storage)

	// Bootstrap HTTP router.
	router := httprouter.New()
	router.HandlerFunc("POST", "/api/v1/posts", postAPI.Create)
	router.HandlerFunc("PATCH", "/api/v1/posts/:id", postAPI.Update)
	router.HandlerFunc("DELETE", "/api/v1/posts/:id", postAPI.Delete)
	router.HandlerFunc("GET", "/api/v1/posts/:id", postAPI.Find)

	// Start HTTP server.
	log.Fatalln(http.ListenAndServe(":3000", router))
}

// b6918b33-2d6f-4496-b119-e524cb01829d
// 486753ee-712d-4335-88d8-75a806a36705
// c99e08eb-f9e9-4b29-84bc-38ff08797423
