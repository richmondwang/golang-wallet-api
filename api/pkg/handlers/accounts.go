package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/richmondwang/golang-wallet-api/ent"
	"github.com/richmondwang/golang-wallet-api/pkg/services"
)

// this is just for workaround to remove edges on swagger docs
type Empty struct{}

type ctxKey struct {
	key int8
}

var (
	CtxAccountKey = &ctxKey{1}
)

// AccountRequest the request body for creating a new acount
type AccountRequest struct {
	Name           string  `json:"name"`
	InitialBalance float64 `json:"initial_balance"`
}

// Bind validation and cleanup
func (ar *AccountRequest) Bind(r *http.Request) error {
	ar.Name = strings.Trim(ar.Name, " ")
	if ar.Name == "" {
		return errors.New("name is required to create an account")
	}
	if ar.InitialBalance < 0 {
		return errors.New("cannot create a new account with a negative balance")
	}
	return nil
}

// TransferRequest request body for creating a transfer
type TransferRequest struct {
	AccountID int     `json:"account_id"`
	Amount    float64 `json:"amount"`
}

// Bind validation
func (tr *TransferRequest) Bind(r *http.Request) error {
	if tr.AccountID == 0 {
		return errors.New("account id is required to transfer money")
	}
	if tr.Amount <= 0 {
		return errors.New("can only transfer an amount more than 0")
	}
	return nil
}

type accounts struct {
	service services.Accounts
}

// NewAccountsHandler creates a new accounts handler
func NewAccountsHandler(db *ent.Client) *accounts {
	return &accounts{
		service: services.NewAccountsService(db),
	}
}

// SetService sets the accounts service
func (a *accounts) SetService(service services.Accounts) {
	a.service = service
}

// GetAccountMiddleware prefetch account before handling
func (h *accounts) GetAccountMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accountID, err := strconv.Atoi(chi.URLParam(r, "accountID"))
		if err != nil {
			render.Render(w, r, NewResponseError(errors.New("invalid account id")).
				WithCode(http.StatusBadRequest))
			return
		}
		pctx := r.Context()
		userAccount, err := h.service.GetAccount(pctx, accountID)
		if err != nil {
			render.Render(w, r, NewResponseError(err).
				WithCode(http.StatusInternalServerError))
			return
		}
		if userAccount == nil {
			render.Render(w, r, NewResponseError(fmt.Errorf("account %d not found", accountID)).
				WithCode(http.StatusNotFound))
			return
		}

		ctx := context.WithValue(pctx, CtxAccountKey, userAccount)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetAccount godoc
// @Summary      Retrieve an account using an account ID
// @Description  Retrieve an account using an account ID
// @Tags         account
// @Produce      json
// @Success      200  {object}   ResponseWrapper{data=ent.Account}
// @Failure      404  {object}   ResponseWrapper
// @Failure      400  {object}   ResponseWrapper
// @Failure      500  {object}   ResponseWrapper
// @Param 		 accountID  path int true "Account ID"
// @Router       /accounts/{accountID} [get]
func (h *accounts) GetAccount(w http.ResponseWriter, r *http.Request) {
	account := r.Context().Value(CtxAccountKey).(*ent.Account)
	render.Render(w, r, NewResponse(account))
}

// Transactions godoc
// @Summary      Retrieve transactions of an account
// @Description  Retrieve transactions of an account
// @Tags         account
// @Produce      json
// @Success      200  {object}   ResponseWrapper{data=ent.Transactions}
// @Failure      404  {object}   ResponseWrapper
// @Failure      400  {object}   ResponseWrapper
// @Failure      500  {object}   ResponseWrapper
// @Param 		 accountID  path  int    true  "Account ID"
// @Param   	 type       query string false "Type of transactions" Enums(all, incoming, outgoing)
// @Router       /accounts/{accountID}/transactions [get]
func (h *accounts) Transactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	account := ctx.Value(CtxAccountKey).(*ent.Account)
	txType := r.URL.Query().Get("type")

	var transactions ent.Transactions
	var err error
	switch {
	case txType == "incoming":
		transactions, err = h.service.IncomingTransactions(ctx, account)
	case txType == "outgoing":
		transactions, err = h.service.OutgoingTransactions(ctx, account)
	default:
		transactions, err = h.service.AllTransactions(ctx, account)
	}
	if err != nil {
		render.Render(w, r, NewResponseError(err).
			WithCode(http.StatusInternalServerError))
		return
	}
	render.Render(w, r, NewResponse(transactions))
}

// CreateAccount godoc
// @Summary      Create an account
// @Description  Create an account
// @Tags         account
// @Produce      json
// @Success      201  {object}   ResponseWrapper{data=ent.Account}
// @Failure      400  {object}   ResponseWrapper
// @Failure      500  {object}   ResponseWrapper
// @Param 		 request  		 body AccountRequest true "Data of the account"
// @Router       /accounts [post]
func (h *accounts) CreateAccount(w http.ResponseWriter, r *http.Request) {
	data := &AccountRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, NewResponseError(err).
			WithCode(http.StatusBadRequest))
		return
	}

	userAccount, err := h.service.AddAccount(r.Context(), data.Name, data.InitialBalance)
	if err != nil {
		render.Render(w, r, NewResponseError(err).
			WithCode(http.StatusInternalServerError))
		return
	}
	render.Render(w, r, NewResponse(userAccount).WithCode(http.StatusCreated))
}

// TransferMoney godoc
// @Summary      Transfer money
// @Description  Transfer an amount from an account's wallet to another
// @Tags         account
// @Produce      json
// @Success      201  {object}   ResponseWrapper{data=ent.Transaction}
// @Failure      400  {object}   ResponseWrapper
// @Failure      500  {object}   ResponseWrapper
// @Param 		 accountID path int             true "Account ID"
// @Param 		 request   body TransferRequest true "Account and amount to transfer"
// @Router       /accounts/{accountID}/transfer [post]
func (h *accounts) TransferMoney(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fromAccount := ctx.Value(CtxAccountKey).(*ent.Account)

	data := &TransferRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, NewResponseError(err).
			WithCode(http.StatusBadRequest))
		return
	}
	if data.AccountID == fromAccount.ID {
		render.Render(w, r, NewResponseError(errors.New("cannot transfer to own account")).
			WithCode(http.StatusBadRequest))
		return
	}

	// querying destination account and check if it exists
	toAccount, err := h.service.GetAccount(ctx, data.AccountID)
	if err != nil {
		render.Render(w, r, NewResponseError(err).
			WithCode(http.StatusInternalServerError))
		return
	}
	if toAccount == nil {
		render.Render(w, r, NewResponseError(fmt.Errorf("account %d not found", data.AccountID)).
			WithCode(http.StatusBadRequest))
		return
	}
	// they can transfer money they dont have
	if fromAccount.Balance < data.Amount {
		render.Render(w, r, NewResponseError(services.ErrInsufficientBalance).
			WithCode(http.StatusBadRequest))
		return
	}

	transaction, err := h.service.Transfer(ctx, fromAccount, toAccount, data.Amount)
	if err == services.ErrInsufficientBalance {
		render.Render(w, r, NewResponseError(err).
			WithCode(http.StatusBadRequest))
		return
	}
	if err != nil {
		render.Render(w, r, NewResponseError(err).
			WithCode(http.StatusInternalServerError))
		return
	}
	// if we dont want to show the other party's balance
	// transaction.Edges.ToAccount.Balance = 0
	render.Render(w, r, NewResponse(transaction).WithCode(http.StatusCreated))
}
