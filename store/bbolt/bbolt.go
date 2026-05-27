package bbolt

import (
	"github.com/ostafen/clover/v2/store"
	"go.etcd.io/bbolt"
)

type boltStore struct {
	db *bbolt.DB
}

const (
	dbFileName = "data.db"
	rootBucket = "root"
)

func Open(dir string) (store.Store, error) {
	_ = "STUB: not implemented"
	return *new(store.Store), nil
}

func OpenWithOptions(dir string, opts *bbolt.Options) (store.Store, error) {
	_ = "STUB: not implemented"
	return *new(store.Store), nil
}

func (store *boltStore) createRootBucketIfNotExists() error { _ = "STUB: not implemented"; return nil }

func (store *boltStore) Begin(update bool) (store.Tx, error) {
	_ = "STUB: not implemented"
	return *new(store.Tx), nil
}

func (store *boltStore) Close() error { _ = "STUB: not implemented"; return nil }

type boltTx struct {
	*bbolt.Tx
}

func (tx *boltTx) bucket() *bbolt.Bucket { _ = "STUB: not implemented"; return nil }

func (tx *boltTx) Set(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (tx *boltTx) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (tx *boltTx) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (tx *boltTx) Cursor(forward bool) (store.Cursor, error) {
	_ = "STUB: not implemented"
	return *new(store.Cursor), nil
}

func (tx *boltTx) Commit() error { _ = "STUB: not implemented"; return nil }

func (tx *boltTx) Rollback() error { _ = "STUB: not implemented"; return nil }

type boltCursor struct {
	*bbolt.Cursor
	forward bool

	currItem *store.Item
}

func (c *boltCursor) Seek(seek []byte) error { _ = "STUB: not implemented"; return nil }

func (c *boltCursor) adjustSeek(key []byte, seek []byte) { _ = "STUB: not implemented"; return }

func (c *boltCursor) Next() { _ = "STUB: not implemented"; return }

func (c *boltCursor) Valid() bool { _ = "STUB: not implemented"; return false }

func (c *boltCursor) Item() (store.Item, error) {
	_ = "STUB: not implemented"
	return *new(store.Item), nil
}

func (c *boltCursor) Close() error { _ = "STUB: not implemented"; return nil }
