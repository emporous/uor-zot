// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"zotregistry.io/zot/ent/element"
)

// Element is the model entity for the Element schema.
type Element struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ResourceType holds the value of the "resourceType" field.
	ResourceType string `json:"resourceType,omitempty"`
	// LocatorType holds the value of the "locatorType" field.
	LocatorType string `json:"locatorType,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ElementQuery when eager-loading is set.
	Edges                ElementEdges `json:"edges"`
	resource_elements    *int
	statement_objects    *int
	statement_predicates *int
	statement_subjects   *int
	statement_statements *int
	selectValues         sql.SelectValues
}

// ElementEdges holds the relations/edges for other nodes in the graph.
type ElementEdges struct {
	// Statements holds the value of the statements edge.
	Statements []*Statement `json:"statements,omitempty"`
	// Resources holds the value of the resources edge.
	Resources []*Resource `json:"resources,omitempty"`
	// Locations holds the value of the locations edge.
	Locations []*Resource `json:"locations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedStatements map[string][]*Statement
	namedResources  map[string][]*Resource
	namedLocations  map[string][]*Resource
}

// StatementsOrErr returns the Statements value or an error if the edge
// was not loaded in eager-loading.
func (e ElementEdges) StatementsOrErr() ([]*Statement, error) {
	if e.loadedTypes[0] {
		return e.Statements, nil
	}
	return nil, &NotLoadedError{edge: "statements"}
}

// ResourcesOrErr returns the Resources value or an error if the edge
// was not loaded in eager-loading.
func (e ElementEdges) ResourcesOrErr() ([]*Resource, error) {
	if e.loadedTypes[1] {
		return e.Resources, nil
	}
	return nil, &NotLoadedError{edge: "resources"}
}

// LocationsOrErr returns the Locations value or an error if the edge
// was not loaded in eager-loading.
func (e ElementEdges) LocationsOrErr() ([]*Resource, error) {
	if e.loadedTypes[2] {
		return e.Locations, nil
	}
	return nil, &NotLoadedError{edge: "locations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Element) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case element.FieldID:
			values[i] = new(sql.NullInt64)
		case element.FieldResourceType, element.FieldLocatorType:
			values[i] = new(sql.NullString)
		case element.ForeignKeys[0]: // resource_elements
			values[i] = new(sql.NullInt64)
		case element.ForeignKeys[1]: // statement_objects
			values[i] = new(sql.NullInt64)
		case element.ForeignKeys[2]: // statement_predicates
			values[i] = new(sql.NullInt64)
		case element.ForeignKeys[3]: // statement_subjects
			values[i] = new(sql.NullInt64)
		case element.ForeignKeys[4]: // statement_statements
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Element fields.
func (e *Element) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case element.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case element.FieldResourceType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resourceType", values[i])
			} else if value.Valid {
				e.ResourceType = value.String
			}
		case element.FieldLocatorType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field locatorType", values[i])
			} else if value.Valid {
				e.LocatorType = value.String
			}
		case element.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field resource_elements", value)
			} else if value.Valid {
				e.resource_elements = new(int)
				*e.resource_elements = int(value.Int64)
			}
		case element.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field statement_objects", value)
			} else if value.Valid {
				e.statement_objects = new(int)
				*e.statement_objects = int(value.Int64)
			}
		case element.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field statement_predicates", value)
			} else if value.Valid {
				e.statement_predicates = new(int)
				*e.statement_predicates = int(value.Int64)
			}
		case element.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field statement_subjects", value)
			} else if value.Valid {
				e.statement_subjects = new(int)
				*e.statement_subjects = int(value.Int64)
			}
		case element.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field statement_statements", value)
			} else if value.Valid {
				e.statement_statements = new(int)
				*e.statement_statements = int(value.Int64)
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Element.
// This includes values selected through modifiers, order, etc.
func (e *Element) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryStatements queries the "statements" edge of the Element entity.
func (e *Element) QueryStatements() *StatementQuery {
	return NewElementClient(e.config).QueryStatements(e)
}

// QueryResources queries the "resources" edge of the Element entity.
func (e *Element) QueryResources() *ResourceQuery {
	return NewElementClient(e.config).QueryResources(e)
}

// QueryLocations queries the "locations" edge of the Element entity.
func (e *Element) QueryLocations() *ResourceQuery {
	return NewElementClient(e.config).QueryLocations(e)
}

// Update returns a builder for updating this Element.
// Note that you need to call Element.Unwrap() before calling this method if this Element
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Element) Update() *ElementUpdateOne {
	return NewElementClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Element entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Element) Unwrap() *Element {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Element is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Element) String() string {
	var builder strings.Builder
	builder.WriteString("Element(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("resourceType=")
	builder.WriteString(e.ResourceType)
	builder.WriteString(", ")
	builder.WriteString("locatorType=")
	builder.WriteString(e.LocatorType)
	builder.WriteByte(')')
	return builder.String()
}

// NamedStatements returns the Statements named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Element) NamedStatements(name string) ([]*Statement, error) {
	if e.Edges.namedStatements == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedStatements[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Element) appendNamedStatements(name string, edges ...*Statement) {
	if e.Edges.namedStatements == nil {
		e.Edges.namedStatements = make(map[string][]*Statement)
	}
	if len(edges) == 0 {
		e.Edges.namedStatements[name] = []*Statement{}
	} else {
		e.Edges.namedStatements[name] = append(e.Edges.namedStatements[name], edges...)
	}
}

// NamedResources returns the Resources named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Element) NamedResources(name string) ([]*Resource, error) {
	if e.Edges.namedResources == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedResources[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Element) appendNamedResources(name string, edges ...*Resource) {
	if e.Edges.namedResources == nil {
		e.Edges.namedResources = make(map[string][]*Resource)
	}
	if len(edges) == 0 {
		e.Edges.namedResources[name] = []*Resource{}
	} else {
		e.Edges.namedResources[name] = append(e.Edges.namedResources[name], edges...)
	}
}

// NamedLocations returns the Locations named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Element) NamedLocations(name string) ([]*Resource, error) {
	if e.Edges.namedLocations == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedLocations[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Element) appendNamedLocations(name string, edges ...*Resource) {
	if e.Edges.namedLocations == nil {
		e.Edges.namedLocations = make(map[string][]*Resource)
	}
	if len(edges) == 0 {
		e.Edges.namedLocations[name] = []*Resource{}
	} else {
		e.Edges.namedLocations[name] = append(e.Edges.namedLocations[name], edges...)
	}
}

// Elements is a parsable slice of Element.
type Elements []*Element
