package main

import (
	"imagedisplay/db"
	"imagedisplay/initializers"
	"imagedisplay/routes"
	"os"

	"github.com/gin-gonic/gin"
)

var app = gin.Default()

func init() {
	db.COnnectDb()
	initializers.LoadEnvVariables()

}

func main() {
	routes.Allroutes(app)
	port := os.Getenv("PORT")
	app.Run(port)
}
