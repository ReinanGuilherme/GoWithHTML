package main

import (
	"GoWithHTML/src/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	routers.SetupRouter(r)

	// Inicia o servidor na porta 8080
	r.Run(":8080")
}
