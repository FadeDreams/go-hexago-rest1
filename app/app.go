package app

import (
	"log"
	"net/http"

	"github.com/fadedreams/go_hexagonal_rest/domain"
	"github.com/fadedreams/go_hexagonal_rest/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Start() {

	router := mux.NewRouter()

	// wiring Stub
	//ch := ItemHandlers{service: service.NewItemService(domain.NewItemRepositoryStub())}

	// wiring Db
	ch := ItemHandlers{service: service.NewItemService(domain.NewItemRepositoryDb())}

	// define routes
	router.HandleFunc("/items", ch.getAllItems).Methods(http.MethodGet)
	router.HandleFunc("/items", ch.createItem).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", ch.getItemById).Methods(http.MethodGet)
	router.HandleFunc("/items/{id}", ch.updateItemById).Methods(http.MethodPut)
	router.HandleFunc("/items/{id}", ch.deleteItemById).Methods(http.MethodDelete)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, //react client
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	log.Println("Server is running on port 8080")
	// starting server
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", corsOpts.Handler(router)))

}
