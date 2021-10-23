package store

import "github.com/adilku/shorturlservice/internal/app/model"

type UrlRepository interface {
	Create(url *model.Url) error
	FindByShort(shortUrl string) (string ,error)
}