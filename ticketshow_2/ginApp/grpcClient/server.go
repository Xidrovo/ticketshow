package main

import (
	"io"
	"log"
	"net/http"

	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	pb "../proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

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

const (
	address = "ec2-18-222-131-157.us-east-2.compute.amazonaws.com:50051"
	//"localhost:50051"
	// localhost
	// 1. ec2-18-191-125-159.us-east-2.compute.amazonaws.com
	// 2. ec2-18-222-131-157.us-east-2.compute.amazonaws.com
	defaultName = "world"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

var router *gin.Engine

func getAllEvents(c *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {

		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMicroClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	stream, err := client.GetEventos(ctx, &pb.RequestEvento{})
	var names [50]gin.H
	i := 0
	for {
		evento, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
		}

		if evento != nil {
			names[i] = gin.H{"Sede": evento.Sede, "id": evento.ID, "Creador": evento.Creador, "Inicio": evento.Inicio, "Fin": evento.Fin, "Valor": evento.Valor}
		}
		i++
	}
	c.JSON(http.StatusOK, names)

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

func main() {
	var router = gin.Default()
	router.Use(Cors())
	gin.SetMode(gin.ReleaseMode)

	api := router.Group("/api")
	{
		api.GET("/sedes", getAllSedes)
		api.GET("/eventos", getAllEvents)
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	log.Println("Now server is running on port 3001!!!")

	router.Run(":3001")
}
