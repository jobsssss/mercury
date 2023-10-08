package v1

import (
	"mercury/app/models/topic"
	"mercury/app/requests"
	"mercury/pkg/auth"
	"mercury/pkg/response"

	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	BaseAPIController
}

func (ctrl *TopicsController) Store(ctx *gin.Context) {

	request := requests.TopicRequest{}
	if ok := requests.Validate(ctx, &request, requests.TopicSave); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserId:     auth.CurrentUID(ctx),
	}
	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(ctx, topicModel)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}
