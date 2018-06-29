package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create the Gin engine.
	g := gin.New()

	fmt.Println(g)
}
