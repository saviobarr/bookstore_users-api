package users

import (
	"bookstore_users-api/domain/users"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: Handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	c.String(http.StatusNotImplemented, "Implement me!")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
