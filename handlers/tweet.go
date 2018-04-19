package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

// Request defines the input into the function
type Request struct {
	Text  string
	Image string //base64 encoded image
}

// Response defines the response from the function
type Response struct {
	Code    int
	Message string
}

//go:generate moq -out mocks_test.go . TwitterPoster

// TwitterPoster defines the behaviour for posting a tweet
type TwitterPoster interface {
	UploadMedia(base64String string) (media anaconda.Media, err error)
	PostTweet(status string, v url.Values) (tweet anaconda.Tweet, err error)
}

type Tweet struct {
	twitterClient TwitterPoster
}

func NewTweet(c TwitterPoster) *Tweet {
	return &Tweet{c}
}

// ServeHTTP a serverless request
func (th *Tweet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("New Tweet")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Bad Request")
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	req, err := marshalRequest(data)
	if err != nil {
		log.Println("Bad Request")
		createResponse(rw, http.StatusBadRequest, "Invalid request message")
		return
	}

	if req.Text == "" {
		log.Println("Empty message")
		createResponse(rw, http.StatusBadRequest, "Empty message")
		return
	}

	// upload the media
	vars := url.Values{}

	if req.Image != "" {
		media, err := th.twitterClient.UploadMedia(req.Image)
		if err != nil {
			log.Println("unable to upload")
			createResponse(
				rw,
				http.StatusInternalServerError,
				fmt.Sprintf("Unable to upload media: %s", err),
			)
			return
		}

		vars.Set("media_ids", media.MediaIDString)
	}

	_, err = th.twitterClient.PostTweet(req.Text, vars)
	if err != nil {
		createResponse(
			rw,
			http.StatusInternalServerError,
			fmt.Sprintf("Tweet failed to send: %s", err),
		)
		log.Println("failed to send")
		return
	}

	log.Println("Tweet sent")
	createResponse(rw, http.StatusOK, "Tweet sent")
}

func marshalRequest(req []byte) (Request, error) {
	r := Request{}
	return r, json.Unmarshal(req, &r)
}

func createResponse(rw http.ResponseWriter, code int, message string) {
	r := Response{Code: code, Message: message}

	data, _ := json.Marshal(r)

	rw.WriteHeader(code)
	rw.Write(data)
}
