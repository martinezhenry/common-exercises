package internal

// Storage interface defines methods for getting, setting, and deleting
type Storage[K comparable, E any] interface {
	// Get retrieves a value by key
	Get(key K) (E, error)
	// Set stores a value with a key
	Set(key K, value E) error
	// Delete removes a value by key
	Delete(key K) error
	// GetAll retrieves all values
	GetAll() ([]E, error)
}
