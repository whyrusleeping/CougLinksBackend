package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8080/users")
	if err != nil {
		panic(err)
	}
	var users []*Student
	buf, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buf, &users)
	for _,s := range users {
		fmt.Printf("%s %s %s\n", s.FirstName, s.Major, s.Email)
	}
}
