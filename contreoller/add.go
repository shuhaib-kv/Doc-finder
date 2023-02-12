package contreoller

import (
	"fmt"
	"imagedisplay/db"
	"imagedisplay/models"
	"net/http"
	"os"
	"path/filepath"

	"github.com/asticode/go-astitesseract"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/otiai10/gosseract"
)

func Add(c *gin.Context) {
	t, _ := astitesseract.New(astitesseract.Options{Languages: []string{"eng"}})
	defer t.Close()
	fmt.Println(t.GetUTF8Text("./testdata/test.png"))
	Image1, _ := c.FormFile("image")
	extension := filepath.Ext(Image1.Filename)
	img1 := uuid.New().String() + extension
	c.SaveUploadedFile(Image1, "./public/images"+img1)

	client := gosseract.NewClient()
	client.Languages = []string{"eng"}

	os.Setenv("TESSDATA_PREFIX", "/usr/share/tesseract-ocr/4.00/tessdata")
	defer client.Close()
	client.SetImage("./public/images" + img1)
	text, err := client.Text()
	if err != nil {
		fmt.Println(err.Error())
	}
	uplode := models.Document{
		FileData: img1,
		Text:     text,
	}
	db.DBS.Create(&uplode)
	c.JSON(http.StatusOK, gin.H{
		"text": text,
	})

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

func FindWord(c *gin.Context) {
	var body struct {
		word string
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}

	var Document models.Document
	db.DBS.Where("text LIKE ?", "%"+body.word+"%").First(&Document)
	if Document.Text != "" {
		c.JSON(http.StatusCreated, gin.H{
			"status":  true,
			"message": "word fount",
			"data":    Document.FileData,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status":  false,
			"message": "word not found",
			"data":    "",
		})
	}
}
