package usecase

import "url-shortener/internal/app/domain/repository"

func GetURLByHash(
	repository repository.IShortenerRepository,
	hash string,
) (string, error) {
	url, err := repository.GetByHash(hash)
	if err != nil {
		return "", err
	}

	return url.URL, nil
}
