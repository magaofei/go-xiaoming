package main

import (
	"net/http"
)



func main() {
	client := http.Client{}
	resp, err := client.Get("http://127.0.0.1:8080/golang")
	if err != nil {
		print(err)
	}


	//body, _ := ioutil.ReadAll(resp.Body)
	//print(body)
	print(resp.Status)
}
