package entity

type ShortenURL struct {
	ID       int64  `gorm:"primaryKey"`
	Code     string `gorm:"unique"`
	URL      string
	Password string
}
