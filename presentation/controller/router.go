package controller

import (
	"encoding/json"
	"github.com/google/martian/log"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

const apiVersionV1 = "v1"

var routes = Routes{
	Route{
		"GetUserById",
		strings.ToUpper("Get"),
		"/" + apiVersionV1 + "/users/{userId}",
		GetUserById,
	},
}

func GetPathParam(r *http.Request) string {
	p := r.URL.Path
	split := strings.Split(p, "/")
	if len(split) <= 0 {
		return ""
	}
	return split[len(split)-1]
}

func WriteReponse(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("to convert interface to json is failed. ")); err != nil {
			HandleUnknownError()
		}
	}
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(b); err != nil {
		HandleUnknownError()
	}
}

func HandleUnknownError() {
	log.Errorf("unknown error")
	os.Exit(1)
}
