package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/trewanek/layered-architecture/presentation/handler"
	"github.com/trewanek/layered-architecture/registory"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

const versionV1 = "v1"

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

func NewRouter(ser registory.Service) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	var GetUserById http.HandlerFunc = handler.NewUserHandler(ser).GetUserById
	router.Methods(GET).Path(formatPath(versionV1, "users/{userId}")).Name("GetUserById").Handler(GetUserById)

	return router
}

func formatPath(version string, path string) string {
	return fmt.Sprintf("/%s/%s", version, path)
}
