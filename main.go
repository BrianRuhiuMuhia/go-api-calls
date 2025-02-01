package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

var newToDo = ToDo{
	TASK:        "Buy milk",
	COMPLETED:   false,
	DESCRIPTION: nil,
	ID:          1900,
}

func main() {
	postData(newToDo)
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
func postData(item ToDo) {

	data, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
		return
	}
	response := bytes.NewBuffer(data)
	resp, err := http.Post("http://127.0.0.1:8080/api/todos", "application/json", response)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("err")
	}
	log.Println(string(data))
}
func deleteItem(url string, id int) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%d", url, id), nil)
	if err != nil {
		log.Fatal("There was an error")
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("there was an error")
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("there was an error")
	}
	log.Println(string(data))
}
func updateItem(url string, item ToDo, id int) {
	sendData, err := json.Marshal(item)
	if err != nil {
		log.Fatal("conversion error")
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%d", url, id), bytes.NewReader(sendData))
	if err != nil {
		log.Fatal("There was an error")
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("There was an error")
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	fmt.Println(string(data))

}
