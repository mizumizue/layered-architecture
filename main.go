package main

import (
	"fmt"
	"github.com/google/martian/log"
	"github.com/trewanek/LayeredArchitectureWithGolang/application"
)

func main() {
	appUser := application.NewUser()
	user, err := appUser.GetUserByID("documentID")
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	fmt.Println(user.GetFullName())
}
