package simplestore

import (
	"errors"
	"github.com/adilku/shorturlservice/internal/app/model"
	"log"
	"sync"
)

//TODO mutex

type UrlRepository struct {
	store *Store
	data  map[string]string
	mu 	  sync.RWMutex
}

func (r *UrlRepository) Create(url *model.Url) error {
	if err := url.Validate(); err != nil {
		return err
	}
	r.mu.Lock()
	r.data[url.ShortAddress] = url.Address
	r.mu.Unlock()
	return nil
}

func (r *UrlRepository) FindByShort(shortUrl string) (string, error) {
	for key, val := range r.data {
		log.Println(key, val)
	}
	log.Println(shortUrl)
	r.mu.RLock()
	val, ok := r.data[shortUrl]
	r.mu.RUnlock()
	if !ok {
		return "", errors.New("cannot find url in database")
	} else {
		return val, nil
	}
}



