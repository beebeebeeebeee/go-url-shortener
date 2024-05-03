package route

import (
	"url-shortener/internal/app/adapter/controller"
	"url-shortener/internal/app/infrastructure/gin"
)

type AppRoute struct {
	handler       gin.RequestHandler
	appController controller.AppController
}

func NewAppRoute(
	handler gin.RequestHandler,
	appController controller.AppController,
) AppRoute {
	return AppRoute{
		handler:       handler,
		appController: appController,
	}
}

func (ar *AppRoute) Setup() {
	app := ar.handler.Gin.Group("/")
	{
		app.GET("/", ar.appController.Index)
		app.POST("/shorten", ar.appController.Shorten)
		app.GET("/:hash", ar.appController.Redirect)
		app.DELETE("/:hash", ar.appController.Delete)
	}
}
