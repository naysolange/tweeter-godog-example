package domain_test

import (
	"testing"

	"github.com/nportas/tweeter/domain"
	"github.com/stretchr/testify/assert"
)

func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("womenwhogo", "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@womenwhogo: This is my tweet"
	assert.Equal(t, expectedText, text)
}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("womenwhogo", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@womenwhogo: This is my tweet"
	assert.Equal(t, expectedText, text)
}
