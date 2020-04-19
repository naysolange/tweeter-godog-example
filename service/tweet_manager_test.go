package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nportas/tweeter/domain"
	"github.com/nportas/tweeter/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tm := service.NewTweeterManager()

	var tweet domain.Tweet

	user := "womenwhogo"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ := tm.PublishTweet(tweet)

	// Validation
	publishedTweets := tm.GetTweetsByUser("womenwhogo")

	assert.NotEmpty(t, publishedTweets)
	assert.Equal(t, 1, len(publishedTweets))

	validateTweet(t, publishedTweets[0], id, user, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tm := service.NewTweeterManager()

	var tweet domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tm.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err)
	assert.Equal(t, "user is required", err.Error())
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tm := service.NewTweeterManager()

	var tweet domain.Tweet

	user := "womenwhogo"
	var text string

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tm.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err)
	assert.Equal(t, "text is required", err.Error())
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	tm := service.NewTweeterManager()

	var tweet domain.Tweet

	user := "womenwhogo"
	text := `The Go project has grown considerably with over half a million users and community members 
	all over the world. To date all community oriented activities have been organized by the community
	with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tm.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err)
	assert.Equal(t, "text exceeds 140 characters", err.Error())
}
func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	tm := service.NewTweeterManager()

	var tweet, secondTweet domain.Tweet

	user := "womenwhogo"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)

	// Operation
	firstId, _ := tm.PublishTweet(tweet)
	secondId, _ := tm.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tm.GetTweetsByUser("womenwhogo")

	assert.Equal(t, 2, len(publishedTweets))

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	validateTweet(t, firstPublishedTweet, firstId, user, text)
	validateTweet(t, secondPublishedTweet, secondId, user, secondText)

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tm := service.NewTweeterManager()

	var tweet, secondTweet, thirdTweet domain.Tweet

	user := "womenwhogo"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	firstId, _ := tm.PublishTweet(tweet)
	secondId, _ := tm.PublishTweet(secondTweet)
	tm.PublishTweet(thirdTweet)

	// Operation
	tweets := tm.GetTweetsByUser(user)

	// Validation
	assert.Equal(t, 2, len(tweets))

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	validateTweet(t, firstPublishedTweet, firstId, user, text)
	validateTweet(t, secondPublishedTweet, secondId, user, secondText)

}

func validateTweet(t *testing.T, tweet domain.Tweet, id int, user, text string) {

	assert.Equal(t, id, tweet.GetId())
	assert.Equal(t, user, tweet.GetUser())
	assert.Equal(t, text, tweet.GetText())
	assert.NotNil(t, tweet.GetDate())
}
