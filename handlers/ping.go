package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	code := c.Query("code")
	if len(code) == 0 || code == "0" {
		c.JSON(http.StatusOK, ok("pong"))
	} else {
		cd, _ := strconv.ParseInt(code, 10, 64)
		c.JSON(http.StatusOK, fail(withCode(int(cd))))
	}
}
