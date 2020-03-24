package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jotranen/go-graphql-server/internal/handlers"
)

func main() {
	host := "localhost"
	port := "7777"
	pathQL := "/graphql"

	r := gin.Default()

	r.GET("/ping", handlers.Ping())

	fmt.Println("Running @ http://" + host + ":" + port + pathQL)

	log.Fatalln(r.Run(host + ":" + port))
}
