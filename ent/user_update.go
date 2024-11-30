// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nevissGo/ent/hype"
	"nevissGo/ent/pixel"
	"nevissGo/ent/predicate"
	"nevissGo/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetDisplayName sets the "display_name" field.
func (uu *UserUpdate) SetDisplayName(s string) *UserUpdate {
	uu.mutation.SetDisplayName(s)
	return uu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDisplayName(s *string) *UserUpdate {
	if s != nil {
		uu.SetDisplayName(*s)
	}
	return uu
}

// SetGameID sets the "game_id" field.
func (uu *UserUpdate) SetGameID(s string) *UserUpdate {
	uu.mutation.SetGameID(s)
	return uu
}

// SetNillableGameID sets the "game_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableGameID(s *string) *UserUpdate {
	if s != nil {
		uu.SetGameID(*s)
	}
	return uu
}

// AddPixelIDs adds the "pixels" edge to the Pixel entity by IDs.
func (uu *UserUpdate) AddPixelIDs(ids ...int) *UserUpdate {
	uu.mutation.AddPixelIDs(ids...)
	return uu
}

// AddPixels adds the "pixels" edges to the Pixel entity.
func (uu *UserUpdate) AddPixels(p ...*Pixel) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPixelIDs(ids...)
}

// SetHypeID sets the "hype" edge to the Hype entity by ID.
func (uu *UserUpdate) SetHypeID(id int) *UserUpdate {
	uu.mutation.SetHypeID(id)
	return uu
}

// SetNillableHypeID sets the "hype" edge to the Hype entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableHypeID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetHypeID(*id)
	}
	return uu
}

// SetHype sets the "hype" edge to the Hype entity.
func (uu *UserUpdate) SetHype(h *Hype) *UserUpdate {
	return uu.SetHypeID(h.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearPixels clears all "pixels" edges to the Pixel entity.
func (uu *UserUpdate) ClearPixels() *UserUpdate {
	uu.mutation.ClearPixels()
	return uu
}

// RemovePixelIDs removes the "pixels" edge to Pixel entities by IDs.
func (uu *UserUpdate) RemovePixelIDs(ids ...int) *UserUpdate {
	uu.mutation.RemovePixelIDs(ids...)
	return uu
}

// RemovePixels removes "pixels" edges to Pixel entities.
func (uu *UserUpdate) RemovePixels(p ...*Pixel) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePixelIDs(ids...)
}

// ClearHype clears the "hype" edge to the Hype entity.
func (uu *UserUpdate) ClearHype() *UserUpdate {
	uu.mutation.ClearHype()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.DisplayName(); ok {
		_spec.SetField(user.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := uu.mutation.GameID(); ok {
		_spec.SetField(user.FieldGameID, field.TypeString, value)
	}
	if uu.mutation.PixelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PixelsTable,
			Columns: []string{user.PixelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pixel.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPixelsIDs(); len(nodes) > 0 && !uu.mutation.PixelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PixelsTable,
			Columns: []string{user.PixelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pixel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PixelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PixelsTable,
			Columns: []string{user.PixelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pixel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.HypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.HypeTable,
			Columns: []string{user.HypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.HypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.HypeTable,
			Columns: []string{user.HypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetDisplayName sets the "display_name" field.
func (uuo *UserUpdateOne) SetDisplayName(s string) *UserUpdateOne {
	uuo.mutation.SetDisplayName(s)
	return uuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDisplayName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDisplayName(*s)
	}
	return uuo
}

// SetGameID sets the "game_id" field.
func (uuo *UserUpdateOne) SetGameID(s string) *UserUpdateOne {
	uuo.mutation.SetGameID(s)
	return uuo
}

// SetNillableGameID sets the "game_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGameID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetGameID(*s)
	}
	return uuo
}

// AddPixelIDs adds the "pixels" edge to the Pixel entity by IDs.
func (uuo *UserUpdateOne) AddPixelIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddPixelIDs(ids...)
	return uuo
}

// AddPixels adds the "pixels" edges to the Pixel entity.
func (uuo *UserUpdateOne) AddPixels(p ...*Pixel) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPixelIDs(ids...)
}

// SetHypeID sets the "hype" edge to the Hype entity by ID.
func (uuo *UserUpdateOne) SetHypeID(id int) *UserUpdateOne {
	uuo.mutation.SetHypeID(id)
	return uuo
}

// SetNillableHypeID sets the "hype" edge to the Hype entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableHypeID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetHypeID(*id)
	}
	return uuo
}

// SetHype sets the "hype" edge to the Hype entity.
func (uuo *UserUpdateOne) SetHype(h *Hype) *UserUpdateOne {
	return uuo.SetHypeID(h.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearPixels clears all "pixels" edges to the Pixel entity.
func (uuo *UserUpdateOne) ClearPixels() *UserUpdateOne {
	uuo.mutation.ClearPixels()
	return uuo
}

// RemovePixelIDs removes the "pixels" edge to Pixel entities by IDs.
func (uuo *UserUpdateOne) RemovePixelIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemovePixelIDs(ids...)
	return uuo
}

// RemovePixels removes "pixels" edges to Pixel entities.
func (uuo *UserUpdateOne) RemovePixels(p ...*Pixel) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePixelIDs(ids...)
}

// ClearHype clears the "hype" edge to the Hype entity.
func (uuo *UserUpdateOne) ClearHype() *UserUpdateOne {
	uuo.mutation.ClearHype()
	return uuo
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.DisplayName(); ok {
		_spec.SetField(user.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.GameID(); ok {
		_spec.SetField(user.FieldGameID, field.TypeString, value)
	}
	if uuo.mutation.PixelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PixelsTable,
			Columns: []string{user.PixelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pixel.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPixelsIDs(); len(nodes) > 0 && !uuo.mutation.PixelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PixelsTable,
			Columns: []string{user.PixelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pixel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PixelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PixelsTable,
			Columns: []string{user.PixelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pixel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.HypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.HypeTable,
			Columns: []string{user.HypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.HypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.HypeTable,
			Columns: []string{user.HypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
