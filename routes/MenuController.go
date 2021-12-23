package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMenu(c *gin.Context) {
	UserId, _ := c.Get("UserId") // ログインユーザの取得
	fmt.Println("UserId", UserId)

	c.HTML(http.StatusOK, "menu", gin.H{"UserId": UserId})
}
