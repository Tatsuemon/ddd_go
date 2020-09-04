package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/Tatsuemon/ddd_go/infrastructure/web/handler"
	"github.com/Tatsuemon/ddd_go/infrastructure/web/router"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	db     *sqlx.DB
	router *mux.Router
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(db *sqlx.DB, h handler.AppHandler) {
	s.router = router.NewMuxRouter(h)
	s.db = db
}

func (s *Server) Run(port int) {
	fmt.Printf("Server running at http://loacalhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), s.router)
	return
}
