package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main(){
	//get will return a response and an err
	response, err := http.Get("http://api.theysaidso.com/qod")
	//basic err handling
	if err != nil{
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(contents))
}