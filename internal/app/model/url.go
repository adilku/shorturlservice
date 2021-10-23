package model

import gonanoid "github.com/matoous/go-nanoid"

type Url struct {
	Address string `json:"address"`
	ShortAddress string `json:"short_address"`
}

func GenerateNewUrl() (string, error) {
	shortUrl, err := gonanoid.Generate(
		"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-",
		10,
		)
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}
