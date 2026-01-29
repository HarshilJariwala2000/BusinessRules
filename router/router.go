package router

import (
	// "calculationengine/service"
	"calculationengine/service/attribute"
	storage "calculationengine/store"
	// "fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Router *gin.Engine

func parseRequest[T any](c *gin.Context) T {
	var request T
	if err := c.BindJSON(&request); err != nil {
        log.Fatalln("Failed to Parse Request")
    }
	return request
}

func Api(){
	Router = gin.Default()

	Router.POST("/v1/attribute/create", func(c *gin.Context){
		request := parseRequest[storage.Attribute](c)
		validate := validator.New()
		validationErr := validate.Struct(request)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}
		// fmt.Println(request)
		result, err := attribute.CreateAttribute(request)
		if err !=nil{
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, result)
	})


	Router.GET("/v1/health", func(c *gin.Context){
		c.JSON(200, gin.H{
			"status":"success",
			"message":"healthy",
		})
	})
}