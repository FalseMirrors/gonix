package apiJson

import (
	"fmt"
	data "github.com/zenbool/gonix/configs/apiJson"
	"io/ioutil"
	"net/http"
)

func getPosts(){
	answ, err := http.Get(data.UrlHostAPI() + data.Resources("posts"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	body, err := ioutil.ReadAll(answ.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}
