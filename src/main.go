package main

import (
	"github.com/gin-gonic/gin"
	"github.com/walbety/go-user-ms/src/controller"
	"github.com/walbety/go-user-ms/src/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoService.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run(":8080")
}
