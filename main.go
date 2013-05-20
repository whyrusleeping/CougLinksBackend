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
	updateStudent chan *Student
}

func (s *CougLink) StartSyncRoutine() {
	for {
		select {
		case ns := <-s.newStudents:
			_,exists := s.studentsByName[ns.Name]
			if exists {
				continue
			}

			s.studentsByName[ns.Name] = ns
			s.students = append(s.students, ns)

			buf := new(bytes.Buffer)
			s.WriteUserList(buf)
			s.userListData = buf.Bytes()
		case us := <-s.updateStudent:
			s.UpdateStudent(us)
		}
	}
}

func (s *CougLink) UpdateStudent(us *Student) {
	stu, ok := s.studentsByName[us.Name]
	if !ok {
		log.Println("this is awkward.")
		//Student doesnt actually exist somehow
		//This is really just a sanity check
		return
	}
	if stu.Update(us) {
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
	switch r.Method {
	case "GET":
		log.Println("User req!")
		w.Write(s.userListData)
	case "POST":
		//We need to somehow authenticate for this
		log.Println("New User Req!")
		dec := json.NewDecoder(r.Body)

		var newStudent Student
		dec.Decode(&newStudent)
		log.Println(newStudent)
		s.newStudents <- &newStudent
	case "PUT":
		//We need to somehow authenticate for this
		log.Println("Update user request!")
		dec := json.NewDecoder(r.Body)

		var studentInfo Student
		dec.Decode(&studentInfo)
		log.Println(studentInfo)
		s.updateStudent <- &studentInfo
	}
}

func main() {
	cl := new(CougLink)
	cl.newStudents = make(chan *Student)
	cl.updateStudent = make(chan *Student)
	cl.studentsByName = make(map[string]*Student)
	go cl.StartSyncRoutine()

	http.HandleFunc("/users", cl.UserRequest)
	http.Handle("/", http.FileServer(http.Dir("html")))
	http.ListenAndServe(":8080", nil)
}

/*
Example Request to add a user:
curl 127.0.0.1:8080/new --data '{"Name":"Jeromy","Resume":"None Yet","Skills":"Golang, C, Genius","Major":"Computer Science","Minor":"","Interests":"Females","Email":"me@jero.my"}'
*/
