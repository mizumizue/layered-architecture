package main

import (
	"github.com/google/martian/log"
	"github.com/trewanek/LayeredArchitectureWithGolang/presentation/controller"
	"net/http"
)

const port = ":8080"

func main() {
	log.Infof("Server started")
	router := controller.NewRouter()
	log.Errorf(http.ListenAndServe(port, router).Error())
}
