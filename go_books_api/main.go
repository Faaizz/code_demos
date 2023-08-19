package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"github.com/faaizz/code_demos/go_books_api/controller"
	"github.com/faaizz/code_demos/go_books_api/handle"
	"github.com/faaizz/code_demos/go_books_api/middleware"
	"github.com/faaizz/code_demos/go_books_api/model"
)

func main() {
	log.Println("Initializing...")
	db, err := controller.SetupDB()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Book{})
	controller.DB = db

	bc := controller.Book{}
	controller.BC = &bc

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := httprouter.New()
	apiDocsPath := "/api-docs"
	router.GET(apiDocsPath, handle.ApiDocs)

	apiPath := "/api/v1"

	healthPath := apiPath + "/healthz"
	router.GET(healthPath, handle.Health)

	bookPath := apiPath + "/book"
	router.GET(bookPath, handle.BookIndex)
	router.POST(bookPath, middleware.BasicAuth(handle.BookCreate))
	router.PUT(bookPath+"/:id", middleware.BasicAuth(handle.BookUpdate))
	router.DELETE(bookPath+"/:id", middleware.BasicAuth(handle.BookDelete))
	router.GET(bookPath+"/:id", middleware.BasicAuth(handle.BookRead))

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
