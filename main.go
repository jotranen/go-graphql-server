package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	host := "localhost"
	port := "7777"
	pathQL := "/graphql"

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	fmt.Println("Running @ http://" + host + ":" + port + pathQL)

	log.Fatalln(r.Run(host + ":" + port))
}
