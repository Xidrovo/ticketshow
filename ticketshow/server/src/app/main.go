package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// db, err := sql.Open("mysql", "user:password@/dbname")
type Sede struct {
	ID   int    `json:"id"`
	name string `json:"name"`
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	// Serve frontend static files
	const view = "./../../../src"
	router.Use(static.Serve("/", static.LocalFile(view, true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	db, err := sql.Open("mysql", "xidrovov:Jamytafy95@tcp(localhost:3306)/ticketshowXIV")
	// checkErr(err)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello Mars!")
	}))

	log.Println("Now server is running on port 3000")

	// Execute the query
	results, err := db.Query("SELECT id, name FROM Sede")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var Sede Sede
		// for each row, scan the result into our tag composite object
		err = results.Scan(&Sede.ID, &Sede.name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(Sede.name)
	}
	router.Run(":3000")

	// http.ListenAndServe(":3000", nil)
}

// func checkErr(err) {
//   fmt.Println(err)
//   if err != nil {
//     panic(err.Error())
//   }
// }
