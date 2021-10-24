package postgrestore

import (
	"database/sql"
	"github.com/adilku/shorturlservice/internal/app/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
	urlRepository *UrlRepository
}

func New (db *sql.DB) *Store {
	return &Store{
		db : db,
	}
}

func (s *Store) GetUrls() store.UrlRepository {
	if s.urlRepository != nil {
		return s.urlRepository
	}
	s.urlRepository = &UrlRepository{
		store: s,
	}
	return s.urlRepository
}