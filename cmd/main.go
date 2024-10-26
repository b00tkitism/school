package main

import (
	"school/db"
	"school/route"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	db.RunMigrations()

	router := gin.Default()

	route.InitRoutes(router)

	router.Run(":8080")
}
