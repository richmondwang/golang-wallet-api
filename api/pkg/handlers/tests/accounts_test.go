package handlers_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/richmondwang/golang-wallet-api/ent"
	"github.com/richmondwang/golang-wallet-api/mocks"
	"github.com/richmondwang/golang-wallet-api/pkg/handlers"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var mockContext = mock.MatchedBy(func(c context.Context) bool { return true })

func TestAccountHandler_GetAccount(t *testing.T) {
	accountsHandler := handlers.NewAccountsHandler(nil)
	serviceMock := mocks.NewAccounts(t)

	tNow := time.Now()
	serviceMock.On("GetAccount", mockContext, 1).
		Return(&ent.Account{
			ID:        1,
			Name:      "Richmond Wang",
			Balance:   100,
			CreatedAt: &tNow,
		}, nil)

	accountsHandler.SetService(serviceMock)
	s := createMockServer()
	s.router.Mount("/", accountsHandler.Routes(s.router))

	req, _ := http.NewRequest(http.MethodGet, "/accounts/1", nil)
	resp := executeMockRequest(req, s)

	require.Equal(t, http.StatusOK, resp.Result().StatusCode)

	rBody, _ := io.ReadAll(resp.Result().Body)
	acc := &handlers.ResponseWrapper{Data: &ent.Account{}}
	err := json.Unmarshal(rBody, acc)

	require.Equal(t, nil, err)
	require.Equal(t, "", acc.Error)
	require.Equal(t, 200, acc.Code)
	require.Equal(t, 1, acc.Data.(*ent.Account).ID)
	require.Equal(t, "Richmond Wang", acc.Data.(*ent.Account).Name)
	require.Equal(t, float64(100), acc.Data.(*ent.Account).Balance)
}

func TestAccountHandler_CreateAccount(t *testing.T) {
	accountsHandler := handlers.NewAccountsHandler(nil)
	serviceMock := mocks.NewAccounts(t)

	tNow := time.Now()
	serviceMock.On("AddAccount", mockContext, "Skyler Chase", float64(50)).
		Return(&ent.Account{
			ID:        2,
			Name:      "Skyler Chase",
			Balance:   50,
			CreatedAt: &tNow,
		}, nil)

	accountsHandler.SetService(serviceMock)
	s := createMockServer()
	s.router.Mount("/", accountsHandler.Routes(s.router))

	body := strings.NewReader(`
	{
		"name": "Skyler Chase",
		"initial_balance": 50
	}`)
	req, _ := http.NewRequest(http.MethodPost, "/accounts/", body)
	resp := executeMockRequest(req, s)

	require.Equal(t, http.StatusCreated, resp.Result().StatusCode)

	rBody, _ := io.ReadAll(resp.Result().Body)
	acc := &handlers.ResponseWrapper{Data: &ent.Account{}}
	err := json.Unmarshal(rBody, acc)

	require.Equal(t, nil, err)
	require.Equal(t, "", acc.Error)
	require.Equal(t, 201, acc.Code)
	require.Equal(t, 2, acc.Data.(*ent.Account).ID)
	require.Equal(t, "Skyler Chase", acc.Data.(*ent.Account).Name)
	require.Equal(t, float64(50), acc.Data.(*ent.Account).Balance)
}

