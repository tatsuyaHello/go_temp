package handler

import "github.com/gin-gonic/gin"

func PingJson(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "good",
	})
	// ちなみに、このreturnは明示的に返してやっているだけであり、別になくても大丈夫そう
	return
}

func PingString(c *gin.Context) {
	c.String(200, "Hello World!!!")
	return
}
