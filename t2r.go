package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://search.twitter.com/search.json?q=from%3Araymondh")

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	type Message struct {
		Completed_in         float64
		Max_id, Min_id, Page float64
		Query                string
		Results              []interface{}
	}

	var m Message
	err = json.Unmarshal(body, &m)

	if err != nil {
		log.Fatal(err)
	}
}
