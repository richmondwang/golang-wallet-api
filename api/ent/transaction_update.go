// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/richmondwang/golang-wallet-api/ent/account"
	"github.com/richmondwang/golang-wallet-api/ent/predicate"
	"github.com/richmondwang/golang-wallet-api/ent/transaction"
)

// TransactionUpdate is the builder for updating Transaction entities.
type TransactionUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionMutation
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tu *TransactionUpdate) Where(ps ...predicate.Transaction) *TransactionUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetAmount sets the "amount" field.
func (tu *TransactionUpdate) SetAmount(f float64) *TransactionUpdate {
	tu.mutation.ResetAmount()
	tu.mutation.SetAmount(f)
	return tu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableAmount(f *float64) *TransactionUpdate {
	if f != nil {
		tu.SetAmount(*f)
	}
	return tu
}

// AddAmount adds f to the "amount" field.
func (tu *TransactionUpdate) AddAmount(f float64) *TransactionUpdate {
	tu.mutation.AddAmount(f)
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TransactionUpdate) SetCreatedAt(t time.Time) *TransactionUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableCreatedAt(t *time.Time) *TransactionUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetFromAccountID sets the "from_account" edge to the Account entity by ID.
func (tu *TransactionUpdate) SetFromAccountID(id int) *TransactionUpdate {
	tu.mutation.SetFromAccountID(id)
	return tu
}

// SetFromAccount sets the "from_account" edge to the Account entity.
func (tu *TransactionUpdate) SetFromAccount(a *Account) *TransactionUpdate {
	return tu.SetFromAccountID(a.ID)
}

// SetToAccountID sets the "to_account" edge to the Account entity by ID.
func (tu *TransactionUpdate) SetToAccountID(id int) *TransactionUpdate {
	tu.mutation.SetToAccountID(id)
	return tu
}

// SetToAccount sets the "to_account" edge to the Account entity.
func (tu *TransactionUpdate) SetToAccount(a *Account) *TransactionUpdate {
	return tu.SetToAccountID(a.ID)
}

// Mutation returns the TransactionMutation object of the builder.
func (tu *TransactionUpdate) Mutation() *TransactionMutation {
	return tu.mutation
}

// ClearFromAccount clears the "from_account" edge to the Account entity.
func (tu *TransactionUpdate) ClearFromAccount() *TransactionUpdate {
	tu.mutation.ClearFromAccount()
	return tu
}

// ClearToAccount clears the "to_account" edge to the Account entity.
func (tu *TransactionUpdate) ClearToAccount() *TransactionUpdate {
	tu.mutation.ClearToAccount()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TransactionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TransactionUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TransactionUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TransactionUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TransactionUpdate) check() error {
	if _, ok := tu.mutation.FromAccountID(); tu.mutation.FromAccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Transaction.from_account"`)
	}
	if _, ok := tu.mutation.ToAccountID(); tu.mutation.ToAccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Transaction.to_account"`)
	}
	return nil
}

func (tu *TransactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(transaction.Table, transaction.Columns, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Amount(); ok {
		_spec.SetField(transaction.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.AddedAmount(); ok {
		_spec.AddField(transaction.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(transaction.FieldCreatedAt, field.TypeTime, value)
	}
	if tu.mutation.FromAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.FromAccountTable,
			Columns: []string{transaction.FromAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.FromAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.FromAccountTable,
			Columns: []string{transaction.FromAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ToAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.ToAccountTable,
			Columns: []string{transaction.ToAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ToAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.ToAccountTable,
			Columns: []string{transaction.ToAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TransactionUpdateOne is the builder for updating a single Transaction entity.
type TransactionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransactionMutation
}

// SetAmount sets the "amount" field.
func (tuo *TransactionUpdateOne) SetAmount(f float64) *TransactionUpdateOne {
	tuo.mutation.ResetAmount()
	tuo.mutation.SetAmount(f)
	return tuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableAmount(f *float64) *TransactionUpdateOne {
	if f != nil {
		tuo.SetAmount(*f)
	}
	return tuo
}

// AddAmount adds f to the "amount" field.
func (tuo *TransactionUpdateOne) AddAmount(f float64) *TransactionUpdateOne {
	tuo.mutation.AddAmount(f)
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TransactionUpdateOne) SetCreatedAt(t time.Time) *TransactionUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableCreatedAt(t *time.Time) *TransactionUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetFromAccountID sets the "from_account" edge to the Account entity by ID.
func (tuo *TransactionUpdateOne) SetFromAccountID(id int) *TransactionUpdateOne {
	tuo.mutation.SetFromAccountID(id)
	return tuo
}

// SetFromAccount sets the "from_account" edge to the Account entity.
func (tuo *TransactionUpdateOne) SetFromAccount(a *Account) *TransactionUpdateOne {
	return tuo.SetFromAccountID(a.ID)
}

// SetToAccountID sets the "to_account" edge to the Account entity by ID.
func (tuo *TransactionUpdateOne) SetToAccountID(id int) *TransactionUpdateOne {
	tuo.mutation.SetToAccountID(id)
	return tuo
}

// SetToAccount sets the "to_account" edge to the Account entity.
func (tuo *TransactionUpdateOne) SetToAccount(a *Account) *TransactionUpdateOne {
	return tuo.SetToAccountID(a.ID)
}

// Mutation returns the TransactionMutation object of the builder.
func (tuo *TransactionUpdateOne) Mutation() *TransactionMutation {
	return tuo.mutation
}

// ClearFromAccount clears the "from_account" edge to the Account entity.
func (tuo *TransactionUpdateOne) ClearFromAccount() *TransactionUpdateOne {
	tuo.mutation.ClearFromAccount()
	return tuo
}

// ClearToAccount clears the "to_account" edge to the Account entity.
func (tuo *TransactionUpdateOne) ClearToAccount() *TransactionUpdateOne {
	tuo.mutation.ClearToAccount()
	return tuo
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tuo *TransactionUpdateOne) Where(ps ...predicate.Transaction) *TransactionUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TransactionUpdateOne) Select(field string, fields ...string) *TransactionUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Transaction entity.
func (tuo *TransactionUpdateOne) Save(ctx context.Context) (*Transaction, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TransactionUpdateOne) SaveX(ctx context.Context) *Transaction {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TransactionUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TransactionUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TransactionUpdateOne) check() error {
	if _, ok := tuo.mutation.FromAccountID(); tuo.mutation.FromAccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Transaction.from_account"`)
	}
	if _, ok := tuo.mutation.ToAccountID(); tuo.mutation.ToAccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Transaction.to_account"`)
	}
	return nil
}

func (tuo *TransactionUpdateOne) sqlSave(ctx context.Context) (_node *Transaction, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(transaction.Table, transaction.Columns, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Transaction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transaction.FieldID)
		for _, f := range fields {
			if !transaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Amount(); ok {
		_spec.SetField(transaction.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.AddedAmount(); ok {
		_spec.AddField(transaction.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(transaction.FieldCreatedAt, field.TypeTime, value)
	}
	if tuo.mutation.FromAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.FromAccountTable,
			Columns: []string{transaction.FromAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.FromAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.FromAccountTable,
			Columns: []string{transaction.FromAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ToAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.ToAccountTable,
			Columns: []string{transaction.ToAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ToAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.ToAccountTable,
			Columns: []string{transaction.ToAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Transaction{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
