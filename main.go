package main

import (
	controllers "go-crud-postgres/controllers" // new
	"go-crud-postgres/models"                  // new
	"go-crud-postgres/middleware"                  // new
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := models.SetupModels() // new

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.Use(middleware.CommonMiddleware())
	r.POST("/register", controllers.RegisterUser) // register
	r.POST("/login", controllers.LoginUser) // login

	b := r.Group("/book")
	b.Use(middleware.AuthMiddleware())
	b.GET("/", controllers.FindBooks)
	b.POST("/", controllers.CreateBook) // create
	b.GET("/:id", controllers.FindBook) // find by id
	b.PATCH("/:id", controllers.UpdateBook) // update by id
	b.DELETE("/:id", controllers.DeleteBook) // delete by id

	http.ListenAndServe(":2020", r)
}
