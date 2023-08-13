// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"zotregistry.io/zot/ent/statement"
	"zotregistry.io/zot/ent/subject"
)

// SubjectCreate is the builder for creating a Subject entity.
type SubjectCreate struct {
	config
	mutation *SubjectMutation
	hooks    []Hook
}

// SetSubjectType sets the "subjectType" field.
func (sc *SubjectCreate) SetSubjectType(s string) *SubjectCreate {
	sc.mutation.SetSubjectType(s)
	return sc
}

// SetSubject sets the "subject" field.
func (sc *SubjectCreate) SetSubject(m map[string]interface{}) *SubjectCreate {
	sc.mutation.SetSubject(m)
	return sc
}

// AddStatementIDs adds the "statement" edge to the Statement entity by IDs.
func (sc *SubjectCreate) AddStatementIDs(ids ...int) *SubjectCreate {
	sc.mutation.AddStatementIDs(ids...)
	return sc
}

// AddStatement adds the "statement" edges to the Statement entity.
func (sc *SubjectCreate) AddStatement(s ...*Statement) *SubjectCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddStatementIDs(ids...)
}

// Mutation returns the SubjectMutation object of the builder.
func (sc *SubjectCreate) Mutation() *SubjectMutation {
	return sc.mutation
}

// Save creates the Subject in the database.
func (sc *SubjectCreate) Save(ctx context.Context) (*Subject, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubjectCreate) SaveX(ctx context.Context) *Subject {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubjectCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubjectCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubjectCreate) check() error {
	if _, ok := sc.mutation.SubjectType(); !ok {
		return &ValidationError{Name: "subjectType", err: errors.New(`ent: missing required field "Subject.subjectType"`)}
	}
	if _, ok := sc.mutation.Subject(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required field "Subject.subject"`)}
	}
	return nil
}

func (sc *SubjectCreate) sqlSave(ctx context.Context) (*Subject, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubjectCreate) createSpec() (*Subject, *sqlgraph.CreateSpec) {
	var (
		_node = &Subject{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(subject.Table, sqlgraph.NewFieldSpec(subject.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.SubjectType(); ok {
		_spec.SetField(subject.FieldSubjectType, field.TypeString, value)
		_node.SubjectType = value
	}
	if value, ok := sc.mutation.Subject(); ok {
		_spec.SetField(subject.FieldSubject, field.TypeJSON, value)
		_node.Subject = value
	}
	if nodes := sc.mutation.StatementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subject.StatementTable,
			Columns: subject.StatementPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statement.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubjectCreateBulk is the builder for creating many Subject entities in bulk.
type SubjectCreateBulk struct {
	config
	builders []*SubjectCreate
}

// Save creates the Subject entities in the database.
func (scb *SubjectCreateBulk) Save(ctx context.Context) ([]*Subject, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subject, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubjectMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubjectCreateBulk) SaveX(ctx context.Context) []*Subject {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubjectCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubjectCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
