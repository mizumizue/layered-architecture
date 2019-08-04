package controller

import (
	"github.com/trewanek/LayeredArchitectureWithGolang/application"
	"github.com/trewanek/LayeredArchitectureWithGolang/presentation/view"
	"log"
	"net/http"
)

type User struct {
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	log.Printf("request: %+v", r)

	userId := GetPathParam(r.URL.Path)
	ctx := r.Context()

	appUser := application.NewUser()
	viewUser := view.NewUser()
	user, err := appUser.GetUserByID(ctx, userId)
	if err != nil {
		switch err.(type) {
		case *application.ResourceNotFoundError:
			viewUser.RenderErrorJSON(w, http.StatusNotFound, err)
			return
		case *application.InternalError:
		default:
			viewUser.RenderErrorJSON(w, http.StatusInternalServerError, err)
			return
		}
	}
	viewUser.RenderJSON(w, http.StatusOK, user)
}
