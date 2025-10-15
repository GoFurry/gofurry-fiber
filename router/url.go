package routers

import (
	user "github.com/GoFurry/gofurry-fiber/apps/system/user/controller"
	"github.com/gofiber/fiber/v2"
)

/*
 * @Desc: 接口层
 * @author: 福狼
 * @version: v1.0.0
 */

func userApi(g fiber.Router) {
	// 非鉴权
	g.Post("/login", user.UserApi.Login)
	// 鉴权

}
