package tweeter

import (
	"fmt"
)

type Manager struct {
	tweets       []Tweet
	tweetsByUser map[string][]Tweet
}

func NewManager() *Manager {

	m := new(Manager)

	m.tweets = make([]Tweet, 0)
	m.tweetsByUser = make(map[string][]Tweet)

	return m
}

func (m *Manager) CreateUser(username string) error {

	_, exists := m.tweetsByUser[username]

	if exists {
		return fmt.Errorf("the user is already in use")
	}

	m.tweetsByUser[username] = nil
	return nil
}

func (m *Manager) InitializeUser(username string) {

	m.tweetsByUser[username] = make([]Tweet, 0)
}

func (m *Manager) PublishTweet(tweetToPublish Tweet) (int, error) {

	if tweetToPublish.GetUser() == "" {
		return 0, fmt.Errorf("user is required")
	}

	if tweetToPublish.GetText() == "" {
		return 0, fmt.Errorf("text is required")
	}

	if len(tweetToPublish.GetText()) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}

	m.tweets = append(m.tweets, tweetToPublish)

	tweetToPublish.SetId(len(m.tweets))

	userTweets := m.tweetsByUser[tweetToPublish.GetUser()]
	m.tweetsByUser[tweetToPublish.GetUser()] = append(userTweets, tweetToPublish)

	return tweetToPublish.GetId(), nil
}

func (m *Manager) GetTweetsByUser(user string) []Tweet {

	return m.tweetsByUser[user]
}
