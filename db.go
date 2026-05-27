package clover

import (
	"errors"

	d "github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/index"
	"github.com/ostafen/clover/v2/query"
	"github.com/ostafen/clover/v2/store"
)

// Collection creation errors
var (
	ErrCollectionExist    = errors.New("collection already exist")
	ErrCollectionNotExist = errors.New("no such collection")

	ErrIndexExist    = errors.New("index already exist")
	ErrIndexNotExist = errors.New("no such index")

	ErrDocumentNotExist = errors.New("no such document")
	ErrDuplicateKey     = errors.New("duplicate key")
)

type docConsumer func(doc *d.Document) error

// DB represents the entry point of each clover database.
type DB struct {
	store  store.Store
	closed uint32
}

type collectionMetadata struct {
	Size    int
	Indexes []index.Info
}

// CreateCollection creates a new empty collection with the given name.
func (db *DB) CreateCollection(name string) error { _ = "STUB: not implemented"; return nil }

func (db *DB) CreateCollectionByQuery(name string, q *query.Query) error {
	_ = "STUB: not implemented"
	return nil
}

// just an empty collection

func (db *DB) saveCollectionMetadata(collection string, meta *collectionMetadata, tx store.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) hasCollection(name string, tx store.Tx) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func getCollectionKey(name string) string { _ = "STUB: not implemented"; return "" }

func getCollectionKeyPrefix() string {
	_ = "STUB: not implemented"

	// DropCollection removes the collection with the given name, deleting any content on disk.
	return ""
}

func (db *DB) DropCollection(name string) error { _ = "STUB: not implemented"; return nil }

func (db *DB) deleteAll(tx store.Tx, collName string) error { _ = "STUB: not implemented"; return nil }

// HasCollection returns true if and only if the database contains a collection with the given name.
func (db *DB) HasCollection(name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func NewObjectId() string { _ = "STUB: not implemented"; return "" }

// Insert adds the supplied documents to a collection.
func (db *DB) Insert(collectionName string, docs ...*d.Document) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) getIndexes(tx store.Tx, collection string, meta *collectionMetadata) []index.Index {
	_ = "STUB: not implemented"
	return nil
}

func saveDocument(doc *d.Document, key []byte, tx store.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) addDocToIndexes(tx store.Tx, indexes []index.Index, doc *d.Document) error {
	_ = "STUB: not implemented"
	// update indexes
	return nil
}

// missing fields are treated as null

func getDocumentKey(collection string, id string) string { _ = "STUB: not implemented"; return "" }

func getDocumentKeyPrefix(collection string) string { _ = "STUB: not implemented"; return "" }

