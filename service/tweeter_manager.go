package service

import (
	"fmt"

	"github.com/nportas/tweeter/domain"
)

type TweeterManager struct {
	tweets       []domain.Tweet
	tweetsByUser map[string][]domain.Tweet
}

func NewTweeterManager() *TweeterManager {

	tm := new(TweeterManager)

	tm.tweets = make([]domain.Tweet, 0)
	tm.tweetsByUser = make(map[string][]domain.Tweet)

	return tm
}

func (tm *TweeterManager) CreateUser(username string) error {

	_, exists := tm.tweetsByUser[username]

	if exists {
		return fmt.Errorf("the user is already in use")
	}

	tm.tweetsByUser[username] = nil
	return nil
}

func (tm *TweeterManager) InitializeUser(username string) {

	tm.tweetsByUser[username] = make([]domain.Tweet, 0)
}

func (tm *TweeterManager) PublishTweet(tweetToPublish domain.Tweet) (int, error) {

	if tweetToPublish.GetUser() == "" {
		return 0, fmt.Errorf("user is required")
	}

	if tweetToPublish.GetText() == "" {
		return 0, fmt.Errorf("text is required")
	}

	if len(tweetToPublish.GetText()) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}

	tm.tweets = append(tm.tweets, tweetToPublish)

	tweetToPublish.SetId(len(tm.tweets))

	userTweets := tm.tweetsByUser[tweetToPublish.GetUser()]
	tm.tweetsByUser[tweetToPublish.GetUser()] = append(userTweets, tweetToPublish)

	return tweetToPublish.GetId(), nil
}

func (tm *TweeterManager) GetTweetsByUser(user string) []domain.Tweet {

	return tm.tweetsByUser[user]
}
