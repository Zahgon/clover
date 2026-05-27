package query

import d "github.com/ostafen/clover/v2/document"

// Query represents a generic query which is submitted to a specific collection.
type Query struct {
	collection string
	criteria   Criteria
	limit      int
	skip       int
	sortOpts   []SortOption
}

// NewQuery simply returns the collection with the supplied name. Use it to initialize a new query.
func NewQuery(collection string) *Query { _ = "STUB: not implemented"; return nil }

func (q *Query) copy() *Query { _ = "STUB: not implemented"; return nil }

func (q *Query) satisfy(doc *d.Document) bool { _ = "STUB: not implemented"; return false }

// MatchFunc selects all the documents which satisfy the supplied predicate function.
func (q *Query) MatchFunc(p func(doc *d.Document) bool) *Query {
	_ = "STUB: not implemented"
	return nil
}

// Where returns a new Query which select all the documents fulfilling the provided Criteria.
func (q *Query) Where(c Criteria) *Query { _ = "STUB: not implemented"; return nil }

// Skip sets the query so that the first n documents of the result set are discarded.
func (q *Query) Skip(n int) *Query { _ = "STUB: not implemented"; return nil }

// Limit sets the query q to consider at most n records.
// As a consequence, the FindAll() method will output at most n documents,
// and any integer m returned by Count() will satisfy the condition m <= n.
func (q *Query) Limit(n int) *Query { _ = "STUB: not implemented"; return nil }

// SortOption is used to specify sorting options to the Sort method.
// It consists of a field name and a sorting direction (1 for ascending and -1 for descending).
// Any other positive of negative value (except from 1 and -1) will be equivalent, respectively, to 1 or -1.
// A direction value of 0 (which is also the default value) is assumed to be ascending.
type SortOption struct {
	Field     string
	Direction int
}

func normalizeSortOptions(opts []SortOption) []SortOption { _ = "STUB: not implemented"; return nil }

// Sort sets the query so that the returned documents are sorted according list of options.
func (q *Query) Sort(opts ...SortOption) *Query {
	_ = "STUB: not implemented"
	// by default, documents are sorted documents by "_id" field
	return nil
}

func (q *Query) Collection() string { _ = "STUB: not implemented"; return "" }

func (q *Query) Criteria() Criteria { _ = "STUB: not implemented"; return *new(Criteria) }

func (q *Query) GetLimit() int { _ = "STUB: not implemented"; return 0 }

func (q *Query) GetSkip() int { _ = "STUB: not implemented"; return 0 }

func (q *Query) SortOptions() []SortOption { _ = "STUB: not implemented"; return nil }
