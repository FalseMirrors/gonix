package read_api_json

import (
	"fmt"
	"net/http"
)

const hostAPI string = "https://jsonplaceholder.typicode.com"

resoureces := []string{
	"/posts",
	"/comments",
	"/albums",
	"/photos",
	"/todos",
	"/users",
}


func Client(){
	answ, err := http.Get(hostAPI + resoureces[0])
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