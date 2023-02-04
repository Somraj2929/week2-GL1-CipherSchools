package main

import (
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Data struct {
		Id           int    `json:"id"`
		Name         string `json:"name"`
		Year         int    `json:"year"`
		Color        string `json:"color"`
		PantoneValue string `json:"pantone_value"`
	} `json:"data"`
}

//rest: server-client communication
//rpc: client-client or server-server communication

func main() {
	url := "https://reqres.in/appi/products/3"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(response.Body)
	data, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))
	var responseObject Response
	json.Unmarshal(data, &responseObject)
	log.Println(responseObject)
}
