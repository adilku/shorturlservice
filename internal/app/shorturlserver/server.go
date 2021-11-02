package shorturlserver

import (
	"encoding/json"
	"github.com/adilku/shorturlservice/internal/app/model"
	"github.com/adilku/shorturlservice/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store: store,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/urls", s.handleCreateShortLink()).Methods("POST", "GET")
}

func (s *server) handleCreateShortLink() http.HandlerFunc {
	//additional things
	type request struct {
		Url string `json:"url"`
	}
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "POST":
			req := &request{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}
			shortUrl, err := model.GenerateNewUrl()
			if err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
			}
			u := model.Url{ShortAddress: shortUrl, Address: req.Url}
			if err := s.store.GetUrls().Create(&u); err != nil {
				s.error(w, r, http.StatusConflict, err)
				return
			}
			responseShort := map[string]string{"shortUrl" : shortUrl}
			s.respond(w, r, http.StatusCreated, responseShort)
		case "GET":
			shortUrl := r.URL.Query().Get("shortUrl")
			originalUrl, err := s.store.GetUrls().FindByShort(shortUrl)
			if err != nil {
				s.error(w, r, http.StatusNoContent, err)
				return
			}
			responseLong := map[string]string{"longUrl" : originalUrl}
			s.respond(w, r, http.StatusOK, responseLong)
		}
	}
}


func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

