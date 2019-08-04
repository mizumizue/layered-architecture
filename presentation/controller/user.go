package controller

import (
	"github.com/trewanek/layered-architecture/application"
	"github.com/trewanek/layered-architecture/domain/exception"
	"github.com/trewanek/layered-architecture/infrastructure/persistence/firestore"
	"github.com/trewanek/layered-architecture/presentation/view"
	"log"
	"net/http"
)

type User struct {
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	log.Printf("request: %+v", r)

	ctx := r.Context()
	userId := GetPathParam(r.URL.Path)
	appUser := application.NewUserUseCase(firestore.NewUserRepository())
	viewUser := view.NewUser()

	user, err := appUser.GetUserByID(ctx, userId)
	if err != nil {
		switch err.(type) {
		case *exception.ResourceNotFoundError:
			viewUser.RenderErrorJSON(w, http.StatusNotFound, err)
			return
		case *exception.InternalError:
		default:
			viewUser.RenderErrorJSON(w, http.StatusInternalServerError, err)
			return
		}
	}
	viewUser.RenderJSON(w, http.StatusOK, user)
}
