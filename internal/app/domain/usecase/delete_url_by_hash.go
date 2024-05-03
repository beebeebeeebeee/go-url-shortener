package usecase

import (
	"errors"
	"url-shortener/internal/app/domain/repository"
	"url-shortener/internal/app/domain/service"
)

func DeleteURLByHash(
	hashService service.IHashService,
	repository repository.IShortenerRepository,
	hash string,
	password string,
) error {
	rsp, err := repository.GetByHash(hash)
	if err != nil {
		return err
	}

	hashedPassword := hashService.MD5(password)
	if rsp.Password != hashedPassword {
		return errors.New("password is incorrect")
	}

	return repository.DeleteByHash(hash)
}
