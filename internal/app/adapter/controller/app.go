package controller

import (
	"github.com/gin-gonic/gin"
	"url-shortener/internal/app/domain/repository"
	"url-shortener/internal/app/domain/service"
	"url-shortener/internal/app/domain/usecase"
	"url-shortener/internal/cfg"
)

type AppController struct {
	hashService         service.IHashService
	shortenerRepository repository.IShortenerRepository
}

func NewAppController(
	hashService service.IHashService,
	shortenerRepository repository.IShortenerRepository,
) AppController {
	return AppController{
		hashService:         hashService,
		shortenerRepository: shortenerRepository,
	}
}

type ShortenRequest struct {
	URL string `json:"url"`
}

func (ac *AppController) Shorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	hash, password, err := usecase.Shorten(ac.hashService, ac.shortenerRepository, req.URL)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"hash":     hash,
		"url":      cfg.Cfg.App.BaseURL + hash,
		"password": password,
	})
}

func (ac *AppController) Redirect(c *gin.Context) {
	hash, isExist := c.Params.Get("hash")
	if !isExist {
		c.Redirect(302, cfg.Cfg.App.BaseURL)
		return
	}

	url, err := usecase.GetURLByHash(ac.shortenerRepository, hash)
	if err != nil {
		c.Redirect(302, cfg.Cfg.App.BaseURL)
		return
	}

	c.Redirect(302, url)
}

type DeleteRequest struct {
	Password string `json:"password"`
}

func (ac *AppController) Delete(c *gin.Context) {
	hash, isExist := c.Params.Get("hash")
	if !isExist {
		c.JSON(500, gin.H{
			"error": "hash is required",
		})
		return
	}

	var req DeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := usecase.DeleteURLByHash(ac.hashService, ac.shortenerRepository, hash, req.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "deleted",
	})

}

func (ac *AppController) Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
