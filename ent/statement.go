// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"zotregistry.io/zot/ent/statement"
)

// Statement is the model entity for the Statement schema.
type Statement struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Namespace holds the value of the "namespace" field.
	Namespace string `json:"namespace,omitempty"`
	// Statement holds the value of the "statement" field.
	Statement map[string]interface{} `json:"statement,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StatementQuery when eager-loading is set.
	Edges        StatementEdges `json:"edges"`
	selectValues sql.SelectValues
}

// StatementEdges holds the relations/edges for other nodes in the graph.
type StatementEdges struct {
	// Objects holds the value of the objects edge.
	Objects []*Object `json:"objects,omitempty"`
	// Predicates holds the value of the predicates edge.
	Predicates []*Spredicate `json:"predicates,omitempty"`
	// Subjects holds the value of the subjects edge.
	Subjects []*Subject `json:"subjects,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedObjects    map[string][]*Object
	namedPredicates map[string][]*Spredicate
	namedSubjects   map[string][]*Subject
}

// ObjectsOrErr returns the Objects value or an error if the edge
// was not loaded in eager-loading.
func (e StatementEdges) ObjectsOrErr() ([]*Object, error) {
	if e.loadedTypes[0] {
		return e.Objects, nil
	}
	return nil, &NotLoadedError{edge: "objects"}
}

// PredicatesOrErr returns the Predicates value or an error if the edge
// was not loaded in eager-loading.
func (e StatementEdges) PredicatesOrErr() ([]*Spredicate, error) {
	if e.loadedTypes[1] {
		return e.Predicates, nil
	}
	return nil, &NotLoadedError{edge: "predicates"}
}

// SubjectsOrErr returns the Subjects value or an error if the edge
// was not loaded in eager-loading.
func (e StatementEdges) SubjectsOrErr() ([]*Subject, error) {
	if e.loadedTypes[2] {
		return e.Subjects, nil
	}
	return nil, &NotLoadedError{edge: "subjects"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Statement) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case statement.FieldStatement:
			values[i] = new([]byte)
		case statement.FieldID:
			values[i] = new(sql.NullInt64)
		case statement.FieldNamespace:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Statement fields.
func (s *Statement) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case statement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case statement.FieldNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace", values[i])
			} else if value.Valid {
				s.Namespace = value.String
			}
		case statement.FieldStatement:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field statement", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Statement); err != nil {
					return fmt.Errorf("unmarshal field statement: %w", err)
				}
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Statement.
// This includes values selected through modifiers, order, etc.
func (s *Statement) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryObjects queries the "objects" edge of the Statement entity.
func (s *Statement) QueryObjects() *ObjectQuery {
	return NewStatementClient(s.config).QueryObjects(s)
}

// QueryPredicates queries the "predicates" edge of the Statement entity.
func (s *Statement) QueryPredicates() *SpredicateQuery {
	return NewStatementClient(s.config).QueryPredicates(s)
}

// QuerySubjects queries the "subjects" edge of the Statement entity.
func (s *Statement) QuerySubjects() *SubjectQuery {
	return NewStatementClient(s.config).QuerySubjects(s)
}

// Update returns a builder for updating this Statement.
// Note that you need to call Statement.Unwrap() before calling this method if this Statement
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Statement) Update() *StatementUpdateOne {
	return NewStatementClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Statement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Statement) Unwrap() *Statement {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Statement is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Statement) String() string {
	var builder strings.Builder
	builder.WriteString("Statement(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("namespace=")
	builder.WriteString(s.Namespace)
	builder.WriteString(", ")
	builder.WriteString("statement=")
	builder.WriteString(fmt.Sprintf("%v", s.Statement))
	builder.WriteByte(')')
	return builder.String()
}

// NamedObjects returns the Objects named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Statement) NamedObjects(name string) ([]*Object, error) {
	if s.Edges.namedObjects == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedObjects[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Statement) appendNamedObjects(name string, edges ...*Object) {
	if s.Edges.namedObjects == nil {
		s.Edges.namedObjects = make(map[string][]*Object)
	}
	if len(edges) == 0 {
		s.Edges.namedObjects[name] = []*Object{}
	} else {
		s.Edges.namedObjects[name] = append(s.Edges.namedObjects[name], edges...)
	}
}

// NamedPredicates returns the Predicates named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Statement) NamedPredicates(name string) ([]*Spredicate, error) {
	if s.Edges.namedPredicates == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedPredicates[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Statement) appendNamedPredicates(name string, edges ...*Spredicate) {
	if s.Edges.namedPredicates == nil {
		s.Edges.namedPredicates = make(map[string][]*Spredicate)
	}
	if len(edges) == 0 {
		s.Edges.namedPredicates[name] = []*Spredicate{}
	} else {
		s.Edges.namedPredicates[name] = append(s.Edges.namedPredicates[name], edges...)
	}
}

// NamedSubjects returns the Subjects named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Statement) NamedSubjects(name string) ([]*Subject, error) {
	if s.Edges.namedSubjects == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedSubjects[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Statement) appendNamedSubjects(name string, edges ...*Subject) {
	if s.Edges.namedSubjects == nil {
		s.Edges.namedSubjects = make(map[string][]*Subject)
	}
	if len(edges) == 0 {
		s.Edges.namedSubjects[name] = []*Subject{}
	} else {
		s.Edges.namedSubjects[name] = append(s.Edges.namedSubjects[name], edges...)
	}
}

// Statements is a parsable slice of Statement.
type Statements []*Statement
