package query

import (
	d "github.com/ostafen/clover/v2/document"
)

const (
	ExistsOp = iota
	EqOp
	NeqOp
	GtOp
	GtEqOp
	LtOp
	LtEqOp
	LikeOp
	InOp
	ContainsOp
	FunctionOp
)

const (
	LogicalAnd = iota
	LogicalOr
)

// Criteria represents a predicate for selecting documents.
// It follows a fluent API style so that you can easily chain together multiple criteria.
type Criteria interface {
	Satisfy(doc *d.Document) bool
	Accept(v CriteriaVisitor) interface{}
	Not() Criteria
	And(c Criteria) Criteria
	Or(c Criteria) Criteria
}

type BinaryCriteria struct {
	OpType int
	C1, C2 Criteria
}

func (c *BinaryCriteria) Accept(v CriteriaVisitor) interface{} {
	_ = "STUB: not implemented"
	return nil
}

type NotCriteria struct {
	C Criteria
}

func (c *NotCriteria) Not() Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (c *NotCriteria) And(other Criteria) Criteria {
	_ = "STUB: not implemented"
	return *new(Criteria)
}

func (c *NotCriteria) Or(other Criteria) Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (c *NotCriteria) Satisfy(doc *d.Document) bool { _ = "STUB: not implemented"; return false }

func (c *NotCriteria) Accept(v CriteriaVisitor) interface{} { _ = "STUB: not implemented"; return nil }

func (c *BinaryCriteria) Not() Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (c *BinaryCriteria) And(other Criteria) Criteria {
	_ = "STUB: not implemented"
	return *new(Criteria)
}

func (c *BinaryCriteria) Or(other Criteria) Criteria {
	_ = "STUB: not implemented"
	return *new(Criteria)
}

func (c *BinaryCriteria) Satisfy(doc *d.Document) bool { _ = "STUB: not implemented"; return false }

type UnaryCriteria struct {
	OpType int
	Field  string
	Value  interface{}
}

func (c *UnaryCriteria) Not() Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (c *UnaryCriteria) And(other Criteria) Criteria {
	_ = "STUB: not implemented"
	return *new(Criteria)
}

func (c *UnaryCriteria) Or(other Criteria) Criteria {
	_ = "STUB: not implemented"
	return *new(Criteria)
}

func (c *UnaryCriteria) Satisfy(doc *d.Document) bool { _ = "STUB: not implemented"; return false }

func (c *UnaryCriteria) Accept(v CriteriaVisitor) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func and(c1, c2 Criteria) Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func or(c1, c2 Criteria) Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func not(c Criteria) Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func newCriteria(opType int, field string, value interface{}) Criteria {
	_ = "STUB: not implemented"
	return *new(Criteria)
}

type field struct {
	name string
}

func IsField(v interface{}) bool { _ = "STUB: not implemented"; return false }

// Field represents a document field. It is used to create a new criteria.
func Field(name string) *field { _ = "STUB: not implemented"; return nil }

func (f *field) Exists() Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (f *field) NotExists() Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (f *field) IsNil() Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (f *field) IsTrue() Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (f *field) IsFalse() Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (f *field) IsNilOrNotExists() Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (f *field) Eq(value interface{}) Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (f *field) Gt(value interface{}) Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (f *field) GtEq(value interface{}) Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (f *field) Lt(value interface{}) Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (f *field) LtEq(value interface{}) Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (f *field) Neq(value interface{}) Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (f *field) In(values ...interface{}) Criteria {
	_ = "STUB: not implemented"
	return *new(Criteria)
}

func (f *field) Like(pattern string) Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (f *field) Contains(elems ...interface{}) Criteria {
	_ = "STUB: not implemented"
	return *new(Criteria)
}

// getFieldOrValue returns dereferenced value if value denotes another document field,
// otherwise returns the value itself directly
func getFieldOrValue(doc *d.Document, value interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *UnaryCriteria) compare(doc *d.Document) bool { _ = "STUB: not implemented"; return false }

func (c *UnaryCriteria) exist(doc *d.Document) bool { _ = "STUB: not implemented"; return false }

func (c *UnaryCriteria) eq(doc *d.Document) bool { _ = "STUB: not implemented"; return false }

func (c *UnaryCriteria) in(doc *d.Document) bool { _ = "STUB: not implemented"; return false }

func (c *UnaryCriteria) contains(doc *d.Document) bool { _ = "STUB: not implemented"; return false }

func (c *UnaryCriteria) like(doc *d.Document) bool { _ = "STUB: not implemented"; return false }

type CriteriaVisitor interface {
	VisitUnaryCriteria(c *UnaryCriteria) interface{}
	VisitNotCriteria(c *NotCriteria) interface{}
	VisitBinaryCriteria(c *BinaryCriteria) interface{}
}
