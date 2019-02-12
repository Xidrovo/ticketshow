// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	protoBuff "../grpc"
// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/net/context"
// 	"google.golang.org/grpc"
// )

// const (
// 	address     = "ec2-18-191-125-159.us-east-2.compute.amazonaws.com:50051" // ec2-18-218-144-18.us-east-2.compute.amazonaws.com
// 	defaultName = "world"
// )

// func getAllEvents(db *sql.DB) []Evento {
// 	Qs := fmt.Sprintf("SELECT name, id, sede_id, creador from evento")
// 	rows, err := db.Query(Qs)
// 	if err != nil {
// 		panic("Err")
// 	}
// 	defer rows.Close()
// 	retVal := []Evento{}

// 	for rows.Next() {
// 		evento := Evento{}
// 		err = rows.Scan(&evento.nombre, &evento.ID, &evento.sede, &evento.creador)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		retVal = append(retVal, evento)
// 	}
// 	return retVal
// }

// func getAllSedes(c *gin.Context) []Sede {
// 	db, err := sql.Open("mysql", "xidrovov:Jamytafy95@tcp(localhost:3306)/ticketshowXIV")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	Qs := fmt.Sprintf("SELECT * from Sede")
// 	rows, err := db.Query(Qs)
// 	if err != nil {
// 		panic("Errrhm")
// 	}
// 	defer rows.Close()
// 	retVal := []Sede{}

// 	for rows.Next() {
// 		sede := Sede{}
// 		err = rows.Scan(&sede.ID, &sede.Nombre, &sede.Lugar, &sede.Capacidad, &sede.Descripcion, &sede.Imagen, &sede.Creador, &sede.Ciudad)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		retVal = append(retVal, sede)
// 	}
// 	return retVal
// }

// func getPin(c *gin.Context) {
// 	conn, err := grpc.Dial(address, grpc.WithInsecure())
// 	if err != nil {

// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	client := protoBuff.NewMicroClient(conn)

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
// 	defer cancel()
// 	// stream, err := client.Ping(ctx, &protoBuff.PingRequest{})
// 	// var names [50]gin.H
// 	// i := 0
// 	// for {
// 	// 	evento, err := stream.Recv()
// 	// 	if err == io.EOF {
// 	// 		break
// 	// 	}
// 	// 	if err != nil {
// 	// 		log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
// 	// 	}

// 	// 	if evento.Nombre != "" {
// 	// 		names[i] = gin.H{"nombre": evento.Nombre, "id": evento.Id, "fechaCreacion": evento.FechaCreacion, "tipoLocalidad": evento.TipoLocalidad, "localidad_id": evento.LocalidadId, "descripcion": evento.Descripcion}
// 	// 	}
// 	// 	// log.Printf("Server: %s", evento.Nombre)
// 	// 	i++
// 	// }
// 	// c.JSON(http.StatusOK, names)

// 	message, err := client.Ping(ctx, &protoBuff.PingRequest{Message: "hola"})
// 	end := time.Now()
// 	start := time.Now()
// 	fmt.Println(end.Sub(start))
// 	log.Println(message)
// 	if message == nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"not connected": "no conectado"})
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"eventos": "Los eventos"})
// 	}

// }

// // func GetAllSedes2(c *gin.Context) {
// // 	var sedes []Sede
// // 	_, err := Db.Find(&sedes)
// // 	if err != nil {
// // 		c.JSON(200, sedes)
// // 	} else {
// // 		panic(err.Error())
// // 		c.JSON(400, gin.H{"error": "something went wrong!!!"})
// // 	}
// // }

// // conn, err := grpc.Dial(address, grpc.WithInsecure())
// // 	if err != nil {
// // 		log.Fatalf("did not connect: %v", err)
// // 	}
// // 	defer conn.Close()
// // 	client := pb.NewMicroClient(conn)

// // 	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
// // 	defer cancel()

// //     // start := time.Now()
// //     // name := defaultName
// //     stream, err := client.GetEventos(ctx, &pb.RequestEvento{})
// //     var names [50]gin.H
// //     i := 0
// //     for {
// // 		evento, err := stream.Recv()
// // 		if err == io.EOF {
// // 			break
// // 		}
// // 		if err != nil {
// // 			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
// //         }

// //         if evento.Nombre != "" {
// //             names[i] = gin.H{"nombre": evento.Nombre,"id": evento.Id, "fechaCreacion": evento.FechaCreacion, "tipoLocalidad": evento.TipoLocalidad, "localidad_id": evento.LocalidadId, "descripcion": evento.Descripcion }
// //         }
// //         // log.Printf("Server: %s", evento.Nombre)
// //         i++
// //     }
// //     c.JSON(http.StatusOK, names)
