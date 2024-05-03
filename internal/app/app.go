package app

import (
	"context"
	"go.uber.org/fx"
	"strconv"
	"url-shortener/internal/app/adapter/controller"
	"url-shortener/internal/app/adapter/repository"
	"url-shortener/internal/app/adapter/route"
	"url-shortener/internal/app/adapter/service"
	"url-shortener/internal/app/infrastructure/gin"
	"url-shortener/internal/app/infrastructure/go_orm"
	"url-shortener/internal/cfg"
)

type App struct {
	Cfg    *cfg.Config
	ctx    context.Context
	module fx.Option
}

func NewApp(cfg *cfg.Config) *App {
	return &App{
		Cfg: cfg,
		ctx: context.Background(),
		module: fx.Options(
			controller.Module,
			repository.Module,
			route.Module,
			service.Module,
			gin.Module,
			go_orm.Module,
		),
	}
}

func (a *App) Run() {
	app := fx.New(a.module, fx.Options(
		fx.Invoke(func(
			routes route.Routes,
			requestHandler gin.RequestHandler,
		) {
			routes.Setup()
			if err := requestHandler.Gin.Run(":" + strconv.Itoa(a.Cfg.App.Port)); err != nil {
				panic(err)
			}
		}),
	))
	if err := app.Start(a.ctx); err != nil {
		panic(err)
	}
}
