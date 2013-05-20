package main

import (
	"testing"
	"encoding/json"
	"bytes"
	"net/http"
)

func TestAPI(t *testing.T) {
	go StartServer(":8080")

	//Generate a bunch of random students and add them, request them back and validate everything
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
