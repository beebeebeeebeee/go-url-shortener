package go_orm

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"url-shortener/internal/app/infrastructure/go_orm/entity"
)

type Client struct {
	orm *gorm.DB
}

func NewClient() Client {
	db, err := gorm.Open(sqlite.Open("db/shortener_db.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	if err = db.AutoMigrate(
		&entity.ShortenURL{},
	); err != nil {
		panic("failed to migrate database")
	}

	return Client{
		orm: db,
	}
}

func (c *Client) CreateShortenURL(code string, url string, password string) error {
	return c.orm.Create(&entity.ShortenURL{
		Code:     code,
		URL:      url,
		Password: password,
	}).Error
}

func (c *Client) GetShortenURLByCode(code string) (*entity.ShortenURL, error) {
	var shortenURL entity.ShortenURL
	if err := c.orm.Where("code = ?", code).First(&shortenURL).Error; err != nil {
		return nil, err
	}

	return &shortenURL, nil
}

func (c *Client) DeleteShortenURLByCode(code string) error {
	return c.orm.Where("code = ?", code).Delete(&entity.ShortenURL{}).Error
}
