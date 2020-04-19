package tweeter

import (
	"fmt"
	"time"
)

type Tweet interface {
	GetUser() string
	GetText() string
	GetDate() *time.Time
	GetId() int
	SetId(int)
	PrintableTweet() string
}

type TextTweet struct {
	User string
	Text string
	Date *time.Time
	Id   int
}

func NewTextTweet(user, text string) *TextTweet {

	date := time.Now()

	tweet := TextTweet{
		User: user,
		Text: text,
		Date: &date,
	}

	return &tweet
}

func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
}

func (tweet *TextTweet) String() string {
	return tweet.PrintableTweet()
}
