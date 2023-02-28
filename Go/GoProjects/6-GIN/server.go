package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/nisarg/6/controller"
	"gitlab.com/nisarg/6/middlewares"
	"gitlab.com/nisarg/6/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	// DEFAULT ...THERE ARE 2 MIDDLEWARES Recovery & Logger
	// server := gin.Default()

	server := gin.New()
	// server.Use(gin.Recovery(), gin.Logger())
	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		m := make(map[string]any)
		m["a"] = 5
		ctx.Keys = m
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		}
	})

	server.Run(":8080")

}
