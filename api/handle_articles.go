package api

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vikram-parashar/rss-rush/database"
)

func (dbApi *dbAPI) handleGetArticles(c *gin.Context) {
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

	user, err := dbApi.AuthUser(c)
	if err != nil {
		log.Println(err)
		return
	}

	res, err := dbApi.DB.GetArticles(context.Background(), database.GetArticlesParams{
    UserID: user.ID,
		Limit:  int32(limit),
		Offset: int32(offset),
    
	})
	if err != nil {
		c.JSON(500, fmt.Sprintf("db fetch failed: %v", err))
		return
	}

	// Convert []Article to []interface{}
	interfaceSlice := make([]interface{}, len(res))
	for i, ch := range res {
		interfaceSlice[i] = ch
	}
	resWithJSONArray(c, interfaceSlice)
}
