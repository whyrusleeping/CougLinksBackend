package main

import (
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	s := new(Student)
	fmt.Println("First Name")
	fmt.Scanln(&s.FirstName)
	fmt.Println("Last Name")
	fmt.Scanln(&s.LastName)
	fmt.Println("Major")
	fmt.Scanln(&s.Major)
	fmt.Println("Email")
	fmt.Scanln(&s.Email)
	fmt.Println("UUID")
	fmt.Scanln(&s.UUID)
	
	req := new(RequestData)
	req.Value = s

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(req)
	http.Post("http://127.0.0.1:8080/users", "application/json", buf)
}
