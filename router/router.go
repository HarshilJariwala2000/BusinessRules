package router

import (
	// "calculationengine/service"
	"calculationengine/models"
	"calculationengine/service/attribute"
	"calculationengine/service/formulas"
	"fmt"

	// storage "calculationengine/store"
	// "fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Router *gin.Engine

func parseRequest[T any](c *gin.Context) T {
	var request T
	if err := c.BindJSON(&request); err != nil {
		fmt.Println(err)
        log.Fatalln("Failed to Parse Request")
    }
	return request
}

func Api(){
	Router = gin.Default()

	Router.POST("/v1/attribute/create", func(c *gin.Context){
		request := parseRequest[models.CreateAttributeRequest](c)
		validate := validator.New()
		validationErr := validate.Struct(request)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}
		ctx := c.Request.Context()
		result, err := attribute.CreateAttribute(ctx, request)
		if err !=nil{
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, result)
	})

	Router.POST("/v1/formula/create", func(c *gin.Context){
		request := parseRequest[models.CreateFormulaRequest](c)
		validate := validator.New()
		validationErr := validate.Struct(request)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}
		ctx := c.Request.Context()
		result, err := formulas.CreateFormula(ctx, request)
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