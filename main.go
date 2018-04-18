package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/selfidrone/twitter/handlers"
)

var port = flag.String("port", "8080", "HTTP port to listen on")

var version = "notset"

func main() {
	flag.Parse()

	c := setupDeps()
	tweetHandler := handlers.NewTweet(c)

	http.HandleFunc("/health", health)
	http.Handle("/tweet", tweetHandler)

	log.Println(http.ListenAndServe(":"+*port, nil))
}

func setupDeps() handlers.TwitterPoster {
	consumerToken := os.Getenv("twitter_consumer_key")
	consumerSecret := os.Getenv("twitter_consumer_secret")
	accessToken := os.Getenv("twitter_access_token")
	accessTokenSecret := os.Getenv("twitter_access_token_secret")

	anaconda.SetConsumerKey(consumerToken)
	anaconda.SetConsumerSecret(consumerSecret)

	return anaconda.NewTwitterApi(accessToken, accessTokenSecret)
}

func health(rw http.ResponseWriter, r *http.Request) {
	r.Body.Close()
	rw.WriteHeader(http.StatusOK)
}
