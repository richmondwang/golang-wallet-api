package services

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/sql"
	"github.com/richmondwang/golang-wallet-api/ent"
	"github.com/richmondwang/golang-wallet-api/ent/account"
	"github.com/richmondwang/golang-wallet-api/ent/transaction"
)

// Accounts service for accounts db
type Accounts interface {
	// Get retrieve a single account using a user_id
	GetAccount(ctx context.Context, id int) (*ent.Account, error)
	// Add creates a new Account
	AddAccount(ctx context.Context, name string, initialBalance float64) (*ent.Account, error)
	// Transfer deduct balance from account and send it to another account
	Transfer(ctx context.Context, fromAccount *ent.Account, toAccount *ent.Account, amount float64) (*ent.Transaction, error)
	// OutgoingTransactions get all the outgoing transactions of a user
	OutgoingTransactions(ctx context.Context, account *ent.Account) (ent.Transactions, error)
	// IncomingTransactions get all the incoming transactions of a user
	IncomingTransactions(ctx context.Context, account *ent.Account) (ent.Transactions, error)
	// AllTransactions get all the transactions of a user, outgoing and incoming
	AllTransactions(ctx context.Context, account *ent.Account) (ent.Transactions, error)
}

var (
	// ErrInsufficientBalance error for account not having enough balance to tx
	ErrInsufficientBalance error = errors.New("account has insufficient balance")
	// ErrNoName error for account without a name
	ErrNoName = errors.New("account cannot have empty name")
	// ErrNegativeBalance error for account with negative balance
	ErrNegativeBalance = errors.New("account cannot have negative balance")
)

// AccountsService service instance for acounts db
type AccountsService struct {
	db *ent.Client
}

// ensure instance implements the service
var _ Accounts = &AccountsService{}

// NewAccountsService create new accounts service instance
func NewAccountsService(db *ent.Client) *AccountsService {
	return &AccountsService{db: db}
}

// GetAccount retrieve an account using an ID
func (s *AccountsService) GetAccount(ctx context.Context, id int) (*ent.Account, error) {
	account, err := s.db.Account.Get(ctx, id)
	if ent.IsNotFound(err) {
		return nil, nil
	}
	return account, err
}

// AddAccount add an acccount
func (s *AccountsService) AddAccount(ctx context.Context, name string, initialBalance float64) (*ent.Account, error) {
	if name == "" {
		return nil, ErrNoName
	}
	if initialBalance < 0 {
		return nil, ErrNegativeBalance
	}
	account, err := s.db.Account.
		Create().
		SetName(name).
		SetBalance(initialBalance).
		Save(ctx)
	return account, err
}

// Transfer transfer amount from one account to another
func (s *AccountsService) Transfer(ctx context.Context, fromAccount *ent.Account, toAccount *ent.Account, amount float64) (*ent.Transaction, error) {
	if fromAccount.Balance < amount {
		return nil, ErrInsufficientBalance
	}
	var transaction *ent.Transaction
	err := withTx(ctx, s.db, func(tx *ent.Tx) error {
		// lock accounts
		fromAcc, err := tx.Account.
			Query().
			Where(account.ID(fromAccount.ID)).
			ForUpdate(sql.WithLockAction(sql.NoWait)).
			Only(ctx)
		if err != nil {
			return err
		}
		toAcc, err := tx.Account.
			Query().
			Where(account.ID(toAccount.ID)).
			ForUpdate(sql.WithLockAction(sql.NoWait)).
			Only(ctx)
		if err != nil {
			return err
		}

		// create the transfer
		transaction, err = tx.Transaction.
			Create().
			SetFromAccount(fromAcc).
			SetToAccount(toAcc).
			SetAmount(amount).
			Save(ctx)
		if err != nil {
			return err
		}
		_, err = fromAcc.
			Update().
			SetBalance(fromAcc.Balance - amount).
			Save(ctx)
		if err != nil {
			return err
		}
		_, err = toAcc.
			Update().
			SetBalance(toAcc.Balance + amount).
			Save(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return transaction.Unwrap(), nil
}

// OutgoingTransactions retrieve all outgoing transactions of an acount
func (s *AccountsService) OutgoingTransactions(ctx context.Context, acc *ent.Account) (ent.Transactions, error) {
	return acc.QueryOutgoingTransactions().
		WithFromAccount(func(q *ent.AccountQuery) {
			q.Select(account.FieldID, account.FieldName)
		}).
		WithToAccount(func(q *ent.AccountQuery) {
			q.Select(account.FieldID, account.FieldName)
		}).
		All(ctx)
}

// IncomingTransactions retrieve all incoming transactions of an account
func (s *AccountsService) IncomingTransactions(ctx context.Context, acc *ent.Account) (ent.Transactions, error) {
	return acc.QueryIncomingTransactions().
		WithFromAccount(func(q *ent.AccountQuery) {
			q.Select(account.FieldID, account.FieldName)
		}).
		WithToAccount(func(q *ent.AccountQuery) {
			q.Select(account.FieldID, account.FieldName)
		}).
		All(ctx)
}

// AllTransactions retrieve all transactions of an account
func (s *AccountsService) AllTransactions(ctx context.Context, acc *ent.Account) (ent.Transactions, error) {
	return s.db.Transaction.Query().
		Where(
			transaction.Or(
				transaction.HasFromAccountWith(account.ID(acc.ID)),
				transaction.HasToAccountWith(account.ID(acc.ID)),
			),
		).
		WithFromAccount(func(q *ent.AccountQuery) {
			q.Select(account.FieldID, account.FieldName)
		}).
		WithToAccount(func(q *ent.AccountQuery) {
			q.Select(account.FieldID, account.FieldName)
		}).
		All(ctx)
}
