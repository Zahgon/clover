package index

type Range struct {
	Start, End                 interface{}
	StartIncluded, EndIncluded bool
}

func (r *Range) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (r *Range) IsNil() bool { _ = "STUB: not implemented"; return false }

func (r *Range) Intersect(r2 *Range) *Range { _ = "STUB: not implemented"; return nil }
