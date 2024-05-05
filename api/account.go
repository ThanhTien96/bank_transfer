package api

import (
	"net/http"
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)


type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof:VND USD EUR"`
}

func (s *Server) CreateAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner: req.Owner,
		Currency: req.Currency,
		Balance: 0,
	}

	account, err := s.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}