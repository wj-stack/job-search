// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wj-stack/job-search/ent/application"
	"github.com/wj-stack/job-search/ent/predicate"
)

// ApplicationUpdate is the builder for updating Application entities.
type ApplicationUpdate struct {
	config
	hooks    []Hook
	mutation *ApplicationMutation
}

// Where appends a list predicates to the ApplicationUpdate builder.
func (au *ApplicationUpdate) Where(ps ...predicate.Application) *ApplicationUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUserID sets the "user_id" field.
func (au *ApplicationUpdate) SetUserID(i int) *ApplicationUpdate {
	au.mutation.ResetUserID()
	au.mutation.SetUserID(i)
	return au
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableUserID(i *int) *ApplicationUpdate {
	if i != nil {
		au.SetUserID(*i)
	}
	return au
}

// AddUserID adds i to the "user_id" field.
func (au *ApplicationUpdate) AddUserID(i int) *ApplicationUpdate {
	au.mutation.AddUserID(i)
	return au
}

// SetJobID sets the "job_id" field.
func (au *ApplicationUpdate) SetJobID(i int) *ApplicationUpdate {
	au.mutation.ResetJobID()
	au.mutation.SetJobID(i)
	return au
}

// SetNillableJobID sets the "job_id" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableJobID(i *int) *ApplicationUpdate {
	if i != nil {
		au.SetJobID(*i)
	}
	return au
}

// AddJobID adds i to the "job_id" field.
func (au *ApplicationUpdate) AddJobID(i int) *ApplicationUpdate {
	au.mutation.AddJobID(i)
	return au
}

// SetResumeURL sets the "resume_url" field.
func (au *ApplicationUpdate) SetResumeURL(s string) *ApplicationUpdate {
	au.mutation.SetResumeURL(s)
	return au
}

// SetNillableResumeURL sets the "resume_url" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableResumeURL(s *string) *ApplicationUpdate {
	if s != nil {
		au.SetResumeURL(*s)
	}
	return au
}

// SetStatus sets the "status" field.
func (au *ApplicationUpdate) SetStatus(s string) *ApplicationUpdate {
	au.mutation.SetStatus(s)
	return au
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableStatus(s *string) *ApplicationUpdate {
	if s != nil {
		au.SetStatus(*s)
	}
	return au
}

// Mutation returns the ApplicationMutation object of the builder.
func (au *ApplicationUpdate) Mutation() *ApplicationMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ApplicationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ApplicationUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ApplicationUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ApplicationUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ApplicationUpdate) check() error {
	if v, ok := au.mutation.ResumeURL(); ok {
		if err := application.ResumeURLValidator(v); err != nil {
			return &ValidationError{Name: "resume_url", err: fmt.Errorf(`ent: validator failed for field "Application.resume_url": %w`, err)}
		}
	}
	if v, ok := au.mutation.Status(); ok {
		if err := application.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Application.status": %w`, err)}
		}
	}
	return nil
}

func (au *ApplicationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(application.Table, application.Columns, sqlgraph.NewFieldSpec(application.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UserID(); ok {
		_spec.SetField(application.FieldUserID, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedUserID(); ok {
		_spec.AddField(application.FieldUserID, field.TypeInt, value)
	}
	if value, ok := au.mutation.JobID(); ok {
		_spec.SetField(application.FieldJobID, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedJobID(); ok {
		_spec.AddField(application.FieldJobID, field.TypeInt, value)
	}
	if value, ok := au.mutation.ResumeURL(); ok {
		_spec.SetField(application.FieldResumeURL, field.TypeString, value)
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.SetField(application.FieldStatus, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{application.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ApplicationUpdateOne is the builder for updating a single Application entity.
type ApplicationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApplicationMutation
}

// SetUserID sets the "user_id" field.
func (auo *ApplicationUpdateOne) SetUserID(i int) *ApplicationUpdateOne {
	auo.mutation.ResetUserID()
	auo.mutation.SetUserID(i)
	return auo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableUserID(i *int) *ApplicationUpdateOne {
	if i != nil {
		auo.SetUserID(*i)
	}
	return auo
}

// AddUserID adds i to the "user_id" field.
func (auo *ApplicationUpdateOne) AddUserID(i int) *ApplicationUpdateOne {
	auo.mutation.AddUserID(i)
	return auo
}

// SetJobID sets the "job_id" field.
func (auo *ApplicationUpdateOne) SetJobID(i int) *ApplicationUpdateOne {
	auo.mutation.ResetJobID()
	auo.mutation.SetJobID(i)
	return auo
}

// SetNillableJobID sets the "job_id" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableJobID(i *int) *ApplicationUpdateOne {
	if i != nil {
		auo.SetJobID(*i)
	}
	return auo
}

// AddJobID adds i to the "job_id" field.
func (auo *ApplicationUpdateOne) AddJobID(i int) *ApplicationUpdateOne {
	auo.mutation.AddJobID(i)
	return auo
}

// SetResumeURL sets the "resume_url" field.
func (auo *ApplicationUpdateOne) SetResumeURL(s string) *ApplicationUpdateOne {
	auo.mutation.SetResumeURL(s)
	return auo
}

// SetNillableResumeURL sets the "resume_url" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableResumeURL(s *string) *ApplicationUpdateOne {
	if s != nil {
		auo.SetResumeURL(*s)
	}
	return auo
}

// SetStatus sets the "status" field.
func (auo *ApplicationUpdateOne) SetStatus(s string) *ApplicationUpdateOne {
	auo.mutation.SetStatus(s)
	return auo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableStatus(s *string) *ApplicationUpdateOne {
	if s != nil {
		auo.SetStatus(*s)
	}
	return auo
}

// Mutation returns the ApplicationMutation object of the builder.
func (auo *ApplicationUpdateOne) Mutation() *ApplicationMutation {
	return auo.mutation
}

// Where appends a list predicates to the ApplicationUpdate builder.
func (auo *ApplicationUpdateOne) Where(ps ...predicate.Application) *ApplicationUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ApplicationUpdateOne) Select(field string, fields ...string) *ApplicationUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Application entity.
func (auo *ApplicationUpdateOne) Save(ctx context.Context) (*Application, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ApplicationUpdateOne) SaveX(ctx context.Context) *Application {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ApplicationUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ApplicationUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ApplicationUpdateOne) check() error {
	if v, ok := auo.mutation.ResumeURL(); ok {
		if err := application.ResumeURLValidator(v); err != nil {
			return &ValidationError{Name: "resume_url", err: fmt.Errorf(`ent: validator failed for field "Application.resume_url": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Status(); ok {
		if err := application.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Application.status": %w`, err)}
		}
	}
	return nil
}

func (auo *ApplicationUpdateOne) sqlSave(ctx context.Context) (_node *Application, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(application.Table, application.Columns, sqlgraph.NewFieldSpec(application.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Application.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, application.FieldID)
		for _, f := range fields {
			if !application.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != application.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UserID(); ok {
		_spec.SetField(application.FieldUserID, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedUserID(); ok {
		_spec.AddField(application.FieldUserID, field.TypeInt, value)
	}
	if value, ok := auo.mutation.JobID(); ok {
		_spec.SetField(application.FieldJobID, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedJobID(); ok {
		_spec.AddField(application.FieldJobID, field.TypeInt, value)
	}
	if value, ok := auo.mutation.ResumeURL(); ok {
		_spec.SetField(application.FieldResumeURL, field.TypeString, value)
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.SetField(application.FieldStatus, field.TypeString, value)
	}
	_node = &Application{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{application.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
