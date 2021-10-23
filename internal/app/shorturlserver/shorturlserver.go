package shorturlserver

import (
	"database/sql"
	"github.com/adilku/shorturlservice/internal/app/store/postgrestore"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := postgrestore.New(db)
	s := newServer(store)
	s.logger.Println("statring at port", config.BindAddr)
	return http.ListenAndServe(config.BindAddr, s)
}

func newDB(dabaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dabaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}