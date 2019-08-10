package main

import (
	"github.com/trewanek/layered-architecture/registory"
	"github.com/trewanek/layered-architecture/router"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Server started")
	rep := registory.NewRepository()
	ser := registory.NewService(rep)
	r := router.NewRouter(ser)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	log.Println(http.ListenAndServe(port, r).Error())
}
