package main

import (
	"fmt"

	"github.com/dhiman1675/mastering-go/structs/basics/user"
)

func main() {
	// var currentUser *User
	// currentUser = newUser("R", "D", "r@d.com", 18)

	currentUser, err := user.New("R", "D", "r@d.com", 18)

	if err != nil {
		fmt.Println(err)
		return
	}

	// We can call a method directly on a struct if we attach it to that struct.
	currentUser.OutputUserData()
	currentUser.ChangeUserName()
	currentUser.OutputUserData()

	admin := user.NewAdmin("Password")
	admin.OutputUserData()
}
