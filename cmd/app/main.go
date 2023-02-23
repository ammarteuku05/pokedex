package main

import (
	"pokedex/controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	_ = godotenv.Load()
)

func main() {
	r := gin.Default()

	controller.Route(r)
	r.Run(":" + "3000")
}
