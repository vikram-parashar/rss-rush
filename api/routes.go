package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vikram-parashar/rss-rush/database"
)

type dbAPI struct {
	DB *database.Queries
}

func SetupRoutes(r *gin.Engine, db *database.Queries) {
	dbApi := dbAPI{DB: db}

	// /user?name=...&email=...
	r.POST("/user", dbApi.handleCreateUser)
	// Header needed: Authorization:api_key
	r.GET("/user", dbApi.handleGetUser)

	// /channel?name=...&htmlUrl=...&xmlUrl=...
	r.POST("/channel", dbApi.handleCreateChannel)
	// /channels?limit=...&offset=...
	r.GET("/channels", dbApi.handleGetChannels)
	// /channel/:channelId
	r.DELETE("/channel/:channelId", dbApi.handleDeleteChannel)

	// /follow/:channelId
	r.POST("/follow/:channelId", dbApi.handleAddFollow)
	// /follows
	r.GET("/follows", dbApi.handleGetFollows)
	// /follow/:channelId
	r.DELETE("/follow/:channelId", dbApi.handleDeleteFollow)

	// /articles?limit=...&offset=...
	r.GET("/articles", dbApi.handleGetArticles)
}
