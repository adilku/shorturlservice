package postgrestore_test

import (
	"github.com/adilku/shorturlservice/internal/app/model"
	"github.com/adilku/shorturlservice/internal/app/store/postgrestore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlRepository_Create(t *testing.T) {
	db, teardown := postgrestore.TestDB(t, databaseURL)
	defer teardown("urls")
	s := postgrestore.New(db)
	m := model.TestFakeModel(t)
	err := s.GetUrls().Create(&m)
	assert.NoError(t, err)
}

func TestUrlRepository_CreateBad(t *testing.T) {
	db, teardown := postgrestore.TestDB(t, databaseURL)
	defer teardown("urls")
	s := postgrestore.New(db)
	m := model.TestFakeModel(t)
	err := s.GetUrls().Create(&m)
	assert.NoError(t, err)
	err = s.GetUrls().Create(&m)
	assert.Error(t, err)
}

func TestUrlRepository_FindByShortBad(t *testing.T) {
	db, teardown := postgrestore.TestDB(t, databaseURL)
	defer teardown("urls")
	s := postgrestore.New(db)
	_, err := s.GetUrls().FindByShort(model.TestFakeModel(t).Address)
	assert.Error(t, err)
}

func TestUrlRepository_FindByShort(t *testing.T) {
	db, teardown := postgrestore.TestDB(t, databaseURL)
	defer teardown("urls")
	s := postgrestore.New(db)
	m := model.TestFakeModel(t)
	err := s.GetUrls().Create(&m)
	assert.NoError(t, err)
	longUrl, err := s.GetUrls().FindByShort(m.ShortAddress)
	assert.NoError(t, err)
	assert.Equal(t, m.Address, longUrl)
}
