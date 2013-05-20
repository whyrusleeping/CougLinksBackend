package main

import (
	"net/http"
	"encoding/json"
	"log"
	"io"
	"bytes"
)

type CougLink struct {
	students []*Student
	studentsByName map[string]*Student
	userListData []byte //Precomputed JSON for user list requests
	newStudents chan *Student
}

func (s *CougLink) StartSyncRoutine() {
	for {
		ns := <-s.newStudents
		_,exists := s.studentsByName[ns.Name]
		if exists {
			continue
		}

		s.students = append(s.students, ns)

		buf := new(bytes.Buffer)
		s.WriteUserList(buf)
		s.userListData = buf.Bytes()
	}
}

func (s *CougLink) WriteUserList(w io.Writer) {
	//TODO: So, eventually we want to cache this. For now, just be lazy
	enc := json.NewEncoder(w)
	enc.Encode(s.students)
}

//Respond to requests about users
func (s *CougLink) UserRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("User req!")
	w.Write(s.userListData)
}

func (s *CougLink) NewUser(w http.ResponseWriter, r *http.Request) {
	log.Println("New User Req!")
	dec := json.NewDecoder(r.Body)
	
	var newStudent Student
	dec.Decode(&newStudent)
	log.Println(newStudent)
	s.newStudents <- &newStudent
}


func (s *CougLink) ServeHtml(w http.ResponseWriter, r *http.Request) {
}


func main() {
	cl := new(CougLink)
	cl.newStudents = make(chan *Student)
	go cl.StartSyncRoutine()

	http.HandleFunc("/users", cl.UserRequest)
	http.HandleFunc("/new", cl.NewUser)
	http.HandleFunc("/update", cl.UpdateUser)
	http.Handle("/", http.FileServer(http.Dir("html")))
	http.ListenAndServe(":8080", nil)
}

/*
Example Request to add a user:
curl 127.0.0.1:8080/new --data '{"Name":"Jeromy","Resume":"None Yet","Skills":"Golang, C, Genius","Major":"Computer Science","Minor":"","Interests":"Females","Email":"me@jero.my"}'
*/
