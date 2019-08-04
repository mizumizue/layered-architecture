package view

import (
	"encoding/json"
	"fmt"
	"github.com/trewanek/layered-architecture/domain/model"
	"log"
	"net/http"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) RenderJSON(w http.ResponseWriter, statusCode int, user *model.User) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	b, _ := json.Marshal(&user)
	if _, err := w.Write(b); err != nil {
		log.Printf(err.Error())
	}
}

func (u *User) RenderErrorJSON(w http.ResponseWriter, statusCode int, responseErr error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write([]byte(fmt.Sprintf("error detail: %+v", responseErr))); err != nil {
		log.Printf("render response json failed: %+v", err)
	}
}
