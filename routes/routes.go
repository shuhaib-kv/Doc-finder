package routes

import (
	"imagedisplay/contreoller"

	"github.com/gin-gonic/gin"
)

func Allroutes(c *gin.Engine) {
	c.POST("/addimage", contreoller.InsertANimageToDatabase)
	c.GET("/viewimage", contreoller.Viewimages)

}
