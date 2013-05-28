package main

import (
	"net/http"
	"os"
	"bytes"
	"encoding/json"
	"fmt"
)


func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: login [userid] [password]")
		return
	}
	login := new(SessionData)

	login.ID = os.Args[1]
	login.Password = os.Args[2]

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(login)

	resp, err := http.Post("http://127.0.0.1:8080/login", "application/json", buf)

	if err != nil {
		panic(err)
	}

	dec := json.NewDecoder(resp.Body)
	reply := new(SessionData)

	dec.Decode(reply)

	fmt.Println(reply)
}
