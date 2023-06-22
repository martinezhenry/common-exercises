package storage

import "fmt"

type MemoryStorage[K comparable, E any] struct {
	data map[K]E
}

func NewMemoryStorage[K comparable, E any]() *MemoryStorage[K, E] {
	return &MemoryStorage[K, E]{
		data: make(map[K]E),
	}
}

func (s *MemoryStorage[K, E]) Get(key K) (E, error) {	
	if value, exists := s.data[key]; exists {
		return value, nil
	}
	var zeroValue E
	return zeroValue, fmt.Errorf("key not found")
}

func (s *MemoryStorage[K, E]) Set(key K, value E) error {
	s.data[key] = value
	return nil
}

func (s *MemoryStorage[K, E]) Delete(key K) error {
	if _, exists := s.data[key]; exists {
		delete(s.data, key)
		return nil
	}
	return fmt.Errorf("key not found")
}

func (s *MemoryStorage[K, E]) GetAll() ([]E, error) {
	var values []E
	for _, value := range s.data {
		values = append(values, value)
	}
	return values, nil
}