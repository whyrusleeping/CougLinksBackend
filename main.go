package main

import (
	"net/http"
)

func StartServer(host string) {
	cl := New()

	//Set up http listeners
	http.HandleFunc("/users", cl.UserRequest)
	http.HandleFunc("/users/", cl.SingleUserRequest)
	http.Handle("/", http.FileServer(http.Dir("html")))
	http.ListenAndServe(host, nil)
}

func main() {
	StartServer(":8080")
}

/*
To update a users info:
curl 127.0.0.1:8080/users -X PUT --data '{"Action":"UPDATE", "Token":"thisismytoken", "Value":{"UUID":"11229324","Name":"Jeromy","Resume":"http://jero.my/resume.md"}}'

*/
