package v1

import (
	"mercury/app/models/topic"
	"mercury/app/policies"
	"mercury/app/requests"
	"mercury/pkg/auth"
	"mercury/pkg/response"

	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	BaseAPIController
}

func (ctrl *TopicsController) Index(ctx *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(ctx, &request, requests.Pagination); !ok {
		return
	}

	data, pager := topic.Paginate(ctx, 10)
	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
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
		UserID:     auth.CurrentUID(ctx),
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
	if ok := policies.CanModifyTopic(ctx, topicModel); !ok {
		response.Abort403(ctx)
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

func (ctrl *TopicsController) Delete(c *gin.Context) {

	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := topicModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *TopicsController) Show(ctx *gin.Context) {
	topicModel := topic.Get(ctx.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(ctx)
		return
	}
	response.Data(ctx, topicModel)
}
