package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vikram-parashar/rss-rush/database"
)

func (dbApi *dbAPI) handleCreateChannel(c *gin.Context) {
	name := c.Query("name")
	htmlUrl := c.Query("htmlUrl")
	xmlUrl := c.Query("xmlUrl")
	if name == "" || xmlUrl == "" {
		c.JSON(400, fmt.Sprintf("empty param/s"))
		return
	}

	user, err := dbApi.AuthUser(c)
	if err != nil {
		log.Println(err)
		return
	}

	resChannel, err := dbApi.DB.CreateChannel(context.Background(), database.CreateChannelParams{
		Name: name,
		HtmlUrl: sql.NullString{
			String: htmlUrl,
			Valid:  htmlUrl != "",
		},
		XmlUrl:  xmlUrl,
		OwnerID: user.ID,
	})
	if err != nil {
		c.JSON(500, fmt.Sprintf("db insert failed: %v", err))
		return
	}

	_, err = dbApi.DB.AddFollow(context.Background(), database.AddFollowParams{
		UserID:    user.ID,
		ChannelID: resChannel.ID,
	})
	if err != nil {
		c.JSON(500, fmt.Sprintf("db insert failed: %v", err))
		return
	}

	resWithJSON(c, resChannel)
}

func (dbApi *dbAPI) handleGetChannels(c *gin.Context) {
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 32)
	if err != nil {
		c.JSON(400, fmt.Sprintf("limit query cant be parsed to number"))
		return
	}
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 32)
	if err != nil {
		c.JSON(400, fmt.Sprintf("offset query cant be parsed to number"))
		return
	}

	res, err := dbApi.DB.GetChannels(context.Background(), database.GetChannelsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		c.JSON(500, fmt.Sprintf("db fetch failed: %v", err))
		return
	}

	// Convert []Channel to []interface{}
	interfaceSlice := make([]interface{}, len(res))
	for i, ch := range res {
		interfaceSlice[i] = ch
	}
	resWithJSONArray(c, interfaceSlice)
}

func (dbApi *dbAPI) handleDeleteChannel(c *gin.Context) {
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
	err = dbApi.DB.DeleteChannel(context.Background(), database.DeleteChannelParams{
		ID:      channelId,
		OwnerID: user.ID,
	})
	if err != nil {
		c.JSON(500, fmt.Sprintf("db exec failed: %v", err))
		return
	}

	c.JSON(200, "success")
}
