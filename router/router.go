package routers

/*
 * @Desc: 路由层
 * @author: 福狼
 * @version: v1.0.0
 */

import (
	"sync"

	"github.com/GoFurry/gofurry-fiber/common"
	"github.com/GoFurry/gofurry-fiber/middleware"
	"github.com/GoFurry/gofurry-fiber/roof/env"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var Router *router

type router struct{}

func NewRouter() *router {
	return &router{}
}

func init() {
	Router = NewRouter()
}

var once = sync.Once{}

func (router *router) Init() *fiber.App {
	once.Do(func() {
	})

	app := fiber.New(fiber.Config{
		Network:      fiber.NetworkTCP6, // tcp tcp4 tcp6 三种模式
		AppName:      common.COMMON_PROJECT_NAME,
		ServerHeader: "GoFurry-Cloud",
		Prefork:      false, // 多核cpu处理高并发 业务量小需关闭
		// 在生产环境禁用错误堆栈跟踪
		EnablePrintRoutes: env.GetServerConfig().Server.Mode == "debug",
		// 配置默认404处理
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// 检查是否是404错误
			if err == fiber.ErrNotFound {
				return common.NewResponse(c).Error("链接不存在")
			}
			// 检查是否是405错误
			if err == fiber.ErrMethodNotAllowed {
				return common.NewResponse(c).Error("方法不存在")
			}
			// 其他错误
			return common.NewResponse(c).Error(err.Error())
		},
	})

	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}
	// 中间件
	if env.GetServerConfig().Middleware.Swagger == "on" {
		app.Use(swagger.New(cfg)) // 访问路径类似 https://[::1]:9999/swagger
	}
	// 调试模式下启用pprof
	if env.GetServerConfig().Server.Mode == "debug" {
		app.Use(pprof.New())
	}
	// 跨域
	app.Use(cors.New())
	// 恢复
	app.Use(recover.New())
	if env.GetServerConfig().Waf.WafSwitch == "on" {
		app.Use(middleware.CorazaMiddleware()) // CorazaWAF
	}

	// 路由分组
	userApi(app.Group("/api/system/user"))

	return app
}
