package rest

import (
	"net/http"

	"github.com/nportas/tweeter/domain"

	"github.com/gin-gonic/gin"
	"github.com/nportas/tweeter/service"
)

type GinTweet struct {
	User string
	Text string
	URL  string
	ID   int
}

type GinServer struct {
	tweeterManager *service.TweeterManager
}

func NewGinServer(tweeterManager *service.TweeterManager) *GinServer {
	return &GinServer{tweeterManager}
}

func (server *GinServer) StartGinServer() {

	router := gin.Default()

	router.POST("createUser", server.createUser)
	router.POST("initializeUser", server.initializeUser)
	router.POST("publishTweet", server.publishTweet)
	router.GET("/listTweets/:user", server.getTweetsByUser)

	go router.Run()
}

func (server *GinServer) createUser(c *gin.Context) {

	var username string
	c.Bind(&username)

	err := server.tweeterManager.CreateUser(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error creating user "+err.Error())
	} else {
		c.JSON(http.StatusOK, "The user was created ok")
	}
}

func (server *GinServer) initializeUser(c *gin.Context) {

	var username string
	c.Bind(&username)

	server.tweeterManager.InitializeUser(username)

	c.JSON(http.StatusOK, "The user was initialized ok")
}

func (server *GinServer) getTweetsByUser(c *gin.Context) {

	user := c.Param("user")
	c.JSON(http.StatusOK, server.tweeterManager.GetTweetsByUser(user))
}

func (server *GinServer) publishTweet(c *gin.Context) {

	var tweetdata GinTweet
	c.Bind(&tweetdata)

	tweetToPublish := domain.NewTextTweet(tweetdata.User, tweetdata.Text)

	id, err := server.tweeterManager.PublishTweet(tweetToPublish)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error publishing tweet "+err.Error())
	} else {
		c.JSON(http.StatusOK, struct{ Id int }{id})
	}
}
