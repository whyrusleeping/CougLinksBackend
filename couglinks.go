package main

import (
	"net/http"
	"encoding/json"
	"log"
)

type RequestData struct {
	Token string
	Value *Student
}

type CougLink struct {
	students []*Student
	studentsByUUID map[string]*Student
	userListData []byte //Precomputed JSON for user list requests `cache`
	newStudents chan *Student
	updateStudent chan *Student

	//Admin auth token
	token string
}

//Constructor for our database manager
//TODO: accept a location parameter or something for the db to be stored
func New() *CougLink {
	cl := new(CougLink)
	cl.newStudents = make(chan *Student, 16)
	cl.updateStudent = make(chan *Student, 16)
	cl.studentsByUUID = make(map[string]*Student)
	cl.userListData = []byte("[]")
	go cl.StartSyncRoutine()
	return cl
}

//Handle New and Update requests here to avoid having to lock data
func (s *CougLink) StartSyncRoutine() {
	for {
		select {
		case ns := <-s.newStudents:
			_,exists := s.studentsByUUID[ns.UUID]
			if exists {
				continue
			}

			s.studentsByUUID[ns.UUID] = ns
			s.students = append(s.students, ns)

			//TODO: Handle any errors here
			s.userListData,_ = json.Marshal(s.students)
		case us := <-s.updateStudent:
			s.UpdateStudent(us)
		}
	}
}

func (s *CougLink) UpdateStudent(us *Student) {
	stu, ok := s.studentsByUUID[us.UUID]
	if !ok {
		log.Println("this is awkward.")
		//Student doesnt actually exist somehow
		//This is really just a sanity check
		return
	}
	if stu.Update(us) {
		//TODO: Handle any errors here
		s.userListData,_ = json.Marshal(s.students)
	}
}

func (s *CougLink) SingleUserRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("single user request!")
	user := r.URL.Path[7:]
	log.Println(user)

}

//Respond to HTTP Requests
func (s *CougLink) UserRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dec := json.NewDecoder(r.Body)
	switch r.Method {
		case "GET": //GET requests get sent back a list of all users
		log.Println("User info request!")
		//TODO: check for body data to see if we should just send a given user

		w.Write(s.userListData)
		case "POST": //POST requests are for creating new users
		//We need to somehow authenticate for this
		log.Println("New User Request!")

		var Req RequestData
		err := dec.Decode(&Req)
		if err != nil {
			log.Println(err)
			return
		}

		if Req.Value == nil {
			log.Println("Invalid JSON Object!")
			return
		}
		s.newStudents <- Req.Value
		case "PUT": //PUT Requests are for updating existing users
		//We need to somehow authenticate for this
		log.Println("Update user request!")

		var Req RequestData
		err := dec.Decode(&Req)
		if err != nil {
			log.Println("Error in put handler")
			log.Println(err)
			return
		}

		if Req.Value == nil {
			log.Println("Invalid JSON Object!")
			return
		}
		s.updateStudent <- Req.Value
	case "DELETE":
		log.Println("DELETE not yet implemented.")
	}
}