func TestAccountHandler_TransferMoney(t *testing.T) {
	accountsHandler := handlers.NewAccountsHandler(nil)
	serviceMock := mocks.NewAccounts(t)

	tNow := time.Now()
	fromAccount := &ent.Account{
		ID:        1,
		Name:      "Richmond Wang",
		Balance:   100,
		CreatedAt: &tNow,
	}
	toAccount := &ent.Account{
		ID:        2,
		Name:      "Skyler Chase",
		Balance:   15,
		CreatedAt: &tNow,
	}
	serviceMock.On("GetAccount", mockContext, 1).Return(fromAccount, nil)
	serviceMock.On("GetAccount", mockContext, 2).Return(toAccount, nil)

	serviceMock.On("Transfer", mockContext, fromAccount, toAccount, float64(20)).
		Return(&ent.Transaction{
			ID:        1,
			Amount:    20,
			CreatedAt: tNow,
		}, nil)

	accountsHandler.SetService(serviceMock)
	s := createMockServer()
	s.router.Mount("/", accountsHandler.Routes(s.router))

	t.Run("Sufficient Balance", func(t *testing.T) {
		body := strings.NewReader(`
		{
			"account_id": 2,
			"amount": 20
		}`)
		req, _ := http.NewRequest(http.MethodPost, "/accounts/1/transfer", body)
		resp := executeMockRequest(req, s)
		require.Equal(t, http.StatusCreated, resp.Result().StatusCode)

		rBody, _ := io.ReadAll(resp.Result().Body)
		acc := &handlers.ResponseWrapper{Data: &ent.Transaction{}}
		err := json.Unmarshal(rBody, acc)
		require.Equal(t, nil, err)
		require.Equal(t, "", acc.Error)
		require.Equal(t, 201, acc.Code)
		require.Equal(t, 1, acc.Data.(*ent.Transaction).ID)
		require.Equal(t, float64(20), acc.Data.(*ent.Transaction).Amount)
	})

	// test insufficient balance
	t.Run("Insufficient Balance", func(t *testing.T) {
		body := strings.NewReader(`
		{
			"account_id": 2,
			"amount": 200
		}`)
		req, _ := http.NewRequest(http.MethodPost, "/accounts/1/transfer", body)
		resp := executeMockRequest(req, s)
		require.Equal(t, http.StatusBadRequest, resp.Result().StatusCode)

		rBody, _ := io.ReadAll(resp.Result().Body)
		acc := &handlers.ResponseWrapper{Data: &ent.Transaction{}}
		err := json.Unmarshal(rBody, acc)
		require.Equal(t, nil, err)
		require.Equal(t, "account has insufficient balance", acc.Error)
		require.Equal(t, 400, acc.Code)
		require.Equal(t, nil, acc.Data)
	})
}

