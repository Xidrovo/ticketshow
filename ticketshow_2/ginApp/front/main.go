package main

import (
	"log"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

var router *gin.Engine

func main() {
	var router = gin.Default()
	router.Use(Cors())
	gin.SetMode(gin.ReleaseMode)

	// Serve frontend static files
	const view = "./../dist"
	router.Use(static.Serve("/", static.LocalFile(view, true)))

	router.Run(":3000")
}
