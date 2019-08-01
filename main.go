package main

import (
	"fmt"
	"github.com/trewanek/LayeredArchitectureWithGolang/application"
)

func main() {
	appUser := application.NewUser()
	fmt.Println(appUser.GetUserByID("documentID"))
}
