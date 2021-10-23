package simplestore

import (
	"errors"
	"github.com/adilku/shorturlservice/internal/app/model"
)

type UrlRepository struct {
	store *Store
	data  map[string]string
}

func (r *UrlRepository) Create(url *model.Url) error {
	if _, ok := r.data[url.ShortAddress]; !ok {
		r.data[url.ShortAddress] = url.Address
	} else {
		return errors.New("collision")
	}
	return nil
}

func (r *UrlRepository) FindByShort(shortUrl string) (string, error) {
	if _, ok := r.data[shortUrl]; !ok {
		return "", errors.New("cannot find url in database")
	} else {
		return r.data[shortUrl], nil
	}
}



