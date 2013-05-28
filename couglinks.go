package main

import (
	"net/http"
	"crypto/rand"
	"crypto/md5"
	"bytes"
	"encoding/json"
	"encoding/base64"
	"log"
	"time"
)

type RequestData struct {
	Token string
	Value *Student
}

type SessionData struct {
	Token string
	loginTime time.Time
	ID string
	Password string
}

type CougLink struct {
	studentsByUUID map[string]*Student
	userListData []byte //Precomputed JSON for user list requests `cache`
	newStudents chan *Student
	updateCh chan *Student
	deleteCh chan *Student

	online map[string]*SessionData

	//Admin auth token
	token string
}

//Constructor for our database manager
//TODO: accept a location parameter or something for the db to be stored
func New() *CougLink {
	cl := new(CougLink)
	cl.newStudents = make(chan *Student, 16)
	cl.updateCh = make(chan *Student, 16)
	cl.deleteCh = make(chan *Student, 16)
	cl.studentsByUUID = make(map[string]*Student)
	cl.online = make(map[string]*SessionData)
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
			//TEMPORARY UNTIL REGISTRATION IS FINISHED!
			ns.password = "PASSWORD"

			s.studentsByUUID[ns.UUID] = ns

			s.UpdateUserCache()
		case us := <-s.updateCh:
			s.UpdateStudent(us)
		case ds := <-s.deleteCh:
			s.deleteStudent(ds)
			//TODO: possibly add an extra channel attatched to a timer, to reload the cache every minute or so if needed
			//That way, we can save on having to reload the cache after every change to the db
		}
	}
}

func (s *CougLink) deleteStudent(ds *Student) {
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

func (s *CougLink) loginRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Login request received!")
	//do login
	var ses SessionData
	dec := json.NewDecoder(r.Body)
	dec.Decode(&ses)
	switch r.Method {
	case "GET":
		//returns whether or not a user is logged in
		_, ok := s.studentsByUUID[ses.ID]
		if ok {
			w.Write([]byte("true"))
		} else {
			w.Write([]byte("false"))
		}
	case "POST":
		//accepts a user ID and password and validates it for a login
		u, ok := s.studentsByUUID[ses.ID]
		if !ok {
			w.WriteHeader(400)
			return
		}
		if u.password == ses.Password {
			log.Println("Sucessful login!")
			ses.loginTime = time.Now()
			tok := MakeSessionToken(ses.ID)
			log.Println(tok)
			u.token = tok
			ses.Token = tok
			ses.Password = ""
			s.online[ses.ID] = &ses
			log.Println(ses)
			enc := json.NewEncoder(w)
			enc.Encode(ses)
		} else {
			log.Println("Login failed...")
			w.WriteHeader(400)
		}
	case "DELETE":
		//this is a logout request, must validate tokens
		u, ok := s.studentsByUUID[ses.ID]
		if !ok {
			w.WriteHeader(400)
			return
		}
		if u.token == ses.Token {
			delete(s.online, ses.ID)
			w.Write([]byte("true"))
		} else {
			w.Write([]byte("false"))
		}
	}
}

func MakeSessionToken(ID string) string {
	hsh := md5.New()
	hsh.Write([]byte(ID))
	salt := make([]byte, 32)
	rand.Read(salt)
	hsh.Write([]byte(time.Now().String()))
	hsh.Write(salt)
	final := hsh.Sum(nil)
	return base64.StdEncoding.EncodeToString(final)
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
		s.updateCh <- Req.Value
	case "DELETE":
		var Req RequestData
		err := dec.Decode(&Req)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Attempting to delete: %s.\n",Req.Value.UUID)
		s.deleteCh <- Req.Value
	}
}
