package main


import (
	"net/http"
	"bytes"
	"encoding/json"
)

func AddDummyStudents() {
	s := NewStudent("11529424","Jeromy","Johnson","resume here", "Computer Science", "jeromyj@gmail.com", []string{"Math"}, []string{"Programming","Hacking", "Longboarding"}, []Skill{{"Go", "Advanced"}, {"C", "Advanced"}, {"C++", "Advanced"}, {"Networking", "Intermediate"}})
	r := NewStudent("13465245","Matt","Hintzke","doesnt need a resume", "Computer Science", "matt@awesome.com", []string{"Horticulture","Botany"}, []string{"Mountain Climbing"}, []Skill{{"Html/Css", "Advanced"}, {"C", "Mediocre"}, {"C++", "Advanced"}})
	n := NewStudent("12345678","Rae","Marks","Too cool for school","Computer Science","rae@awesome.com", []string{"Math"}, []string{"Rowing","Being Awesome"}, []Skill{{"C","Advanced"},{"C++","Intermediate"},{"Java","SuperAwesome"}})
	m := NewStudent("77777777", "Jimmy", "Lee","http://jimmyl.ee","Computer Science","JIMMY AT LEE DOT COM", []string{"Business"}, []string{"Drawing 'art'", "Being a Bro"}, []Skill{{"Javascript", "Uuber"},{"C", "Meh"},{"Brogramming","Over 9000"}})
	PostStudent(n)
	PostStudent(s)
	PostStudent(m)
	PostStudent(r)

}

func PostStudent(s *Student) {
	req := new(RequestData)
	req.Value = s

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(req)
	http.Post("http://127.0.0.1:8080/users", "application/json", buf)
}

