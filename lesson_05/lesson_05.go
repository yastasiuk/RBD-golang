package main

import (
	"fmt"
	. "lesson_05/users"
)

func main() {
	usersService := InitService()

	if u, err := usersService.CreateUser(`{"name": "Jon Doe", "id": "unique-id"}`); true {
		fmt.Println("CreateUser: success", u, err)
	}

	if u, err := usersService.ListUsers(); true {
		fmt.Println("List: successsers", u, err)
	}

	if u, err := usersService.GetUser("unique-id"); true {
		fmt.Println("GetUser:beforeDelete", u, err)
	}

	if err := usersService.DeleteUser("unique-id"); true {
		fmt.Println("DeleteUser: success", err)
	}

	if u, err := usersService.GetUser("unique-id"); true {
		fmt.Println("GetUser:afterDelete", u, err)
	}

	if err := usersService.DeleteUser("not-existing"); true {
		fmt.Println("DeleteUser:error", err)
	}
}
