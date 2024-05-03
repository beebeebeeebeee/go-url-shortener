package service

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"url-shortener/internal/app/domain/service"
)

type HashService struct {
}

func NewHashService() service.IHashService {
	return &HashService{}
}

func (hs *HashService) CreateHash(length int64) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	hash := base64.URLEncoding.EncodeToString(b)
	return hash[:length], nil
}

func (hs *HashService) MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
