package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ToDo struct {
	TASK        string  `json: "task"`
	COMPLETED   bool    `json: "completed"`
	DESCRIPTION *string `json: "description"`
	ID          int32   `json: "id"`
}

func main() {
	postData()
}
func getData() {
	resp, err := http.Get("http://127.0.0.1:8080/api/todos")
	if err != nil {
		log.Fatal(err)
		return
	}
	data, derr := io.ReadAll(resp.Body)
	if derr != nil {
		log.Fatal(derr)
		return
	}
	newstr := string(data)
	log.Println(newstr)
}
func postData() {
	newToDo := ToDo{
		TASK:        "Buy milk",
		COMPLETED:   false,
		DESCRIPTION: nil,
		ID:          1900,
	}
	data, err := json.Marshal(newToDo)
	if err != nil {
		log.Fatal("error")
		return
	}
	response := bytes.NewBuffer(data)
	resp, err := http.Post("http://127.0.0.1:8080/api/todos", "application/json", response)
	if err != nil {
		log.Fatal("error")
		return
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("err")
	}
	log.Println(string(data))
}
func deleteItem(url string) {

}
