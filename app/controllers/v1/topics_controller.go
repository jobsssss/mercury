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

func (ctrl *TopicsController) Update(ctx *gin.Context) {

	topicModel := topic.Get(ctx.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	request := requests.TopicRequest{}
	if ok := requests.Validate(ctx, &request, requests.TopicSave); !ok {
		return
	}

	topicModel.Title = request.Title
	topicModel.Body = request.Body
	topicModel.CategoryID = request.CategoryID
	rowsAffected := topicModel.Save()
	if rowsAffected > 0 {
		response.Data(ctx, topicModel)
	} else {
		response.Abort500(ctx, "更新失败，请稍后尝试~")
	}
}
