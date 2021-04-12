package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//var msg string
	client()


}

func client(){
	answ, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	body, err := ioutil.ReadAll(answ.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	res := json.NewDecoder(body).Decode()
	fmt.Println(string(body))



}
