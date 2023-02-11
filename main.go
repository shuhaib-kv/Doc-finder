package main

import (
	"imagedisplay/db"
	"imagedisplay/initializers"
)

func init() {
	db.COnnectDb()
	initializers.LoadEnvVariables()

}

func main() {

}
