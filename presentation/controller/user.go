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
	userId := GetPathParam(r.URL.Path)
	ctx := r.Context()

	appUser := application.NewUser()
	viewUser := view.NewUser()
	user, err := appUser.GetUserByID(ctx, userId)
	if err != nil {
		if err = viewUser.RenderErrorJSON(w, http.StatusInternalServerError, err); err != nil {
			log.Printf("render response json failed: %v", err)
		}
		return
	}
	viewUser.RenderJSON(w, http.StatusOK, user)
}
