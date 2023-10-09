package v1

import (
	"mercury/app/models/link"
	"mercury/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(ctx *gin.Context) {
	links := link.AllCached()
	response.Data(ctx, links)
}
