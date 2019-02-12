package main

import (
	"log"
	"net/http"

	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	protoBuff "../grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Evento struct {
	Nombre  string `db:"name" json:name`
	ID      int    `db:"id" json:id `
	Sede    int    `db:sede_id json:sede`
	Creador int    `db:creador json:creador`
	Inicio  string `db:start_time json:start_time`
	Fin     string `db:end_time json:end_time`
	Valor   string `db:valor json:valor`
}

// type Eventos []evento

//Sede :)
type Sede struct {
	ID          int    `db:"id" json:id`
	Nombre      string `db:"name" json:"name"`
	Lugar       string `db:"place" json:"place"`
	Capacidad   int    `db:"capacity" json:"capacity"`
	Descripcion string `db:"description" json:"description"`
	Imagen      string `db:"url_img" json:"url_img"`
	Creador     int    `db:"creador" json:"creador"`
	Ciudad      string `db:"ciudad" json:"ciudad"`
}

// type Sede []sede

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

	// even := getAllSedes(db)
	api := router.Group("/api")
	{
		api.GET("/sedes", getAllSedes)
		api.GET("/pin", getPin)
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	log.Println("Now server is running on port 3000!!!")

	router.Run(":3000")
}

const (
	address     = "ec2-18-191-125-159.us-east-2.compute.amazonaws.com:50051" // ec2-18-218-144-18.us-east-2.compute.amazonaws.com
	defaultName = "world"
)

func getAllEvents(c *gin.Context) []Evento {
	db, err := sql.Open("mysql", "xavier:xavier@tcp(ec2-18-191-125-159.us-east-2.compute.amazonaws.com)/ticketshowXIV")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	Qs := fmt.Sprintf("SELECT * from evento")
	rows, err := db.Query(Qs)
	if err != nil {
		panic("Errrhm")
	}
	defer rows.Close()
	retVal := []Evento{}

	for rows.Next() {
		evento := Evento{}
		err = rows.Scan(&evento.id, &evento.sede_id, &evento.Inicio, &evento.Fin, &evento.name, &evento.creador, &evento.Valor)
		if err != nil {
			panic(err.Error())
		}
		retVal = append(retVal, evento)
	}
	c.JSON(200, retVal)
}
func getAllSedes(c *gin.Context) {
	db, err := sql.Open("mysql", "xavier:xavier@tcp(ec2-18-191-125-159.us-east-2.compute.amazonaws.com)/ticketshowXIV")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	Qs := fmt.Sprintf("SELECT * from Sede")
	rows, err := db.Query(Qs)
	if err != nil {
		panic("Errrhm")
	}
	defer rows.Close()
	retVal := []Sede{}

	for rows.Next() {
		sede := Sede{}
		err = rows.Scan(&sede.ID, &sede.Nombre, &sede.Lugar, &sede.Capacidad, &sede.Descripcion, &sede.Imagen, &sede.Creador, &sede.Ciudad)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%v\n", sede)
		retVal = append(retVal, sede)
	}
	c.JSON(200, retVal)
}

func getPin(c *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := protoBuff.NewMicroClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	message, err := client.Ping(ctx, &protoBuff.PingRequest{Message: "hola"})
	end := time.Now()
	start := time.Now()
	fmt.Println(end.Sub(start))
	log.Println(message)
	if message == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"not connected": "no conectado"})
	} else {
		c.JSON(http.StatusOK, gin.H{"eventos": "Los eventos"})
	}

}