func (db *DB) getCollectionMeta(collection string, tx store.Tx) (*collectionMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Save or update a document, If you pass in a custom struct instead of a Document object,
// it is recommended to specify the _id field using struct tags.
func (db *DB) Save(collectionName string, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// InsertOne inserts a single document to an existing collection. It returns the id of the inserted document.
func (db *DB) InsertOne(collectionName string, doc *d.Document) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Open opens a new clover database on the supplied path. If such a folder doesn't exist, it is automatically created.
func Open(dir string) (*DB, error) { _ = "STUB: not implemented"; return nil, nil }

// OpenWithStore opens a new clover database using the provided store.
func OpenWithStore(store store.Store) (*DB, error) { _ = "STUB: not implemented"; return nil, nil }

// Close releases all the resources and closes the database. After the call, the instance will no more be usable.
func (db *DB) Close() error { _ = "STUB: not implemented"; return nil }

// FindAll selects all the documents satisfying q.
func (db *DB) FindAll(q *query.Query) ([]*d.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *DB) IterateDocs(q *query.Query, consumer docConsumer) error {
	_ = "STUB: not implemented"
	return nil
}

// FindFirst returns the first document (if any) satisfying the query.
func (db *DB) FindFirst(q *query.Query) (*d.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ForEach runs the consumer function for each document matching the provided query.
// If false is returned from the consumer function, then the iteration is stopped.
func (db *DB) ForEach(q *query.Query, consumer func(_ *d.Document) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Count returns the number of documents which satisfy the query (i.e. len(q.FindAll()) == q.Count()).
func (db *DB) Count(q *query.Query) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// simply return the size of the collection in this case

func (db *DB) countCollection(q *query.Query) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (db *DB) getCollectionSize(collection string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Exists returns true if and only if the query result set is not empty.
func (db *DB) Exists(q *query.Query) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// FindById returns the document with the given id, if such a document exists and satisfies the underlying query, or null.
func (db *DB) FindById(collection string, id string) (*d.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDocumentById(collectionName string, id string, tx store.Tx) (*d.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteById removes the document with the given id from the underlying collection, provided that such a document exists and satisfies the underlying query.
func (db *DB) DeleteById(collection string, id string) error { _ = "STUB: not implemented"; return nil }

func (db *DB) getDocAndDeleteFromIndexes(tx store.Tx, indexes []index.Index, collection string, docId string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateById updates the document with the specified id using the supplied update map.
// If no document with the specified id exists, an ErrDocumentNotExist is returned.
func (db *DB) UpdateById(collectionName string, docId string, updater func(doc *d.Document) *d.Document) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) updateIndexesOnDocUpdate(tx store.Tx, indexes []index.Index, oldDoc, newDoc *d.Document) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) deleteDocFromIndexes(indexes []index.Index, doc *d.Document) error {
	_ = "STUB: not implemented"
	return nil
}

// ReplaceById replaces the document with the specified id with the one provided.
// If no document exists, an ErrDocumentNotExist is returned.
func (db *DB) ReplaceById(collection, docId string, doc *d.Document) error {
	_ = "STUB: not implemented"
	return nil
}

// Update updates all the document selected by q using the provided updateMap.
// Each update is specified by a mapping fieldName -> newValue.
func (db *DB) Update(q *query.Query, updateMap map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateFunc updates all the document selected by q using the provided function.
func (db *DB) UpdateFunc(q *query.Query, updateFunc func(doc *d.Document) *d.Document) error {
	_ = "STUB: not implemented"
	return nil
}

type docUpdater func(doc *d.Document) *d.Document

func (db *DB) replaceDocs(tx store.Tx, q *query.Query, updater docUpdater) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) iterateDocs(tx store.Tx, q *query.Query, consumer docConsumer) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete removes all the documents selected by q from the underlying collection.
func (db *DB) Delete(q *query.Query) error { _ = "STUB: not implemented"; return nil }

// ListCollections returns a slice of strings containing the name of each collection stored in the db.
func (db *DB) ListCollections() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func iteratePrefix(prefix []byte, tx store.Tx, itemConsumer func(item store.Item) error) error {
	_ = "STUB: not implemented"
	return nil
}

// do not propagate iteration stop error

// CreateIndex creates an index for the specified for the specified (index, collection) pair.
func (db *DB) CreateIndex(collection, field string) error { _ = "STUB: not implemented"; return nil }

func (db *DB) createIndex(collection, field string, indexType index.Type) error {
	_ = "STUB: not implemented"
	return nil
}

// HasIndex returns true if an index exists for the specified (index, collection) pair.
func (db *DB) HasIndex(collection, field string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (db *DB) hasIndex(tx store.Tx, collection, field string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// DropIndex deletes the index, is such index exists for the specified (index, collection) pair.
func (db *DB) DropIndex(collection, field string) error { _ = "STUB: not implemented"; return nil }

// ListIndexes returns a list containing the names of all the indexes for the specified collection.
func (db *DB) ListIndexes(collection string) ([]index.Info, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *DB) listIndexes(collection string, tx store.Tx) ([]index.Info, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func normalizeCriteria(q *query.Query) (*query.Query, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
