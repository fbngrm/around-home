// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fbngrm/around-home/ent/material"
	"github.com/fbngrm/around-home/ent/partner"
	"github.com/fbngrm/around-home/ent/predicate"
	"github.com/fbngrm/around-home/pkg/materials"
)

// MaterialUpdate is the builder for updating Material entities.
type MaterialUpdate struct {
	config
	hooks    []Hook
	mutation *MaterialMutation
}

// Where appends a list predicates to the MaterialUpdate builder.
func (mu *MaterialUpdate) Where(ps ...predicate.Material) *MaterialUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetMaterial sets the "material" field.
func (mu *MaterialUpdate) SetMaterial(m materials.Material) *MaterialUpdate {
	mu.mutation.SetMaterial(m)
	return mu
}

// AddPartnerIDs adds the "partners" edge to the Partner entity by IDs.
func (mu *MaterialUpdate) AddPartnerIDs(ids ...int) *MaterialUpdate {
	mu.mutation.AddPartnerIDs(ids...)
	return mu
}

// AddPartners adds the "partners" edges to the Partner entity.
func (mu *MaterialUpdate) AddPartners(p ...*Partner) *MaterialUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.AddPartnerIDs(ids...)
}

// Mutation returns the MaterialMutation object of the builder.
func (mu *MaterialUpdate) Mutation() *MaterialMutation {
	return mu.mutation
}

// ClearPartners clears all "partners" edges to the Partner entity.
func (mu *MaterialUpdate) ClearPartners() *MaterialUpdate {
	mu.mutation.ClearPartners()
	return mu
}

// RemovePartnerIDs removes the "partners" edge to Partner entities by IDs.
func (mu *MaterialUpdate) RemovePartnerIDs(ids ...int) *MaterialUpdate {
	mu.mutation.RemovePartnerIDs(ids...)
	return mu
}

// RemovePartners removes "partners" edges to Partner entities.
func (mu *MaterialUpdate) RemovePartners(p ...*Partner) *MaterialUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.RemovePartnerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MaterialUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MaterialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MaterialUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MaterialUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MaterialUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MaterialUpdate) check() error {
	if v, ok := mu.mutation.Material(); ok {
		if err := material.MaterialValidator(v); err != nil {
			return &ValidationError{Name: "material", err: fmt.Errorf(`ent: validator failed for field "Material.material": %w`, err)}
		}
	}
	return nil
}

func (mu *MaterialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   material.Table,
			Columns: material.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: material.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Material(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: material.FieldMaterial,
		})
	}
	if mu.mutation.PartnersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   material.PartnersTable,
			Columns: material.PartnersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partner.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedPartnersIDs(); len(nodes) > 0 && !mu.mutation.PartnersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   material.PartnersTable,
			Columns: material.PartnersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partner.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.PartnersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   material.PartnersTable,
			Columns: material.PartnersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partner.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{material.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MaterialUpdateOne is the builder for updating a single Material entity.
type MaterialUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MaterialMutation
}

// SetMaterial sets the "material" field.
func (muo *MaterialUpdateOne) SetMaterial(m materials.Material) *MaterialUpdateOne {
	muo.mutation.SetMaterial(m)
	return muo
}

// AddPartnerIDs adds the "partners" edge to the Partner entity by IDs.
func (muo *MaterialUpdateOne) AddPartnerIDs(ids ...int) *MaterialUpdateOne {
	muo.mutation.AddPartnerIDs(ids...)
	return muo
}

// AddPartners adds the "partners" edges to the Partner entity.
func (muo *MaterialUpdateOne) AddPartners(p ...*Partner) *MaterialUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.AddPartnerIDs(ids...)
}

// Mutation returns the MaterialMutation object of the builder.
func (muo *MaterialUpdateOne) Mutation() *MaterialMutation {
	return muo.mutation
}

// ClearPartners clears all "partners" edges to the Partner entity.
func (muo *MaterialUpdateOne) ClearPartners() *MaterialUpdateOne {
	muo.mutation.ClearPartners()
	return muo
}

// RemovePartnerIDs removes the "partners" edge to Partner entities by IDs.
func (muo *MaterialUpdateOne) RemovePartnerIDs(ids ...int) *MaterialUpdateOne {
	muo.mutation.RemovePartnerIDs(ids...)
	return muo
}

// RemovePartners removes "partners" edges to Partner entities.
func (muo *MaterialUpdateOne) RemovePartners(p ...*Partner) *MaterialUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.RemovePartnerIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MaterialUpdateOne) Select(field string, fields ...string) *MaterialUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Material entity.
func (muo *MaterialUpdateOne) Save(ctx context.Context) (*Material, error) {
	var (
		err  error
		node *Material
	)
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MaterialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Material)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MaterialMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MaterialUpdateOne) SaveX(ctx context.Context) *Material {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MaterialUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MaterialUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MaterialUpdateOne) check() error {
	if v, ok := muo.mutation.Material(); ok {
		if err := material.MaterialValidator(v); err != nil {
			return &ValidationError{Name: "material", err: fmt.Errorf(`ent: validator failed for field "Material.material": %w`, err)}
		}
	}
	return nil
}

func (muo *MaterialUpdateOne) sqlSave(ctx context.Context) (_node *Material, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   material.Table,
			Columns: material.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: material.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Material.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, material.FieldID)
		for _, f := range fields {
			if !material.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != material.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Material(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: material.FieldMaterial,
		})
	}
	if muo.mutation.PartnersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   material.PartnersTable,
			Columns: material.PartnersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partner.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedPartnersIDs(); len(nodes) > 0 && !muo.mutation.PartnersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   material.PartnersTable,
			Columns: material.PartnersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partner.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.PartnersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   material.PartnersTable,
			Columns: material.PartnersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partner.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Material{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{material.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}