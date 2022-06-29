package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sapirbismot/URLShortener/api"
)

func main() {
	fmt.Println("hello world!")

	r := gin.Default()
	r.GET("/", api.Hello)
	r.POST("/create_url", api.CreateUrl)
	r.GET("/:short_url", api.Redirect)

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failure: %v", err))
	}
}
