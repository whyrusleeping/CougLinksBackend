package main

import (
	"net/http"
	"bytes"
	"fmt"
	"os"
	"encoding/json"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must specify user to delete!")
		return
	}
	
	var delreq RequestData
	delreq.Value = new(Student)
	delreq.Value.UUID = os.Args[1]

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(delreq)

	req, err := http.NewRequest("DELETE", "http://127.0.0.1:8080/users", buf)

	if err != nil {
		panic(err)
	}

	http.DefaultClient.Do(req)
}
