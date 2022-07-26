// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/prowsuites"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
)

// ProwSuitesCreate is the builder for creating a ProwSuites entity.
type ProwSuitesCreate struct {
	config
	mutation *ProwSuitesMutation
	hooks    []Hook
}

// SetJobID sets the "job_id" field.
func (psc *ProwSuitesCreate) SetJobID(s string) *ProwSuitesCreate {
	psc.mutation.SetJobID(s)
	return psc
}

// SetName sets the "Name" field.
func (psc *ProwSuitesCreate) SetName(s string) *ProwSuitesCreate {
	psc.mutation.SetName(s)
	return psc
}

// SetStatus sets the "Status" field.
func (psc *ProwSuitesCreate) SetStatus(s string) *ProwSuitesCreate {
	psc.mutation.SetStatus(s)
	return psc
}

// SetTime sets the "time" field.
func (psc *ProwSuitesCreate) SetTime(f float64) *ProwSuitesCreate {
	psc.mutation.SetTime(f)
	return psc
}

// SetProwSuitesID sets the "prow_suites" edge to the Repository entity by ID.
func (psc *ProwSuitesCreate) SetProwSuitesID(id uuid.UUID) *ProwSuitesCreate {
	psc.mutation.SetProwSuitesID(id)
	return psc
}

// SetNillableProwSuitesID sets the "prow_suites" edge to the Repository entity by ID if the given value is not nil.
func (psc *ProwSuitesCreate) SetNillableProwSuitesID(id *uuid.UUID) *ProwSuitesCreate {
	if id != nil {
		psc = psc.SetProwSuitesID(*id)
	}
	return psc
}

// SetProwSuites sets the "prow_suites" edge to the Repository entity.
func (psc *ProwSuitesCreate) SetProwSuites(r *Repository) *ProwSuitesCreate {
	return psc.SetProwSuitesID(r.ID)
}

// Mutation returns the ProwSuitesMutation object of the builder.
func (psc *ProwSuitesCreate) Mutation() *ProwSuitesMutation {
	return psc.mutation
}

// Save creates the ProwSuites in the database.
func (psc *ProwSuitesCreate) Save(ctx context.Context) (*ProwSuites, error) {
	var (
		err  error
		node *ProwSuites
	)
	if len(psc.hooks) == 0 {
		if err = psc.check(); err != nil {
			return nil, err
		}
		node, err = psc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProwSuitesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = psc.check(); err != nil {
				return nil, err
			}
			psc.mutation = mutation
			if node, err = psc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(psc.hooks) - 1; i >= 0; i-- {
			if psc.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = psc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (psc *ProwSuitesCreate) SaveX(ctx context.Context) *ProwSuites {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psc *ProwSuitesCreate) Exec(ctx context.Context) error {
	_, err := psc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psc *ProwSuitesCreate) ExecX(ctx context.Context) {
	if err := psc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psc *ProwSuitesCreate) check() error {
	if _, ok := psc.mutation.JobID(); !ok {
		return &ValidationError{Name: "job_id", err: errors.New(`db: missing required field "job_id"`)}
	}
	if _, ok := psc.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`db: missing required field "Name"`)}
	}
	if _, ok := psc.mutation.Status(); !ok {
		return &ValidationError{Name: "Status", err: errors.New(`db: missing required field "Status"`)}
	}
	if _, ok := psc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`db: missing required field "time"`)}
	}
	return nil
}

func (psc *ProwSuitesCreate) sqlSave(ctx context.Context) (*ProwSuites, error) {
	_node, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (psc *ProwSuitesCreate) createSpec() (*ProwSuites, *sqlgraph.CreateSpec) {
	var (
		_node = &ProwSuites{config: psc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: prowsuites.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prowsuites.FieldID,
			},
		}
	)
	if value, ok := psc.mutation.JobID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prowsuites.FieldJobID,
		})
		_node.JobID = value
	}
	if value, ok := psc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prowsuites.FieldName,
		})
		_node.Name = value
	}
	if value, ok := psc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prowsuites.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := psc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: prowsuites.FieldTime,
		})
		_node.Time = value
	}
	if nodes := psc.mutation.ProwSuitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prowsuites.ProwSuitesTable,
			Columns: []string{prowsuites.ProwSuitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.repository_prow_suites = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProwSuitesCreateBulk is the builder for creating many ProwSuites entities in bulk.
type ProwSuitesCreateBulk struct {
	config
	builders []*ProwSuitesCreate
}

// Save creates the ProwSuites entities in the database.
func (pscb *ProwSuitesCreateBulk) Save(ctx context.Context) ([]*ProwSuites, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*ProwSuites, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProwSuitesMutation)
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
					_, err = mutators[i+1].Mutate(root, pscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pscb *ProwSuitesCreateBulk) SaveX(ctx context.Context) []*ProwSuites {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscb *ProwSuitesCreateBulk) Exec(ctx context.Context) error {
	_, err := pscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscb *ProwSuitesCreateBulk) ExecX(ctx context.Context) {
	if err := pscb.Exec(ctx); err != nil {
		panic(err)
	}
}
