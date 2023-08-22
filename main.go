package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saeedmaldosary/Go-Authentication-JWT/controllers"
	"github.com/saeedmaldosary/Go-Authentication-JWT/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.Run()
}
