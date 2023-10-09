package requests

import (
	"fmt"
	"mercury/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(ctx *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := ctx.ShouldBind(obj); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"error":   err.Error(),
		})
		fmt.Println(err.Error())
		return false
	}

	// 2. 表单验证
	errs := handler(obj, ctx)

	// 3. 判断验证是否通过
	if len(errs) > 0 {
		response.ValidationError(ctx, errs)
		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData, msg govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      msg,
	}

	return govalidator.New(opts).ValidateStruct()
}

func validateFile(ctx *gin.Context, data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Request:       ctx.Request,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}
	// 调用 govalidator 的 Validate 方法来验证文件
	return govalidator.New(opts).Validate()
}
