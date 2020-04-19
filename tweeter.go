package main

import (
	"github.com/nportas/tweeter/rest"
	"github.com/nportas/tweeter/service"
)

func main() {

	tweetManager := service.NewTweetManager()

	ginServer := rest.NewGinServer(tweetManager)
	ginServer.StartGinServer()

}
