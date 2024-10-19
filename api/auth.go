package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vikram-parashar/rss-rush/database"
)

func (dbApi *dbAPI) AuthUser(c *gin.Context) (*database.User, error) {
	apiKey := c.GetHeader("Authorization")

	res, err := dbApi.DB.GetUser(context.Background(), apiKey)
	if err != nil {
		c.JSON(500, fmt.Sprintf("invalid api key: %v", err))
		return nil, fmt.Errorf("invalid api key: %v", err)
	}

	return &res, nil
}
