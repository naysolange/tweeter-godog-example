package acceptance

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

func FeaturePublishTweetContext(s *godog.Suite) {

	execution := AcceptanceContext{}

	s.BeforeScenario(execution.InitializeContext)

	s.Step(`que existe el usuario (.*)$`, execution.theUserExists)
	s.Step(`que el usuario (.*) no ha twiteado nunca$`, execution.theUserNeverHasTweeted)
	s.Step(`el usuario (.*) envía el tweet con texto (.*)$`, execution.theUserSendATweet)
	s.Step(`en el timeline de (.*) aparece el tweet con texto (.*)$`, execution.theTweetIsInTimeline)
}

func TestPublishTweetAcceptance(t *testing.T) {

	format := "progress"
	for _, arg := range os.Args[1:] {
		if arg == "-test.v=true" { // go test transforms -v option
			format = "pretty"
			break
		}
	}
	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		FeaturePublishTweetContext(s)
	}, godog.Options{
		Format: format,
		//Paths:  []string{"features"},
	})

	assert.Equal(t, 0, status)
}
