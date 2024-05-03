package converter

import (
	domain "url-shortener/internal/app/domain/entity"
	infrastructure "url-shortener/internal/app/infrastructure/go_orm/entity"
)

func ShortenURLInfrastructureToDomain(
	url infrastructure.ShortenURL,
) domain.ShortenURL {
	return domain.ShortenURL{
		ID:       url.ID,
		Code:     url.Code,
		URL:      url.URL,
		Password: url.Password,
	}
}
