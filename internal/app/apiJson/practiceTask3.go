package apiJson

import (
	"fmt"
	cfg "github.com/zenbool/gonix/configs/apiJson"
	"io/ioutil"
	"net/http"
)


func Client(){
	answ, err := http.Get(cfg.UrlHostAPI() + cfg.Resources("posts"))
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
