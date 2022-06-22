package internal

import (
	"fmt"
	"main/internal/handler"
	"net/http"
)

var baseURL = "http://localhost:8000"

func DoesSomethingAndReturn5() uint {
	fmt.Println("I did something")
	return 5
}

func SetRoutes() {

	http.HandleFunc("/user", handler.GetUser())
	http.HandleFunc("/user-add", handler.AddUser())
	http.HandleFunc("/user-show", handler.ShowUser())
	http.HandleFunc("/user-update", handler.UpdateUser())
	http.HandleFunc("/user-delete", handler.DeleteUser())

	fmt.Println("Listening in Server: " + baseURL)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
