package postgrestore

import (
	"github.com/adilku/shorturlservice/internal/app/model"
)

type UrlRepository struct {
	store *Store
}

func (r UrlRepository) Create(url *model.Url) error {
	err := r.store.db.QueryRow(
		"INSERT INTO urls (short_url, long_url) VALUES ($1, $2)",
		url.ShortAddress,
		url.Address,
		).Err()
	return err
}

func (r UrlRepository) FindByShort(shortUrl string) (string, error) {
	var longUrl string
	err := r.store.db.QueryRow(
		"SELECT long_url FROM urls WHERE short_url=$1", shortUrl).Scan(&longUrl)
	if err != nil {
		return "", err
	}
	return longUrl, nil
}


