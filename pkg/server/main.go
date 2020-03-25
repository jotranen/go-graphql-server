package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jotranen/go-graphql-server/internal/handlers"
)

// HOST exported variables
var HOST string

// PORT exported variables
var PORT string

func init() {
	HOST = "localhost"
	PORT = "7777"
}

// Run gin server
func Run() {
	r := gin.Default()

	r.GET("/ping", handlers.Ping())

	fmt.Println("Running @ http://" + HOST + ":" + PORT)

	log.Fatalln(r.Run(HOST + ":" + PORT))
}
