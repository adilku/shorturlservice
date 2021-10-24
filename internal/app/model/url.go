package model

import (
	//validation "github.com/go-ozzo/ozzo-validation"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	gonanoid "github.com/matoous/go-nanoid"
)

type Url struct {
	Address string `json:"address"`
	ShortAddress string `json:"short_address"`
}

func (u *Url) Validate() error {
	return validation.Validate(
		u.Address,
		validation.Required,
		is.URL,
		)
}

func GenerateNewUrl() (string, error) {
	shortUrl, err :=  gonanoid.Generate(
		"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-",
		4,
		)
	if err != nil {
		return "", err
	}
	return "bi.ly/" + shortUrl, nil
}
