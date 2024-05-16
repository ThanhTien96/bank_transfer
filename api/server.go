package api

import (
	"net/http"
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	route *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	route := gin.Default()

	route.GET("/check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok you are pass")
	})

	route.POST("/accounts", server.CreateAccount)
	route.GET("/accounts/:id", server.GetAccount)
	route.GET("/accounts", server.ListAccounts)

	server.route = route

	return server
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(addr string) error {
	return s.route.Run(addr)
}


func errResponse(status http.ConnState, err error) *gin.H {
	return &gin.H{
		"error": err.Error(),
		"status": status,
		"success": false,
	}
}


func successResponse(message string, data interface{}) *gin.H {
	return &gin.H{
		"status": http.StatusOK,
		"message": message,
		"data": data,
		"success": true,
	}
}
