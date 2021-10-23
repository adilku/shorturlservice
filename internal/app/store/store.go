package store

type Store interface {
	GetUrls() UrlRepository
}