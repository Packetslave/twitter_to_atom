package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const Twitter = "http://search.twitter.com/search.json?q=from%%3A%s"

type Tweet struct {
	Id_Str     string
	Text       string
	Created_At string
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
	if len(os.Args) < 2 {
		log.Fatal("usage: t2s username")
	}
	u := os.Args[1]

	result := GetTweets(u)
	pubTime := time.Now()

	feed := &Feed{
		Title:   "Twitter Feed for " + u,
		Link:    "http://twitter.com/" + u,
		PubDate: pubTime,
	}

	for _, e := range result.Results {
		entryTime, err := time.Parse(time.RFC1123Z, e.Created_At)

		if err != nil {
			log.Printf("failed to parsed date %q\n", e.Created_At)
			continue
		}

		e := &Entry{
			Title:       e.Text,
			Link:        fmt.Sprintf("http://twitter.com/%s/status/%s", u, e.Id_Str),
			Description: e.Text,
			PubDate:     entryTime,
		}
		feed.AddEntry(e)
	}
	if s, err := feed.GenXml(); err == nil {
		fmt.Printf(s)
	}
}
