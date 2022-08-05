package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/walbety/go-user-ms/src/domain/entity"
	"github.com/walbety/go-user-ms/src/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {

	return &controller{
		service: service,
	}

}
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	return c.service.Save(video)

}
