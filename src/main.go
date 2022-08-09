package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/walbety/go-user-ms/src/controller"
	"github.com/walbety/go-user-ms/src/middleware"
	"github.com/walbety/go-user-ms/src/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	// server := gin.Default()
	SetupLogOutput()

	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(),
		middleware.Logger())
	// middleware.BasicAuth()

	apiRoutes := server.Group("/api")
	{

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoService.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {

			err := VideoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid!"})
			}

		})

	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", VideoController.ShowAll)
	}

	server.Run(":8080")
}

func SetupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
