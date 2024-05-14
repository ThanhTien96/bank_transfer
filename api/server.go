package api

import (
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store db.Store
	route *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}
	route := gin.Default()

	route.POST("/accounts", server.CreateAccount)

	server.route = route

	return server
}


func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}