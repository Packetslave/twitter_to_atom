package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const Twitter = "http://search.twitter.com/search.json?q=from%%3A%s"

type Tweet struct {
	Id   float64
	Text string
}

type Result struct {
	Completed_in float64
	Results      []Tweet
}

func GetTweets(user string) Result {
	resp, err := http.Get(fmt.Sprintf(Twitter, user))
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var m Result
	err = json.Unmarshal(body, &m)

	if err != nil {
		log.Fatal(err)
	}

	return m
}

func main() {
	result := GetTweets("raymondh")

	for _, value := range result.Results {
		fmt.Printf("%s: %s\n", value.Id, value.Text)
	}
}
