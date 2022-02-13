// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"envoyproxy-dashboard/backend/db/route"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RouteCreate is the builder for creating a Route entity.
type RouteCreate struct {
	config
	mutation *RouteMutation
	hooks    []Hook
}

// SetDomain sets the "domain" field.
func (rc *RouteCreate) SetDomain(s string) *RouteCreate {
	rc.mutation.SetDomain(s)
	return rc
}

// SetPath sets the "path" field.
func (rc *RouteCreate) SetPath(s string) *RouteCreate {
	rc.mutation.SetPath(s)
	return rc
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (rc *RouteCreate) SetNillablePath(s *string) *RouteCreate {
	if s != nil {
		rc.SetPath(*s)
	}
	return rc
}

// Mutation returns the RouteMutation object of the builder.
func (rc *RouteCreate) Mutation() *RouteMutation {
	return rc.mutation
}

// Save creates the Route in the database.
func (rc *RouteCreate) Save(ctx context.Context) (*Route, error) {
	var (
		err  error
		node *Route
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RouteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RouteCreate) SaveX(ctx context.Context) *Route {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RouteCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RouteCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RouteCreate) check() error {
	if _, ok := rc.mutation.Domain(); !ok {
		return &ValidationError{Name: "domain", err: errors.New(`db: missing required field "Route.domain"`)}
	}
	if v, ok := rc.mutation.Domain(); ok {
		if err := route.DomainValidator(v); err != nil {
			return &ValidationError{Name: "domain", err: fmt.Errorf(`db: validator failed for field "Route.domain": %w`, err)}
		}
	}
	return nil
}

func (rc *RouteCreate) sqlSave(ctx context.Context) (*Route, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *RouteCreate) createSpec() (*Route, *sqlgraph.CreateSpec) {
	var (
		_node = &Route{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: route.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: route.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Domain(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: route.FieldDomain,
		})
		_node.Domain = value
	}
	if value, ok := rc.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: route.FieldPath,
		})
		_node.Path = &value
	}
	return _node, _spec
}

// RouteCreateBulk is the builder for creating many Route entities in bulk.
type RouteCreateBulk struct {
	config
	builders []*RouteCreate
}

// Save creates the Route entities in the database.
func (rcb *RouteCreateBulk) Save(ctx context.Context) ([]*Route, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Route, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RouteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RouteCreateBulk) SaveX(ctx context.Context) []*Route {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RouteCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RouteCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}