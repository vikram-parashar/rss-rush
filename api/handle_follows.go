package api

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vikram-parashar/rss-rush/database"
)

func (dbApi *dbAPI) handleAddFollow(c *gin.Context) {
	channelId, err := uuid.Parse(c.Param("channelId"))
	if err != nil {
		c.JSON(400, fmt.Sprintf("channelId cant be parsed: %v", err))
		return
	}

	user, err := dbApi.AuthUser(c)
	if err != nil {
		log.Println(err)
		return
	}
	res, err := dbApi.DB.AddFollow(context.Background(), database.AddFollowParams{
		UserID:    user.ID,
		ChannelID: channelId,
	})
	if err != nil {
		c.JSON(500, fmt.Sprintf("db exec failed: %v", err))
		return
	}

	resWithJSON(c, res)
}

func (dbApi *dbAPI) handleGetFollows(c *gin.Context) {
	user, err := dbApi.AuthUser(c)
	if err != nil {
		log.Println(err)
		return
	}

	res, err := dbApi.DB.GetFollows(context.Background(), user.ID)
	if err != nil {
		c.JSON(500, fmt.Sprintf("db exec failed: %v", err))
		return
	}

	// Convert []Follow to []interface{}
	interfaceSlice := make([]interface{}, len(res))
	for i, ch := range res {
		interfaceSlice[i] = ch
	}
	resWithJSONArray(c, interfaceSlice)
}

func (dbApi *dbAPI) handleDeleteFollow(c *gin.Context) {
	channelId, err := uuid.Parse(c.Param("channelId"))
	if err != nil {
		c.JSON(400, fmt.Sprintf("channelId cant be parsed: %v", err))
		return
	}

	user, err := dbApi.AuthUser(c)
	if err != nil {
		log.Println(err)
		return
	}
	err = dbApi.DB.DeleteFollow(context.Background(), database.DeleteFollowParams{
		ChannelID: channelId,
		UserID:    user.ID,
	})
	if err != nil {
		c.JSON(500, fmt.Sprintf("db exec failed: %v", err))
		return
	}

	c.JSON(200, "success")
}
