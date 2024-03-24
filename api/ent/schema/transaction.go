package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Float("amount"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("from_account", Account.Type).
			Ref("outgoing_transactions").
			Unique().
			Required(),
		edge.From("to_account", Account.Type).
			Ref("incoming_transactions").
			Unique().
			Required(),
	}
}

func (Transaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(),
	}
}
