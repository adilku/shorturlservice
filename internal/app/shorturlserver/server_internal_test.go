package shorturlserver

import (
	"bytes"
	"encoding/json"
	"github.com/adilku/shorturlservice/internal/app/model"
	"github.com/adilku/shorturlservice/internal/app/store/simplestore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_handleCreateShortLink(t *testing.T) {
	s := newServer(simplestore.New())
	testCases := []struct {
		name 			string
		payload			interface{}
		expectedCode	int
	} {
		{
			name: "valid",
			payload: map[string]string{
				"url": model.TestURL(t),
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid payload",
			payload: "",
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/urls", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleFind(t *testing.T) {
	s := newServer(simplestore.New())
	testCases := []struct {
		name 			string
		payload			interface{}
		expectedCode	int
	} {
		{
			name: "invalid",
			payload: map[string]string{
				"url": model.TestURL(t),
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid payload",
			payload: "",
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodGet, "/urls", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleFindGood(t *testing.T) {
	type responseShort struct {
		ShortUrl string `json:"shortUrl"`
	}
	type responseLong struct {
		LongUrl string `json:"longUrl"`
	}

	s := newServer(simplestore.New())
	rec := httptest.NewRecorder()
	b := &bytes.Buffer{}
	testCase := map[string]string{"url": model.TestURL(t)}
	json.NewEncoder(b).Encode(testCase)
	req, _ := http.NewRequest(http.MethodPost, "/urls", b)
	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)
	res := rec.Result()
	resp := &responseShort{}
	json.NewDecoder(res.Body).Decode(resp)
	expected := resp.ShortUrl

	rec2 := httptest.NewRecorder()
	b2 := &bytes.Buffer{}
	testCase2 := map[string]string{"url": expected}
	json.NewEncoder(b2).Encode(testCase2)
	req, _ = http.NewRequest(http.MethodGet, "/urls", b2)
	s.ServeHTTP(rec2, req)
	assert.Equal(t, http.StatusOK, rec2.Code)

	res = rec2.Result()
	respLong := &responseLong{}
	json.NewDecoder(res.Body).Decode(respLong)
	assert.Equal(t, model.TestURL(t), respLong.LongUrl)
}