package app

import (
	"net/http"

	"../controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
