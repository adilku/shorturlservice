package shorturlserver

import (
	"github.com/adilku/shorturlservice/internal/app/store/simplestore"
	"net/http"
)

func Start(config *Config) error {
	db := simplestore.New()
	s := newServer(db)
	s.logger.Println("statring at port", config.BindAddr)
	return http.ListenAndServe(config.BindAddr, s)
}