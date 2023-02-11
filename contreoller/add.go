package contreoller

import (
	"imagedisplay/db"
	"imagedisplay/models"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InsertANimageToDatabase(c *gin.Context) {

	Image1, _ := c.FormFile("image")
	extension := filepath.Ext(Image1.Filename)
	img1 := uuid.New().String() + extension
	c.SaveUploadedFile(Image1, "./public/images"+img1)
	images := models.Document{
		Image: img1,
	}
	db.Db.Create(&images)
}
func Viewimages(c *gin.Context) {

	var docs []models.Document
	db.Db.Find(&docs)
	for _, i := range docs {
		c.JSON(200, gin.H{
			"id":    i.ID,
			"image": i.Image,
		})
	}

}
