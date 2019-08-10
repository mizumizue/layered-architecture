package handler

import (
	"github.com/trewanek/layered-architecture/application"
	"github.com/trewanek/layered-architecture/presentation/renderer"
	"github.com/trewanek/layered-architecture/registory"
	"log"
	"net/http"
	"strings"
)

type UserHandler struct {
	ser registory.Service
}

func NewUserHandler(ser registory.Service) *UserHandler {
	return &UserHandler{ser: ser}
}

func (u *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	log.Printf("request: %+v", r)
	ctx := r.Context()
	userId := getPathParam(r.URL.Path)
	use := application.NewUserUseCase(u.ser.NewUserService())
	user, err := use.GetUserByID(ctx, userId)
	renderer.RenderJSONResult(w, user, err)
}

func getPathParam(path string) string {
	split := strings.Split(path, "/")
	if len(split) <= 0 {
		return ""
	}
	return split[len(split)-1]
}
