package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.googleapis.com/youtube/v3/channels?part=statistics&id=UCT8FGLq9AU1H6U3ePhegxJA&key=AIzaSyBLEdmFzHwqsRg-Y_n4bGKcjBdw4AjrXjg")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
