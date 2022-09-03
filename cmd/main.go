package main

import (
	"github.com/NatSau02/main-task/internal/app/controller"
	"github.com/NatSau02/main-task/internal/app/service"
	"github.com/NatSau02/main-task/internal/optional"
	"github.com/gin-gonic/gin"
    "time"
)
var (
	rateService1    service.RateServicee       = service.New()
	rateController controller.RateController = controller.New1(rateService1)
)
func main(){
	server := gin.Default()

	server.GET("/api/story-convert", func(ctx *gin.Context) {
		ctx.JSON(200, rateController.FindAll())
	})

	server.POST("/api/create", func(ctx *gin.Context) {
		ctx.JSON(200, rateController.Save(ctx))
	})
   
    go  optional.Gorut()
   
	server.Run(":8080")
	time.Sleep(4 * time.Second) 
}