package main

import (
	"log"
	"net"

	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	pb "../proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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

func checkErr(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

const (
	port = ":50051"
)

type server struct{}

const (
	address     = "ec2-18-191-125-159.us-east-2.compute.amazonaws.com:50051" // ec2-18-218-144-18.us-east-2.compute.amazonaws.com
	defaultName = "world"
)

var router *gin.Engine

func cliGetAllEvents() ([50]pb.Evento, error) {
	var eventos [50]pb.Evento
	MYSQL_DATABASE := "ticketshowXIV"
	MYSQL_PASSWORD := "xavier"
	MYSQL_USER := "xavier"
	MYSQL_HOST := "ec2-18-191-125-159.us-east-2.compute.amazonaws.com"

	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_DATABASE)

	db, err := sql.Open("mysql", dbConfig)
	rows, err := db.Query("SELECT id, sede_id, start_time, end_time, name, creador, valor from evento")
	checkErr(err, "algo aca")
	defer db.Close()
	index := 0
	for rows.Next() {
		var id int64
		var sede_id int64
		var start_time string
		var end_time string
		var name string
		var creador int64
		var valor string
		err = rows.Scan(&id, &sede_id, &start_time, &end_time, &name, &creador, &valor)
		checkErr(err, "algo aqui")
		evt := pb.Evento{}
		evt.ID = id
		evt.Sede = sede_id
		evt.Creador = creador
		evt.Inicio = start_time
		evt.Fin = end_time
		evt.Valor = valor
		eventos[index] = evt
		index++
		if err != nil {
			panic(err.Error())
		}
	}
	return eventos, nil
}

func (s *server) GetEventos(in *pb.RequestEvento, stream pb.Micro_GetEventosServer) error {
	var eventos [50]pb.Evento
	eventos, err := cliGetAllEvents()
	checkErr(err, "get Eventos")
	for _, evento := range eventos {
		if err := stream.Send(&evento); err != nil {
			return err
		}
	}
	return nil
}
func cliGetAllSedes(c *gin.Context) {
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
	fmt.Println("Servidor corriendo")
	lis, err := net.Listen("tcp", port)
	checkErr(err, "main")
	s := grpc.NewServer()
	pb.RegisterMicroServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

