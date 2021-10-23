package simplestore_test

import (
	"github.com/adilku/shorturlservice/internal/app/model"
	"github.com/adilku/shorturlservice/internal/app/store/simplestore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlRepository_Create(t *testing.T) {
	s := simplestore.New()
	longUrl := model.TestURL(t)
	shortURL, err := model.GenerateNewUrl()
	assert.NoError(t, err)
	m := model.Url{Address: longUrl, ShortAddress: shortURL}
	err = s.GetUrls().Create(&m)
	assert.NoError(t, err)
}

func TestUrlRepository_CreateExist(t *testing.T) {
	s := simplestore.New()
	longUrl := model.TestURL(t)
	shortURL, err := model.GenerateNewUrl()
	assert.NoError(t, err)
	m := model.Url{Address: longUrl, ShortAddress: shortURL}
	err = s.GetUrls().Create(&m)
	assert.NoError(t, err)
	err = s.GetUrls().Create(&m)
	assert.Error(t, err)
}

func TestUrlRepository_FindByShort(t *testing.T) {
	s := simplestore.New()
	longUrl := model.TestURL(t)
	shortURL, err := model.GenerateNewUrl()
	assert.NoError(t, err)
	m := model.Url{Address: longUrl, ShortAddress: shortURL}
	err = s.GetUrls().Create(&m)
	assert.NoError(t, err)
	longAfter, err := s.GetUrls().FindByShort(shortURL)
	assert.NoError(t, err)
	assert.Equal(t, longUrl, longAfter)
}

func TestUrlRepository_FindByShort2(t *testing.T) {
	s := simplestore.New()
	_, err := s.GetUrls().FindByShort("dedwede")
	assert.Error(t, err)
}

