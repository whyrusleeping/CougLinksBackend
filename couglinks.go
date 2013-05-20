package main

import (
	"net/http"
	"encoding/json"
	"log"
)

type RequestData struct {
	Action string
	Token string
	Value *Student
}

type CougLink struct {
	students []*Student
	studentsByName map[string]*Student
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
	cl.newStudents = make(chan *Student)
	cl.updateStudent = make(chan *Student)
	cl.studentsByName = make(map[string]*Student)
	go cl.StartSyncRoutine()
	return cl
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

			//TODO: Handle any errors here
			s.userListData,_ = json.Marshal(s.students)
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
		//TODO: Handle any errors here
		s.userListData,_ = json.Marshal(s.students)
	}
}

//Respond to HTTP Requests
func (s *CougLink) UserRequest(w http.ResponseWriter, r *http.Request) {
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

		//Send newly recieved student off to the sync thread
		s.newStudents <- Req.Value
	case "PUT": //PUT Requests are for updating existing users
		//We need to somehow authenticate for this
		log.Println("Update user request!")

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
		s.updateStudent <- Req.Value
	}
}
