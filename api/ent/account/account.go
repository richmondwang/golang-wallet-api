// Code generated by ent, DO NOT EDIT.

package account

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeOutgoingTransactions holds the string denoting the outgoing_transactions edge name in mutations.
	EdgeOutgoingTransactions = "outgoing_transactions"
	// EdgeIncomingTransactions holds the string denoting the incoming_transactions edge name in mutations.
	EdgeIncomingTransactions = "incoming_transactions"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// OutgoingTransactionsTable is the table that holds the outgoing_transactions relation/edge.
	OutgoingTransactionsTable = "transactions"
	// OutgoingTransactionsInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	OutgoingTransactionsInverseTable = "transactions"
	// OutgoingTransactionsColumn is the table column denoting the outgoing_transactions relation/edge.
	OutgoingTransactionsColumn = "account_outgoing_transactions"
	// IncomingTransactionsTable is the table that holds the incoming_transactions relation/edge.
	IncomingTransactionsTable = "transactions"
	// IncomingTransactionsInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	IncomingTransactionsInverseTable = "transactions"
	// IncomingTransactionsColumn is the table column denoting the incoming_transactions relation/edge.
	IncomingTransactionsColumn = "account_incoming_transactions"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldBalance,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultBalance holds the default value on creation for the "balance" field.
	DefaultBalance float64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Account queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByBalance orders the results by the balance field.
func ByBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBalance, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByOutgoingTransactionsCount orders the results by outgoing_transactions count.
func ByOutgoingTransactionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOutgoingTransactionsStep(), opts...)
	}
}

// ByOutgoingTransactions orders the results by outgoing_transactions terms.
func ByOutgoingTransactions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOutgoingTransactionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByIncomingTransactionsCount orders the results by incoming_transactions count.
func ByIncomingTransactionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIncomingTransactionsStep(), opts...)
	}
}

// ByIncomingTransactions orders the results by incoming_transactions terms.
func ByIncomingTransactions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIncomingTransactionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOutgoingTransactionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OutgoingTransactionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OutgoingTransactionsTable, OutgoingTransactionsColumn),
	)
}
func newIncomingTransactionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IncomingTransactionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IncomingTransactionsTable, IncomingTransactionsColumn),
	)
}
