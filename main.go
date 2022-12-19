package main

import (
	"github.com/gin-gonic/gin"
	swagger "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"openapi/docs"
)

// @title Core App API
// @BasePath /api/v1
// @version 1.0
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization

// Helloworld
// @Summary Get list of academy articles
// @Tags Hello World
// @Accept json
// @Produce json
// @Param page query Page false "page"
// @Success 200 {object} JSONResult{data=Order} "Success Response"
// @Router /api/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

type Page struct {
	Page int `minimum:"1" default:"1"`
}

type JSONResult struct {
	Count    int         `json:"count" default:"1"`
	NextPage string      `json:"nextPage"`
	Data     interface{} `json:"data"`
}

type Order struct {
	Id   uint        `json:"id"`
	Data interface{} `json:"data"`
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	v1 := r.Group("/api")
	{
		v1.GET("/helloworld", Helloworld)
		v1.GET("/doc/*any", ginSwagger.WrapHandler(swagger.Handler))
	}
	r.Run(":8080")
}
