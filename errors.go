// Sentinel errors returned by lionet operations.
package lionet

import "errors"

var (
	ErrKeyNotFound       = errors.New("lionet: key not found")
	ErrEmptyKey          = errors.New("lionet: key cannot be empty")
	ErrDBClosed          = errors.New("lionet: database is closed")
	ErrReadOnlyTxn       = errors.New("lionet: read-only transaction")
	ErrManifestCorrupted = errors.New("lionet: manifest corrupted")
	ErrWALCorrupted      = errors.New("lionet: WAL corrupted")
	ErrValueTooLarge     = errors.New("lionet: value too large")
)
