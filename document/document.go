package document

import (
	"time"
)

const (
	ObjectIdField  = "_id"
	ExpiresAtField = "_expiresAt"
)

// Document represents a document as a map.
type Document struct {
	fields map[string]interface{}
}

// ObjectId returns the id of the document, provided that the document belongs to some collection. Otherwise, it returns the empty string.
func (doc *Document) ObjectId() string { _ = "STUB: not implemented"; return "" }

// NewDocument creates a new empty document.
func NewDocument() *Document { _ = "STUB: not implemented"; return nil }

// NewDocumentOf creates a new document and initializes it with the content of the provided object.
// It returns nil if the object cannot be converted to a valid Document.
func NewDocumentOf(o interface{}) *Document { _ = "STUB: not implemented"; return nil }

func newDocumentOf(o interface{}) *Document { _ = "STUB: not implemented"; return nil }

// Copy returns a shallow copy of the underlying document.
func (doc *Document) Copy() *Document { _ = "STUB: not implemented"; return nil }

func (doc *Document) AsMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func lookupField(name string, fieldMap map[string]interface{}, force bool) (map[string]interface{}, interface{}, string) {
	_ = "STUB: not implemented"
	return nil, nil, ""
}

// Has tells returns true if the document contains a field with the supplied name.
func (doc *Document) Has(name string) bool { _ = "STUB: not implemented"; return false }

// Get retrieves the value of a field. Nested fields can be accessed using dot.
func (doc *Document) Get(name string) interface{} { _ = "STUB: not implemented"; return nil }

// Set maps a field to a value. Nested fields can be accessed using dot.
func (doc *Document) Set(name string, value interface{}) { _ = "STUB: not implemented"; return }

// SetAll sets each field specified in the input map to the corresponding value. Nested fields can be accessed using dot.
func (doc *Document) SetAll(values map[string]interface{}) { _ = "STUB: not implemented"; return }

// ToMap returns a map of all available fields in the document. Nested fields are represented by sub-maps. This is a deep copy, but values are not cloned.
func (doc *Document) ToMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

// Fields returns a lexicographically sorted slice of all available field names in the document.
// Nested fields, if included, are represented using dot notation.
func (doc *Document) Fields(includeSubFields bool) []string { _ = "STUB: not implemented"; return nil }

// ExpiresAt returns the document expiration instant
func (doc *Document) ExpiresAt() *time.Time { _ = "STUB: not implemented"; return nil }

// SetExpiresAt sets document expiration
func (doc *Document) SetExpiresAt(expiration time.Time) { _ = "STUB: not implemented"; return }

// TTL returns a duration representing the time to live of the document before expiration.
// A negative duration means that the document has no expiration, while a zero value represents an already expired document.
func (doc *Document) TTL() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// document already expired

// Unmarshal stores the document in the value pointed by v.
func (doc *Document) Unmarshal(v interface{}) error { _ = "STUB: not implemented"; return nil }

func isValidObjectId(id string) bool { _ = "STUB: not implemented"; return false }

func Validate(doc *Document) error { _ = "STUB: not implemented"; return nil }

func Decode(data []byte) (*Document, error) { _ = "STUB: not implemented"; return nil, nil }

func Encode(doc *Document) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
