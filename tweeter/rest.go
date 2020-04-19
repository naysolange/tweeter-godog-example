package tweeter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinTweet struct {
	User string
	Text string
	URL  string
	ID   int
}

type GinServer struct {
	manager *Manager
}

func NewGinServer(manager *Manager) *GinServer {
	return &GinServer{manager}
}

func (gs *GinServer) Start() {

	router := gin.Default()

	router.POST("createUser", gs.createUser)
	router.POST("initializeUser", gs.initializeUser)
	router.POST("publishTweet", gs.publishTweet)
	router.GET("/listTweets/:user", gs.getTweetsByUser)

	go router.Run()
}

func (gs *GinServer) createUser(c *gin.Context) {

	var username string
	c.Bind(&username)

	err := gs.manager.CreateUser(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error creating user "+err.Error())
	} else {
		c.JSON(http.StatusOK, "The user was created ok")
	}
}

func (gs *GinServer) initializeUser(c *gin.Context) {

	var username string
	c.Bind(&username)

	gs.manager.InitializeUser(username)

	c.JSON(http.StatusOK, "The user was initialized ok")
}

func (gs *GinServer) getTweetsByUser(c *gin.Context) {

	user := c.Param("user")
	c.JSON(http.StatusOK, gs.manager.GetTweetsByUser(user))
}

func (gs *GinServer) publishTweet(c *gin.Context) {

	var tweetdata GinTweet
	c.Bind(&tweetdata)

	tweetToPublish := NewTextTweet(tweetdata.User, tweetdata.Text)

	id, err := gs.manager.PublishTweet(tweetToPublish)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error publishing tweet "+err.Error())
	} else {
		c.JSON(http.StatusOK, struct{ Id int }{id})
	}
}
