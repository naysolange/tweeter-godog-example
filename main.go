package main

import "github.com/nportas/tweeter-godog-example/tweeter"

func main() {

	manager := tweeter.NewManager()

	ginServer := tweeter.NewGinServer(manager)
	ginServer.Start()

}
