package main

import (
	"github.com/trewanek/layered-architecture/presentation/controller"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	log.Println("Server started")
	router := controller.NewRouter()
	log.Println(http.ListenAndServe(port, router).Error())
}
