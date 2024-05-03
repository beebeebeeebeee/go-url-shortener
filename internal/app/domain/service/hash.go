package service

type IHashService interface {
	CreateHash(length int64) (string, error)
	MD5(text string) string
}
