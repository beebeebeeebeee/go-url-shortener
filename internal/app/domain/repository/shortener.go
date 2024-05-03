package repository

import "url-shortener/internal/app/domain/entity"

type IShortenerRepository interface {
	Save(hash string, url string, password string) error
	GetByHash(hash string) (entity.ShortenURL, error)
	DeleteByHash(hash string) error
}
