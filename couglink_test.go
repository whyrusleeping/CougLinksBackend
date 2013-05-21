package main

import (
	"testing"
	"encoding/json"
	"bytes"
	"net/http"
	"time"
)

func TestAPI(t *testing.T) {
	go StartServer(":8080")

	//Wait for server to startup
	time.Sleep(time.Second / 4)

	student := new(Student)
	student.Name = "Test Jones"
	student.Major = "Computer Science"
	student.Minor = "Math"
	student.Interests = "Cheesecake and Biking"

	req := new(RequestData)
	req.Action = "NEW"
	req.Value = student

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(req)

	_, err := http.Post("http://127.0.0.1:8080/users", "json", buf)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get("http://127.0.0.1:8080/users")
	if err != nil {
		t.Fatal(err)
	}

	var reply []*Student
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&reply)

	if err != nil {
		t.Fatal(err)
	}

	rStu := reply[0]

	if !student.Equal(rStu) {
		t.Fatal("Reply student is not the same!")
	}
}
