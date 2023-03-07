package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"testserver/apis"
	"testserver/db"
	"testserver/db/id_gen"
	"testserver/grpc_server"
)

const (
	MiddlewarePort = "MIDDLEWARE_PORT"
	CreationPort   = "CREATION_PORT"
)

func main() {
	address := ":" + os.Getenv(MiddlewarePort)
	creationAddress := ":" + os.Getenv(CreationPort)
	db := db.NewHumanizeDb()
	idGen := id_gen.NewULIDGenerator()

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	middlewareServer, err := apis.NewMiddlewareApiImplementation(db, idGen)
	if err != nil {
		log.Fatal(err)
	}
	creationServer := apis.NewCreationApi(db, idGen)
	fmt.Println("serving on " + address)
	waitChan := make(chan bool, 0)
	go func() {
		if err := grpc_server.ServeMiddlewareAPI(address, middlewareServer); err != nil {
			fmt.Println(err.Error())
		}
		waitChan <- true
	}()
	fmt.Println("serving on " + creationAddress)
	go func() {
		if err := grpc_server.ServeCreationAPI(creationAddress, creationServer); err != nil {
			fmt.Println(err.Error())
		}
		waitChan <- true
	}()
	<-waitChan
}
