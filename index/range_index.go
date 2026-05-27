package index

import (
	"time"

	"github.com/ostafen/clover/v2/store"
)

type RangeIndex interface {
	Index
	IterateRange(vRange *Range, reverse bool, onValue func(docId string) error) error
}

type RangeIndexQuery struct {
	Range   *Range
	Reverse bool
	Idx     RangeIndex
}

func (q *RangeIndexQuery) Run(onValue func(docId string) error) error {
	_ = "STUB: not implemented"
	return nil
}

type rangeIndex struct {
	indexBase
	tx store.Tx
}

func extractDocId(key []byte) ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

func (idx *rangeIndex) getKeyPrefix() []byte { _ = "STUB: not implemented"; return nil }

func (idx *rangeIndex) getKeyPrefixForType(typeId int) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (idx *rangeIndex) getKey(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *rangeIndex) encodeValueAndId(value interface{}, docId string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *rangeIndex) Add(docId string, v interface{}, ttl time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (idx *rangeIndex) Remove(docId string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (idx *rangeIndex) Drop() error { _ = "STUB: not implemented"; return nil }

func (idx *rangeIndex) encodeRange(vRange *Range) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (idx *rangeIndex) IterateRange(vRange *Range, reverse bool, onValue func(docId string) error) error {
	_ = "STUB: not implemented"
	return nil
}

// skip all values equals to range.start

// skip all values equals to range.end

func (idx *rangeIndex) Iterate(reverse bool, onValue func(docId string) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (idx *rangeIndex) Type() Type { _ = "STUB: not implemented"; return *new(Type) }
