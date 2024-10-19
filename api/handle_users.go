package api

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/vikram-parashar/rss-rush/database"
)

func isValidEmail(email string) bool {
	// Regular expression for validating an email address
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func (dbApi *dbAPI) handleCreateUser(c *gin.Context) {
	name := c.Query("name")
	email := c.Query("email")
	if name == "" || email == "" {
		c.JSON(400, fmt.Sprintf("empty params name=%v, email=%v", name, email))
		return
	}
	if isValidEmail(email) == true {
		c.JSON(400, "Email not valid")
		return
	}

	res, err := dbApi.DB.CreateUser(context.Background(), database.CreateUserParams{
		Name: name,
		Email:email,
	})
	if err != nil {
		c.JSON(500, fmt.Sprintf("db insert failed: %v", err))
		return
	}

	resWithJSON(c, res)
}

func (dbApi *dbAPI) handleGetUser(c *gin.Context) {
	res, err := dbApi.AuthUser(c)
	if err != nil {
		log.Println(err)
		return
	}

	resWithJSON(c, *res)
}
