package main

import (
	"url-shortener/internal/app"
	"url-shortener/internal/cfg"
)

func main() {
	cfg.LoadEnv(".")
	a := app.NewApp(&cfg.Cfg)
	a.Run()

}
