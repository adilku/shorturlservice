package simplestore

import "github.com/adilku/shorturlservice/internal/app/store"

type Store struct {
	urlRepository *UrlRepository
}


func New() *Store {
	return &Store{}
}

func (s *Store) GetUrls() store.UrlRepository {
	if s.urlRepository != nil {
		return s.urlRepository
	}

	s.urlRepository = &UrlRepository{
		store: s,
		data: make(map[string]string, 0),
	}

	return s.urlRepository
}