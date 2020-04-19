package domain_test

import (
	"testing"

	"github.com/nportas/tweeter/domain"
)

func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("womenwhogo", "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@womenwhogo: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("womenwhogo", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@womenwhogo: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}
