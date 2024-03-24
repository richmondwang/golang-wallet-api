// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/richmondwang/golang-wallet-api/ent/account"
	"github.com/richmondwang/golang-wallet-api/ent/transaction"
)

// Transaction is the model entity for the Transaction schema.
type Transaction struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionQuery when eager-loading is set.
	Edges                         TransactionEdges `json:"edges"`
	account_outgoing_transactions *int
	account_incoming_transactions *int
	selectValues                  sql.SelectValues
}

// TransactionEdges holds the relations/edges for other nodes in the graph.
type TransactionEdges struct {
	// FromAccount holds the value of the from_account edge.
	FromAccount *Account `json:"from_account,omitempty"`
	// ToAccount holds the value of the to_account edge.
	ToAccount *Account `json:"to_account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FromAccountOrErr returns the FromAccount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) FromAccountOrErr() (*Account, error) {
	if e.FromAccount != nil {
		return e.FromAccount, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: account.Label}
	}
	return nil, &NotLoadedError{edge: "from_account"}
}

// ToAccountOrErr returns the ToAccount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) ToAccountOrErr() (*Account, error) {
	if e.ToAccount != nil {
		return e.ToAccount, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: account.Label}
	}
	return nil, &NotLoadedError{edge: "to_account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transaction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transaction.FieldAmount:
			values[i] = new(sql.NullFloat64)
		case transaction.FieldID:
			values[i] = new(sql.NullInt64)
		case transaction.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case transaction.ForeignKeys[0]: // account_outgoing_transactions
			values[i] = new(sql.NullInt64)
		case transaction.ForeignKeys[1]: // account_incoming_transactions
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transaction fields.
func (t *Transaction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case transaction.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				t.Amount = value.Float64
			}
		case transaction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case transaction.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field account_outgoing_transactions", value)
			} else if value.Valid {
				t.account_outgoing_transactions = new(int)
				*t.account_outgoing_transactions = int(value.Int64)
			}
		case transaction.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field account_incoming_transactions", value)
			} else if value.Valid {
				t.account_incoming_transactions = new(int)
				*t.account_incoming_transactions = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Transaction.
// This includes values selected through modifiers, order, etc.
func (t *Transaction) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryFromAccount queries the "from_account" edge of the Transaction entity.
func (t *Transaction) QueryFromAccount() *AccountQuery {
	return NewTransactionClient(t.config).QueryFromAccount(t)
}

// QueryToAccount queries the "to_account" edge of the Transaction entity.
func (t *Transaction) QueryToAccount() *AccountQuery {
	return NewTransactionClient(t.config).QueryToAccount(t)
}

// Update returns a builder for updating this Transaction.
// Note that you need to call Transaction.Unwrap() before calling this method if this Transaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transaction) Update() *TransactionUpdateOne {
	return NewTransactionClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Transaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transaction) Unwrap() *Transaction {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transaction is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transaction) String() string {
	var builder strings.Builder
	builder.WriteString("Transaction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", t.Amount))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Transactions is a parsable slice of Transaction.
type Transactions []*Transaction