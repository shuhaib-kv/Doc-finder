package contreoller

import (
	"imagedisplay/db"
	"imagedisplay/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/otiai10/gosseract"
)

func Add(c *gin.Context) {
	Image1, _ := c.FormFile("image")
	extension := filepath.Ext(Image1.Filename)

	img1 := uuid.New().String() + extension
	c.SaveUploadedFile(Image1, "./public/images"+img1)

	client := gosseract.NewClient()
	client.Languages = []string{"eng"}

	client.SetImage("./public/images" + img1)
	text, err := client.Text()
	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "fild to convert",
			"error":   err.Error(),
		})
		return
	}
	uplode := models.Document{
		FileData:    img1,
		Textinimage: text,
	}
	db.DBS.Create(&uplode)
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"text":   text,
	})
	defer client.Close()

}

func Viewimages(c *gin.Context) {

	var docs []models.Document
	db.DBS.Find(&docs)
	for _, i := range docs {
		c.JSON(200, gin.H{
			"id":    i.ID,
			"image": i.FileData,
		})
	}

}

func Search(c *gin.Context) {
	var body struct {
		Word string `json:"word"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}
	var count int64
	if err := db.DBS.Model(&models.Document{}).Where("text ilike ?", "%"+body.Word+"%").Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Error searching for word in database",
			"data":    nil,
		})
		return
	}

	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Word not found",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Word found",
		"data":    nil,
	})

}
