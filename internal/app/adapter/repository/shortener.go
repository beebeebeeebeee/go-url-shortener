package repository

import (
	"url-shortener/internal/app/adapter/util/converter"
	"url-shortener/internal/app/domain/entity"
	"url-shortener/internal/app/domain/repository"
	"url-shortener/internal/app/infrastructure/go_orm"
)

type ShortenerRepository struct {
	client go_orm.Client
}

func NewShortenerRepository(
	client go_orm.Client,
) repository.IShortenerRepository {
	return &ShortenerRepository{
		client: client,
	}
}

func (r *ShortenerRepository) Save(hash string, url string, password string) error {
	return r.client.CreateShortenURL(hash, url, password)
}

func (r *ShortenerRepository) GetByHash(hash string) (entity.ShortenURL, error) {
	shortenURL, err := r.client.GetShortenURLByCode(hash)
	if err != nil {
		return entity.ShortenURL{}, err
	}

	return converter.ShortenURLInfrastructureToDomain(*shortenURL), nil
}

func (r *ShortenerRepository) DeleteByHash(hash string) error {
	return r.client.DeleteShortenURLByCode(hash)
}
