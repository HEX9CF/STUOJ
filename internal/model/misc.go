package model

import "fmt"

type Page struct {
	Page  uint64 `json:"page,omitempty"`
	Size  uint64 `json:"size,omitempty"`
	Total uint64 `json:"total"`
}

type Field[T any] struct {
	exist bool
	value T
}

func (f *Field[T]) Exist() bool {
	return f.exist
}

func (f *Field[T]) Value() T {
	return f.value
}

func (f *Field[T]) Set(value interface{}) error {
	if v, ok := value.(T); ok {
		f.exist = true
		f.value = v
		return nil
	}
	return fmt.Errorf("cannot set value of type %T to field of type %T", value, f.value)
}
