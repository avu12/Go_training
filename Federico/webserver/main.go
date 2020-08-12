package main

import "net/http"

func main() {

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello wolrd!"))
	})
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
