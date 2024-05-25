package api

import (
	"database/sql"
	"fmt"
	"net/http"
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)


type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=VND USD EUR"`
}

func (s *Server) CreateAccount(ctx *gin.Context) {
	fmt.Println("CreateAccount API is called.")
	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(http.StatusBadRequest,err))
		return
	}

	arg := db.CreateAccountParams{
		Owner: req.Owner,
		Currency: req.Currency,
		Balance: 0,
	}

	account, err := s.store.CreateAccount(ctx, arg)
	if err != nil {
		fmt.Println("error occured when db create account")
		ctx.JSON(http.StatusInternalServerError, errResponse(http.StatusInternalServerError, err))
		return
	}
	ctx.JSON(http.StatusOK, successResponse("Created account", account))
}

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (s *Server) GetAccount(ctx *gin.Context) {
	var req GetAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(http.StatusBadRequest, err))
		return
	}

	account, err := s.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(http.StatusNotFound, err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type ListAccountRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *Server) ListAccounts(ctx *gin.Context) {
	var req ListAccountRequest 
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(http.StatusBadRequest, err))
		return 
	}

	arg := &db.ListAccountsParams{
		Limit: req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,

	}

	accounts, err := s.store.ListAccounts(ctx, *arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Get account successfully", accounts))
}