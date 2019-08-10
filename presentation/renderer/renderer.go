package renderer

import (
	"encoding/json"
	"fmt"
	"github.com/trewanek/layered-architecture/domain/exception"
	"log"
	"net/http"
)

func RenderJSONResult(w http.ResponseWriter, i interface{}, err error) {
	if err != nil {
		switch err.(type) {
		case *exception.ResourceNotFoundError:
			renderErrorJSON(w, http.StatusNotFound, err)
			return
		case *exception.InternalError:
			renderErrorJSON(w, http.StatusInternalServerError, err)
			return
		default:
			renderErrorJSON(w, http.StatusInternalServerError, err)
			return
		}
	}
	renderJSON(w, http.StatusOK, i)
}

func renderJSON(w http.ResponseWriter, statusCode int, i interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	b, _ := json.Marshal(i)
	if _, err := w.Write(b); err != nil {
		log.Printf(err.Error())
	}
}

func renderErrorJSON(w http.ResponseWriter, statusCode int, responseErr error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write([]byte(fmt.Sprintf("error detail: %+v", responseErr))); err != nil {
		log.Printf("render response json failed: %+v", err)
	}
}
