package badger

import (
	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/ostafen/clover/v2/store"
)

type badgerStore struct {
	db     *badger.DB
	chWg   sync.WaitGroup
	chQuit chan struct{}
}

func (store *badgerStore) Begin(update bool) (store.Tx, error) {
	_ = "STUB: not implemented"
	return *new(store.Tx), nil
}

func (store *badgerStore) Close() error { _ = "STUB: not implemented"; return nil }

type badgerTx struct {
	*badger.Txn
}

func (tx *badgerTx) Set(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func getItemValue(item *badger.Item) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (tx *badgerTx) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (tx *badgerTx) Commit() error { _ = "STUB: not implemented"; return nil }

func (tx *badgerTx) Rollback() error { _ = "STUB: not implemented"; return nil }

func (tx *badgerTx) Cursor(forward bool) (store.Cursor, error) {
	_ = "STUB: not implemented"
	return *new(store.Cursor), nil
}

type badgerCursor struct {
	it *badger.Iterator
}

func (cursor *badgerCursor) Seek(key []byte) error { _ = "STUB: not implemented"; return nil }

func (cursor *badgerCursor) Next() { _ = "STUB: not implemented"; return }

func (cursor *badgerCursor) Valid() bool { _ = "STUB: not implemented"; return false }

func (cursor *badgerCursor) Item() (store.Item, error) {
	_ = "STUB: not implemented"
	return *new(store.Item), nil
}

func (cursor *badgerCursor) Close() error { _ = "STUB: not implemented"; return nil }

func Open(dir string) (store.Store, error) {
	_ = "STUB: not implemented"
	return *new(store.Store), nil
}

func OpenWithOptions(opts badger.Options) (store.Store, error) {
	_ = "STUB: not implemented"
	return *new(store.Store), nil
}

const (
	GCReclaimInterval = time.Minute * 5
	GCDiscardRatio    = 0.5
)

func (store *badgerStore) startGC() { _ = "STUB: not implemented"; return }

func (store *badgerStore) stopGC() { _ = "STUB: not implemented"; return }
