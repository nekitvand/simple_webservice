package app

import (
	"webservice/internals/cfg"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type AppServer struct {
	config cfg.Cfg
	srv    *http.Server
	db     *sqlx.DB
}

func NewServer(config cfg.Cfg) *AppServer {
	server := new(AppServer)
	server.config = config
	return server
}

func (server *AppServer) Start() {
	var err error
	server.db, err = sqlx.Open("pgx",server.config.GetDBString())
	if err != nil {
		log.Fatalln(err)
	}
	server.db.SetMaxIdleConns(10)
	server.db.SetMaxOpenConns(10)

	
	server.srv = &http.Server{
		Addr:    ":" + server.config.Port,
		// Handler: routes,
	}

	log.Println("Server started")

	err = server.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}

	return 
}

func (server *AppServer) Shutdown() {

	server.db.Close()

}
