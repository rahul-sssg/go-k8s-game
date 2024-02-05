package main

import (
	"fmt"
	"go-k8s-game/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRoutes(router)

	router.Run(fmt.Sprintf(":%v", 7076))
}