func TestAccountHandler_Transactions(t *testing.T) {
	accountsHandler := handlers.NewAccountsHandler(nil)
	serviceMock := mocks.NewAccounts(t)

	tNow := time.Now()
	fromAccount := &ent.Account{
		ID:        1,
		Name:      "Richmond Wang",
		Balance:   100,
		CreatedAt: &tNow,
	}
	toAccount := &ent.Account{
		ID:        2,
		Name:      "Skyler Chase",
		Balance:   15,
		CreatedAt: &tNow,
	}
	serviceMock.On("GetAccount", mockContext, 1).Return(fromAccount, nil)
	accountsHandler.SetService(serviceMock)
	s := createMockServer()
	s.router.Mount("/", accountsHandler.Routes(s.router))
	type ResponseWrapper struct {
		Data  ent.Transactions `json:"data"`
		Error string           `json:"error,omitempty"`
		Code  int              `json:"code"`
	}

	// all transactions
	t.Run("All Transactions", func(t *testing.T) {
		serviceMock.On("AllTransactions", mockContext, fromAccount).
			Return(ent.Transactions{
				{
					ID:        5,
					Amount:    85,
					CreatedAt: tNow,
					Edges: ent.TransactionEdges{
						FromAccount: fromAccount,
						ToAccount:   toAccount,
					},
				},
				{
					ID:        6,
					Amount:    65,
					CreatedAt: tNow,
					Edges: ent.TransactionEdges{
						FromAccount: toAccount,
						ToAccount:   fromAccount,
					},
				},
				{
					ID:        7,
					Amount:    35,
					CreatedAt: tNow,
					Edges: ent.TransactionEdges{
						FromAccount: fromAccount,
						ToAccount:   toAccount,
					},
				},
			}, nil)
		req, _ := http.NewRequest(http.MethodGet, "/accounts/1/transactions", nil)
		resp := executeMockRequest(req, s)

		require.Equal(t, http.StatusOK, resp.Result().StatusCode)

		rBody, _ := io.ReadAll(resp.Result().Body)
		acc := &ResponseWrapper{}
		err := json.Unmarshal(rBody, acc)

		require.Equal(t, nil, err)
		require.Equal(t, "", acc.Error)
		require.Equal(t, 200, acc.Code)
		require.Equal(t, 3, len(acc.Data))
		require.Equal(t, float64(85), acc.Data[0].Amount)
		require.Equal(t, float64(65), acc.Data[1].Amount)
		require.Equal(t, float64(35), acc.Data[2].Amount)
		require.Equal(t, 1, acc.Data[0].Edges.FromAccount.ID)
		require.Equal(t, 2, acc.Data[0].Edges.ToAccount.ID)
		require.Equal(t, 2, acc.Data[1].Edges.FromAccount.ID)
		require.Equal(t, 1, acc.Data[1].Edges.ToAccount.ID)
		require.Equal(t, 1, acc.Data[2].Edges.FromAccount.ID)
		require.Equal(t, 2, acc.Data[2].Edges.ToAccount.ID)
	})
	// // incoming transactions

	t.Run("Incoming Transactions", func(t *testing.T) {
		serviceMock.On("IncomingTransactions", mockContext, fromAccount).
			Return(ent.Transactions{
				{
					ID:        1,
					Amount:    10,
					CreatedAt: tNow,
					Edges: ent.TransactionEdges{
						FromAccount: toAccount,
						ToAccount:   fromAccount,
					},
				},
				{
					ID:        3,
					Amount:    25,
					CreatedAt: tNow,
					Edges: ent.TransactionEdges{
						FromAccount: toAccount,
						ToAccount:   fromAccount,
					},
				},
			}, nil)
		req, _ := http.NewRequest(http.MethodGet, "/accounts/1/transactions?type=incoming", nil)
		resp := executeMockRequest(req, s)

		require.Equal(t, http.StatusOK, resp.Result().StatusCode)

		rBody, _ := io.ReadAll(resp.Result().Body)
		acc := &ResponseWrapper{}
		err := json.Unmarshal(rBody, acc)

		require.Equal(t, nil, err)
		require.Equal(t, "", acc.Error)
		require.Equal(t, 200, acc.Code)
		require.Equal(t, 2, len(acc.Data))
		require.Equal(t, float64(10), acc.Data[0].Amount)
		require.Equal(t, float64(25), acc.Data[1].Amount)
		require.Equal(t, 2, acc.Data[0].Edges.FromAccount.ID)
		require.Equal(t, 1, acc.Data[0].Edges.ToAccount.ID)
		require.Equal(t, 2, acc.Data[1].Edges.FromAccount.ID)
		require.Equal(t, 1, acc.Data[1].Edges.ToAccount.ID)
	})

	// outgoing transactions
	t.Run("Outgoing Transactions", func(t *testing.T) {
		serviceMock.On("OutgoingTransactions", mockContext, fromAccount).
			Return(ent.Transactions{
				{
					ID:        2,
					Amount:    40,
					CreatedAt: tNow,
					Edges: ent.TransactionEdges{
						FromAccount: fromAccount,
						ToAccount:   toAccount,
					},
				},
				{
					ID:        4,
					Amount:    35,
					CreatedAt: tNow,
					Edges: ent.TransactionEdges{
						FromAccount: fromAccount,
						ToAccount:   toAccount,
					},
				},
			}, nil)
		req, _ := http.NewRequest(http.MethodGet, "/accounts/1/transactions?type=outgoing", nil)
		resp := executeMockRequest(req, s)

		require.Equal(t, http.StatusOK, resp.Result().StatusCode)

		rBody, _ := io.ReadAll(resp.Result().Body)
		acc := &ResponseWrapper{}
		err := json.Unmarshal(rBody, acc)

		require.Equal(t, nil, err)
		require.Equal(t, "", acc.Error)
		require.Equal(t, 200, acc.Code)
		require.Equal(t, 2, len(acc.Data))
		require.Equal(t, float64(40), acc.Data[0].Amount)
		require.Equal(t, float64(35), acc.Data[1].Amount)
		require.Equal(t, 1, acc.Data[0].Edges.FromAccount.ID)
		require.Equal(t, 2, acc.Data[0].Edges.ToAccount.ID)
		require.Equal(t, 1, acc.Data[1].Edges.FromAccount.ID)
		require.Equal(t, 2, acc.Data[1].Edges.ToAccount.ID)
	})

}

type mockServer struct {
	router *chi.Mux
}

func createMockServer() *mockServer {
	s := &mockServer{}
	s.router = chi.NewRouter()
	return s
}

func executeMockRequest(req *http.Request, s *mockServer) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	s.router.ServeHTTP(rr, req)

	return rr
}
