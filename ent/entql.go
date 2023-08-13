// Code generated by ent, DO NOT EDIT.

package ent

import (
	"zotregistry.io/zot/ent/object"
	"zotregistry.io/zot/ent/predicate"
	"zotregistry.io/zot/ent/spredicate"
	"zotregistry.io/zot/ent/statement"
	"zotregistry.io/zot/ent/subject"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 4)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   object.Table,
			Columns: object.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: object.FieldID,
			},
		},
		Type: "Object",
		Fields: map[string]*sqlgraph.FieldSpec{
			object.FieldObjectType: {Type: field.TypeString, Column: object.FieldObjectType},
			object.FieldObject:     {Type: field.TypeJSON, Column: object.FieldObject},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   spredicate.Table,
			Columns: spredicate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: spredicate.FieldID,
			},
		},
		Type: "Spredicate",
		Fields: map[string]*sqlgraph.FieldSpec{
			spredicate.FieldPredicateType: {Type: field.TypeString, Column: spredicate.FieldPredicateType},
			spredicate.FieldPredicate:     {Type: field.TypeJSON, Column: spredicate.FieldPredicate},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   statement.Table,
			Columns: statement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statement.FieldID,
			},
		},
		Type: "Statement",
		Fields: map[string]*sqlgraph.FieldSpec{
			statement.FieldNamespace: {Type: field.TypeString, Column: statement.FieldNamespace},
			statement.FieldStatement: {Type: field.TypeJSON, Column: statement.FieldStatement},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   subject.Table,
			Columns: subject.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subject.FieldID,
			},
		},
		Type: "Subject",
		Fields: map[string]*sqlgraph.FieldSpec{
			subject.FieldSubjectType: {Type: field.TypeString, Column: subject.FieldSubjectType},
			subject.FieldSubject:     {Type: field.TypeJSON, Column: subject.FieldSubject},
		},
	}
	graph.MustAddE(
		"statement",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   object.StatementTable,
			Columns: object.StatementPrimaryKey,
			Bidi:    false,
		},
		"Object",
		"Statement",
	)
	graph.MustAddE(
		"statement",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   spredicate.StatementTable,
			Columns: spredicate.StatementPrimaryKey,
			Bidi:    false,
		},
		"Spredicate",
		"Statement",
	)
	graph.MustAddE(
		"objects",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statement.ObjectsTable,
			Columns: statement.ObjectsPrimaryKey,
			Bidi:    false,
		},
		"Statement",
		"Object",
	)
	graph.MustAddE(
		"predicates",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statement.PredicatesTable,
			Columns: statement.PredicatesPrimaryKey,
			Bidi:    false,
		},
		"Statement",
		"Spredicate",
	)
	graph.MustAddE(
		"subjects",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statement.SubjectsTable,
			Columns: statement.SubjectsPrimaryKey,
			Bidi:    false,
		},
		"Statement",
		"Subject",
	)
	graph.MustAddE(
		"statement",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subject.StatementTable,
			Columns: subject.StatementPrimaryKey,
			Bidi:    false,
		},
		"Subject",
		"Statement",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (oq *ObjectQuery) addPredicate(pred func(s *sql.Selector)) {
	oq.predicates = append(oq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ObjectQuery builder.
func (oq *ObjectQuery) Filter() *ObjectFilter {
	return &ObjectFilter{config: oq.config, predicateAdder: oq}
}

// addPredicate implements the predicateAdder interface.
func (m *ObjectMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ObjectMutation builder.
func (m *ObjectMutation) Filter() *ObjectFilter {
	return &ObjectFilter{config: m.config, predicateAdder: m}
}

// ObjectFilter provides a generic filtering capability at runtime for ObjectQuery.
type ObjectFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ObjectFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *ObjectFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(object.FieldID))
}

// WhereObjectType applies the entql string predicate on the objectType field.
func (f *ObjectFilter) WhereObjectType(p entql.StringP) {
	f.Where(p.Field(object.FieldObjectType))
}

// WhereObject applies the entql json.RawMessage predicate on the object field.
func (f *ObjectFilter) WhereObject(p entql.BytesP) {
	f.Where(p.Field(object.FieldObject))
}

// WhereHasStatement applies a predicate to check if query has an edge statement.
func (f *ObjectFilter) WhereHasStatement() {
	f.Where(entql.HasEdge("statement"))
}

// WhereHasStatementWith applies a predicate to check if query has an edge statement with a given conditions (other predicates).
func (f *ObjectFilter) WhereHasStatementWith(preds ...predicate.Statement) {
	f.Where(entql.HasEdgeWith("statement", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (sq *SpredicateQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the SpredicateQuery builder.
func (sq *SpredicateQuery) Filter() *SpredicateFilter {
	return &SpredicateFilter{config: sq.config, predicateAdder: sq}
}

// addPredicate implements the predicateAdder interface.
func (m *SpredicateMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the SpredicateMutation builder.
func (m *SpredicateMutation) Filter() *SpredicateFilter {
	return &SpredicateFilter{config: m.config, predicateAdder: m}
}

// SpredicateFilter provides a generic filtering capability at runtime for SpredicateQuery.
type SpredicateFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *SpredicateFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *SpredicateFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(spredicate.FieldID))
}

// WherePredicateType applies the entql string predicate on the predicateType field.
func (f *SpredicateFilter) WherePredicateType(p entql.StringP) {
	f.Where(p.Field(spredicate.FieldPredicateType))
}

// WherePredicate applies the entql json.RawMessage predicate on the predicate field.
func (f *SpredicateFilter) WherePredicate(p entql.BytesP) {
	f.Where(p.Field(spredicate.FieldPredicate))
}

// WhereHasStatement applies a predicate to check if query has an edge statement.
func (f *SpredicateFilter) WhereHasStatement() {
	f.Where(entql.HasEdge("statement"))
}

// WhereHasStatementWith applies a predicate to check if query has an edge statement with a given conditions (other predicates).
func (f *SpredicateFilter) WhereHasStatementWith(preds ...predicate.Statement) {
	f.Where(entql.HasEdgeWith("statement", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (sq *StatementQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the StatementQuery builder.
func (sq *StatementQuery) Filter() *StatementFilter {
	return &StatementFilter{config: sq.config, predicateAdder: sq}
}

// addPredicate implements the predicateAdder interface.
func (m *StatementMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the StatementMutation builder.
func (m *StatementMutation) Filter() *StatementFilter {
	return &StatementFilter{config: m.config, predicateAdder: m}
}

// StatementFilter provides a generic filtering capability at runtime for StatementQuery.
type StatementFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *StatementFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *StatementFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(statement.FieldID))
}

// WhereNamespace applies the entql string predicate on the namespace field.
func (f *StatementFilter) WhereNamespace(p entql.StringP) {
	f.Where(p.Field(statement.FieldNamespace))
}

// WhereStatement applies the entql json.RawMessage predicate on the statement field.
func (f *StatementFilter) WhereStatement(p entql.BytesP) {
	f.Where(p.Field(statement.FieldStatement))
}

// WhereHasObjects applies a predicate to check if query has an edge objects.
func (f *StatementFilter) WhereHasObjects() {
	f.Where(entql.HasEdge("objects"))
}

// WhereHasObjectsWith applies a predicate to check if query has an edge objects with a given conditions (other predicates).
func (f *StatementFilter) WhereHasObjectsWith(preds ...predicate.Object) {
	f.Where(entql.HasEdgeWith("objects", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasPredicates applies a predicate to check if query has an edge predicates.
func (f *StatementFilter) WhereHasPredicates() {
	f.Where(entql.HasEdge("predicates"))
}

// WhereHasPredicatesWith applies a predicate to check if query has an edge predicates with a given conditions (other predicates).
func (f *StatementFilter) WhereHasPredicatesWith(preds ...predicate.Spredicate) {
	f.Where(entql.HasEdgeWith("predicates", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasSubjects applies a predicate to check if query has an edge subjects.
func (f *StatementFilter) WhereHasSubjects() {
	f.Where(entql.HasEdge("subjects"))
}

// WhereHasSubjectsWith applies a predicate to check if query has an edge subjects with a given conditions (other predicates).
func (f *StatementFilter) WhereHasSubjectsWith(preds ...predicate.Subject) {
	f.Where(entql.HasEdgeWith("subjects", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (sq *SubjectQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the SubjectQuery builder.
func (sq *SubjectQuery) Filter() *SubjectFilter {
	return &SubjectFilter{config: sq.config, predicateAdder: sq}
}

// addPredicate implements the predicateAdder interface.
func (m *SubjectMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the SubjectMutation builder.
func (m *SubjectMutation) Filter() *SubjectFilter {
	return &SubjectFilter{config: m.config, predicateAdder: m}
}

// SubjectFilter provides a generic filtering capability at runtime for SubjectQuery.
type SubjectFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *SubjectFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *SubjectFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(subject.FieldID))
}

// WhereSubjectType applies the entql string predicate on the subjectType field.
func (f *SubjectFilter) WhereSubjectType(p entql.StringP) {
	f.Where(p.Field(subject.FieldSubjectType))
}

// WhereSubject applies the entql json.RawMessage predicate on the subject field.
func (f *SubjectFilter) WhereSubject(p entql.BytesP) {
	f.Where(p.Field(subject.FieldSubject))
}

// WhereHasStatement applies a predicate to check if query has an edge statement.
func (f *SubjectFilter) WhereHasStatement() {
	f.Where(entql.HasEdge("statement"))
}

// WhereHasStatementWith applies a predicate to check if query has an edge statement with a given conditions (other predicates).
func (f *SubjectFilter) WhereHasStatementWith(preds ...predicate.Statement) {
	f.Where(entql.HasEdgeWith("statement", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
