package acceptance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cucumber/messages-go/v10"
	"github.com/nportas/tweeter/domain"
	"github.com/nportas/tweeter/rest"
	"github.com/nportas/tweeter/service"
)

type AcceptanceContext struct {
	tweeterManager *service.TweeterManager
	publishedTweet domain.Tweet
}

func (ac *AcceptanceContext) InitializeContext(scenario *messages.Pickle) {

	// Create the manager
	tweeterManager := service.NewTweeterManager()
	ac.tweeterManager = tweeterManager

	// Run the server
	ginServer := rest.NewGinServer(tweeterManager)
	ginServer.StartGinServer()
}

func (ac *AcceptanceContext) theUserExists(username string) error {

	// Create body with the new user
	body := []byte(`"` + username + `"`)

	// Send POST
	http.Post("http://localhost:8080/createUser", "application/json", bytes.NewBuffer(body))

	return nil
}

func (ac *AcceptanceContext) theUserNeverHasTweeted(username string) error {

	// Create body with the username
	body := []byte(`"` + username + `"`)

	// Send PUT
	http.Post("http://localhost:8080/initializeUser", "application/json", bytes.NewBuffer(body))

	return nil
}

func (ac *AcceptanceContext) theUserSendATweet(username, text string) error {

	// Create body with a tweet
	tweet := domain.NewTextTweet(username, text)
	body, _ := json.Marshal(tweet)

	// Send POST
	http.Post("http://localhost:8080/publishTweet", "application/json", bytes.NewBuffer(body))

	return nil
}

func (ac *AcceptanceContext) theTweetIsInTimeline(username, expectedText string) error {

	// Send GET
	res, err := http.Get("http://localhost:8080/listTweets/" + username)

	// Validate result
	if err != nil {
		return fmt.Errorf("error getting tweets from user")
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error getting tweet from user")
	}

	// Validate published tweet against expected tweet
	textTweets := []domain.TextTweet{}
	body, _ := ioutil.ReadAll(res.Body)
	errUnmarshal := json.Unmarshal(body, &textTweets)

	if errUnmarshal != nil {
		return fmt.Errorf("error unmarshaling body: %v", string(body))
	}

	if len(textTweets) != 1 {
		return fmt.Errorf("the amount of tweets in the timeline user is %v but should be 1", len(textTweets))
	}

	textTweet := textTweets[0]

	if textTweet.GetText() != expectedText {
		return fmt.Errorf("the text is %v but should be %v", textTweet.GetText(), expectedText)
	}

	if textTweet.GetUser() != username {
		return fmt.Errorf("the user is %v but should be %v", textTweet.GetUser(), username)
	}

	return nil
}
