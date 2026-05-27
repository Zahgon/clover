package internal

import (
	"time"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	msgpack.RegisterExt(1, (*LocalizedTime)(nil))
}

type LocalizedTime struct {
	time.Time
}

var _ msgpack.Marshaler = (*LocalizedTime)(nil)
var _ msgpack.Unmarshaler = (*LocalizedTime)(nil)

func (tm *LocalizedTime) MarshalMsgpack() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tm *LocalizedTime) UnmarshalMsgpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func replaceTimes(v interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func removeLocalizedTimes(v interface{}) interface{} { _ = "STUB: not implemented"; return nil }
