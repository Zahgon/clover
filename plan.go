package clover

import (
	d "github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/index"
	"github.com/ostafen/clover/v2/query"
	"github.com/ostafen/clover/v2/store"
)

type planNode interface {
	SetNext(next planNode)
	NextNode() planNode
	Callback(doc *d.Document) error
	Finish() error
}

type inputNode interface {
	planNode
	Run(tx store.Tx) error
}

type planNodeBase struct {
	next planNode
}

func (nd *planNodeBase) NextNode() planNode { _ = "STUB: not implemented"; return *new(planNode) }

func (nd *planNodeBase) SetNext(next planNode) { _ = "STUB: not implemented"; return }

func (nd *planNodeBase) CallNext(doc *d.Document) error { _ = "STUB: not implemented"; return nil }

func (nd *planNodeBase) Callback(doc *d.Document) error { _ = "STUB: not implemented"; return nil }

func (nd *planNodeBase) Finish() error { _ = "STUB: not implemented"; return nil }

type iterNode struct {
	planNodeBase
	filter     query.Criteria
	collection string

	//vRange     *valueRange
	//index      RangeIndex

	idxQuery index.Query
	//iterIndexReverse bool
}

func (nd *iterNode) iterateFullCollection(tx store.Tx) error { _ = "STUB: not implemented"; return nil }

func (nd *iterNode) iterateIndex(tx store.Tx) error { _ = "STUB: not implemented"; return nil }

// doc == nil when index record expires after document record

func (nd *iterNode) Run(tx store.Tx) error { _ = "STUB: not implemented"; return nil }

func getIndexQueries(q *query.Query, indexes []index.Index) []index.Query {
	_ = "STUB: not implemented"
	return nil
}

func tryToSelectIndex(q *query.Query, indexes []index.Index) (*iterNode, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type skipLimitNode struct {
	planNodeBase
	skipped  int
	consumed int
	skip     int
	limit    int
}

func (nd *skipLimitNode) Callback(doc *d.Document) error { _ = "STUB: not implemented"; return nil }

type sortNode struct {
	planNodeBase
	opts []query.SortOption
	docs []*d.Document
}

func (nd *sortNode) Callback(doc *d.Document) error { _ = "STUB: not implemented"; return nil }

func (nd *sortNode) Finish() error { _ = "STUB: not implemented"; return nil }

func buildQueryPlan(q *query.Query, indexes []index.Index, outputNode planNode) inputNode {
	_ = "STUB: not implemented"
	return *new(inputNode)
}

//isOutputSorted := (len(q.sortOpts) == 1 && itNode.index != nil && itNode.index.Field() == q.sortOpts[0].Field)

//log.Println("output sorted: ", len(q.SortOptions()) > 0 && !isOutputSorted)

func execPlan(nd inputNode, tx store.Tx) error { _ = "STUB: not implemented"; return nil }

type consumerNode struct {
	planNodeBase
	consumer docConsumer
}

func (nd *consumerNode) Callback(doc *d.Document) error { _ = "STUB: not implemented"; return nil }

func compareDocuments(first *d.Document, second *d.Document, sortOpts []query.SortOption) int {
	_ = "STUB: not implemented"
	return 0
}
