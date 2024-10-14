package main

import (
	"github.com/RPW-11/inventory_management_be/api/route"
	"github.com/RPW-11/inventory_management_be/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	server := gin.Default()

	// set up the routes
	route.Setup(env, app.Db, server)

	server.Run(env.ServerAddress)
}
