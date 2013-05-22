package main

import (
	"net/http"
	"bytes"
	"encoding/json"
)

func main() {
	s := NewStudent("Jeromy","Johnson","resume here", "Computer Science", "jeromyj@gmail.com", []string{"Math"}, []string{"Programming","Hacking", "Longboarding"}, []Skill{{"Go", "Advanced"}, {"C", "Advanced"}, {"C++", "Advanced"}, {"Networking", "Intermediate"}})
	req := new(RequestData)
	req.Value = s

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(req)
	http.Post("http://127.0.0.1:8080/users", "application/json", buf)
}
