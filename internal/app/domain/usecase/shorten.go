package usecase

import (
	"url-shortener/internal/app/domain/repository"
	"url-shortener/internal/app/domain/service"
)

func Shorten(
	hashService service.IHashService,
	shortenerRepository repository.IShortenerRepository,
	url string,
) (string, string, error) {
	var hash string
	for {
		h, err := hashService.CreateHash(5)
		if err != nil {
			return "", "", err
		}

		if _, err := shortenerRepository.GetByHash(h); err != nil {
			hash = h
			break
		}
	}

	password, err := hashService.CreateHash(10)
	if err != nil {
		return "", "", err
	}

	hashedPassword := hashService.MD5(password)

	if err = shortenerRepository.Save(hash, url, hashedPassword); err != nil {
		return "", "", err
	}

	return hash, password, nil
}
