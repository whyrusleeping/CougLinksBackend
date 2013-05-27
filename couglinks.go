package main

import (
	"net/http"
	"bytes"
	"encoding/json"
	"log"
)

type RequestData struct {
	Token string
	Value *Student
}

type CougLink struct {
	studentsByUUID map[string]*Student
	userListData []byte //Precomputed JSON for user list requests `cache`
	newStudents chan *Student
	updateStudent chan *Student
	deleteStudent chan *Student

	//Admin auth token
	token string
}

//Constructor for our database manager
//TODO: accept a location parameter or something for the db to be stored
func New() *CougLink {
	cl := new(CougLink)
	cl.newStudents = make(chan *Student, 16)
	cl.updateStudent = make(chan *Student, 16)
	cl.deleteStudent = make(chan *Student, 16)
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
			log.Printf("New student: %s\n", ns.FirstName)
			_,exists := s.studentsByUUID[ns.UUID]
			if exists {
				log.Println("already exists...")
				continue
			}

			s.studentsByUUID[ns.UUID] = ns

			s.UpdateUserCache()
		case us := <-s.updateStudent:
			s.UpdateStudent(us)
		case ds := <-s.deleteStudent:
			s.DeleteStudent(ds)
		//TODO: possibly add an extra channel attatched to a timer, to reload the cache every minute or so if needed
		//That way, we can save on having to reload the cache after every change to the db
		}
	}
}

func (s *CougLink) DeleteStudent(ds *Student) {
	delete(s.studentsByUUID, ds.UUID)
	s.UpdateUserCache()
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
		s.UpdateUserCache()
	}
}

func (s *CougLink) UpdateUserCache() {
	//TODO: Preallocate a buffer ? maybe.
	buf := new(bytes.Buffer)
	buf.Write([]byte("["))
	first := true
	for _, s := range s.studentsByUUID {
		if !first {
			buf.Write([]byte(","))
		} else {
			first = false
		}
		js,_ := json.Marshal(s)
		buf.Write(js)
	}
	buf.Write([]byte("]"))
	s.userListData = buf.Bytes()
}

func (s *CougLink) SingleUserRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("single user request!")
	user := r.URL.Path[7:]
	log.Println(user)
	switch r.Method {
		case "GET":
			u, ok := s.studentsByUUID[user]
			if !ok {
				w.WriteHeader(404)
				return
			}
			rs,_ := json.Marshal(u)
			w.Write(rs)
		default:
			log.Printf("Unsupported request method.\n")
	}
}

//Respond to HTTP Requests
func (s *CougLink) UserRequest(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	switch r.Method {
		case "GET": //GET requests get sent back a list of all users
		log.Println("User info request!")
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
			w.WriteHeader(400)
			return
		}
		w.WriteHeader(201)
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
		w.WriteHeader(200)
		s.updateStudent <- Req.Value
	case "DELETE":
		var Req RequestData
		err := dec.Decode(&Req)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Attempting to delete: %s.\n",Req.Value.UUID)
		s.deleteStudent <- Req.Value
	}
}
