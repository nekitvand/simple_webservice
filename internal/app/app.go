package app

import (
	"log"
	"net/http"
	"webservice/api"
	"webservice/internal/cfg"
	"webservice/internal/handler"
	"webservice/internal/repository"
	"webservice/internal/service"

	_ "webservice/internal/migrations"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
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
	server.db, err = sqlx.Open("pgx", server.config.GetDBString())
	if err != nil {
		log.Fatalln(err)
	}
	server.db.SetMaxIdleConns(10)
	server.db.SetMaxOpenConns(10)
	log.Println("Migration start")
	if err := goose.Up(server.db.DB,"/var"); 
	err != nil {
		panic(err)
	}

	todoRepository := repository.NewToDoRepository(server.db)
	todoService := service.NewUsersService(todoRepository)
	todoHandler := handler.NewToDoHandler(todoService)
	route := api.CreateRoutes(todoHandler)

	server.srv = &http.Server{
		Addr:    ":" + server.config.Port,
		Handler: route,
	}

	log.Println("Server started")

	err = server.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

func (server *AppServer) Shutdown() {

	server.db.Close()

}
