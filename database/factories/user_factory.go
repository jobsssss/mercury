// Package factories 存放工厂方法
package factories

import (
	"mercury/app/models/user"
	"mercury/pkg/hash"
	"mercury/pkg/helpers"

	"github.com/bxcodec/faker/v3"
)

func MakeUsers(times int) []user.User {

	var objs []user.User
	start := user.User{
		Name:     "mercury",
		Email:    "jobs@testing.com",
		Phone:    "00012312312",
		Password: hash.BcryptHash("secret"),
	}
	objs = append(objs, start)
	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		}
		objs = append(objs, model)
	}

	return objs
}
