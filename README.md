# Twitter
This service is responsible for tweeting selfies

## Running
To run the service obtain your credentials from twitter and run either using your favorite scheduler or with docker as illustrated below

```bash
docker run \
  -p 8080:8080 
  -e "twitter_consumer_key=abcxyz" \
  -e "twitter_consumer_secret=abcxyz" \
  -e "twitter_access_token=abcxyz" \
  -e "twitter_access_token_secret=abcxyz" \
  quay.io/selfidrone/tweets \
  --port 8080
```

## Interface

### POST /tweet
Tweet endpoint is used for sending tweets to twitter, to post an image a base64 encoded jpeg or png must be included in the request.

#### Request
```json
{
  "text": "This is the message to tweet",
  "image": "abc123=" // base64 encoded image
}
```

#### Response
```json
{
  "code": 200,
  "message": "tweet sent"
}
```

### GET /health
Health endpoint returns 200 OK if status is functioning correctly
