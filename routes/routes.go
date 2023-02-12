package routes

import (
	"imagedisplay/contreoller"

	"github.com/gin-gonic/gin"
)

func Allroutes(c *gin.Engine) {
	c.POST("/add", contreoller.Add)
	c.GET("/view", contreoller.Viewimages)
	c.GET("/search", contreoller.Search)

}
