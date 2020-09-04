package server

import (
	"fmt"
	"net/http"

	"ddd_go/infrastructure/web/router"
)

func Run(port int) {
	router := router.NewMuxRouter()
	fmt.Printf("Server running at http://loacalhost:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
