package view

import (
	"encoding/json"
	"fmt"
	"github.com/google/martian/log"
	"github.com/trewanek/LayeredArchitectureWithGolang/domain"
	"net/http"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) RenderJSON(w http.ResponseWriter, statusCode int, user *domain.User) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	b, _ := json.Marshal(&user)
	if _, err := w.Write(b); err != nil {
		log.Errorf(err.Error())
	}
}

func (u *User) RenderErrorJSON(w http.ResponseWriter, statusCode int, responseErr error) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write([]byte(fmt.Sprintf("error detail: %+v", responseErr))); err != nil {
		return err
	}
	return nil
}
