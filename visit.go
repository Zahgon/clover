package clover

import (
	"github.com/ostafen/clover/v2/index"
	"github.com/ostafen/clover/v2/query"
)

type NotFlattenVisitor struct {
}

func (v *NotFlattenVisitor) VisitUnaryCriteria(c *query.UnaryCriteria) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (v *NotFlattenVisitor) VisitBinaryCriteria(c *query.BinaryCriteria) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (v *NotFlattenVisitor) VisitNotCriteria(c *query.NotCriteria) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// extract this into a separate method

func (v *NotFlattenVisitor) removeNotCriteria(c *query.NotCriteria) query.Criteria {
	_ = "STUB: not implemented"
	return *new(query.Criteria)
}

type IndexSelectVisitor struct {
	Fields map[string]*index.Info
}

func (v *IndexSelectVisitor) VisitUnaryCriteria(c *query.UnaryCriteria) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (v *IndexSelectVisitor) VisitBinaryCriteria(c *query.BinaryCriteria) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// select the indexes with the lowest number of queries

func (v *IndexSelectVisitor) VisitNotCriteria(c *query.NotCriteria) interface{} {
	_ = "STUB: not implemented"
	return nil
}

type FieldRangeVisitor struct {
	Fields map[string]bool
}

func NewFieldRangeVisitor(fields []string) *FieldRangeVisitor {
	_ = "STUB: not implemented"
	return nil
}

func (v *FieldRangeVisitor) VisitUnaryCriteria(c *query.UnaryCriteria) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (v *FieldRangeVisitor) VisitBinaryCriteria(c *query.BinaryCriteria) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (v *FieldRangeVisitor) VisitNotCriteria(c *query.NotCriteria) interface{} {
	_ = "STUB: not implemented"
	return nil
}

type CriteriaNormalizeVisitor struct {
	err error
}

func (v *CriteriaNormalizeVisitor) VisitUnaryCriteria(c *query.UnaryCriteria) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (v *CriteriaNormalizeVisitor) VisitBinaryCriteria(c *query.BinaryCriteria) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (v *CriteriaNormalizeVisitor) VisitNotCriteria(c *query.NotCriteria) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func unaryCriteriaToRange(c *query.UnaryCriteria) *index.Range {
	_ = "STUB: not implemented"
	return nil
}
