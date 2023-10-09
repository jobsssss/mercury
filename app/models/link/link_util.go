package link

import (
    "mercury/pkg/app"
    "mercury/pkg/database"
    "mercury/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (link Link) {
    database.DB.Where("id", idstr).First(&link)
    return
}

func GetBy(field, value string) (link Link) {
    database.DB.Where("? = ?", field, value).First(&link)
    return
}

func All() (links []Link) {
    database.DB.Find(&links)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Link{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(ctx *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
    paging = paginator.Paginate(
        ctx,
        database.DB.Model(Link{}),
        &links,
        app.V1URL(database.TableName(&Link{})),
        perPage,
    )
    return
}