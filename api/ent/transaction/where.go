// Code generated by ent, DO NOT EDIT.

package transaction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/richmondwang/golang-wallet-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldID, id))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldAmount, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCreatedAt, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldAmount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldCreatedAt, v))
}

// HasFromAccount applies the HasEdge predicate on the "from_account" edge.
func HasFromAccount() predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FromAccountTable, FromAccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFromAccountWith applies the HasEdge predicate on the "from_account" edge with a given conditions (other predicates).
func HasFromAccountWith(preds ...predicate.Account) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := newFromAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasToAccount applies the HasEdge predicate on the "to_account" edge.
func HasToAccount() predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ToAccountTable, ToAccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasToAccountWith applies the HasEdge predicate on the "to_account" edge with a given conditions (other predicates).
func HasToAccountWith(preds ...predicate.Account) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := newToAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.NotPredicates(p))
}