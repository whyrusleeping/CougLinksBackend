package main

import (
	"net/http"
	"bytes"
	"encoding/json"
)

func PostStudent(s *Student) {
	req := new(RequestData)
	req.Value = s

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(req)
	http.Post("http://127.0.0.1:8080/users", "application/json", buf)
}

func main() {
	s := NewStudent("11229324","Jeromy","Johnson","resume here", "Computer Science", "jeromyj@gmail.com", []string{"Math"}, []string{"Programming","Hacking", "Longboarding"}, []Skill{{"Go", "Advanced"}, {"C", "Advanced"}, {"C++", "Advanced"}, {"Networking", "Intermediate"}})
	r := NewStudent("13465245","Matt","Hintzke","doesnt need a resume", "Computer Science", "matt@awesome.com", []string{"Horticulture","Botany"}, []string{"Mountain Climbing"}, []Skill{{"Html/Css", "Advanced"}, {"C", "Mediocre"}, {"C++", "Advanced"}})
	PostStudent(s)
	PostStudent(r)
}
