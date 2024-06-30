package collection

import (
	"encoding/json"
	"fmt"
)

// Set represents an unordered collection with no duplicate values.
type Set[Value comparable] map[Value]struct{}

// Add ensures that the set contains the specified value.
func (collection Set[Value]) Add(value Value) (modified bool) {
	_, modified = collection[value]
	collection[value] = struct{}{}
	return !modified
}

// AddAll ensures that the set contains all of the specified values.
func (collection Set[Value]) AddAll(values ...Value) (modified bool) {
	for _, value := range values {
		_, contains := collection[value]
		collection[value] = struct{}{}
		modified = modified || !contains
	}
	return modified
}

// Clear removes all of the values from the set.
func (collection *Set[Value]) Clear() (modified bool) {
	modified = len(*collection) > 0
	*collection = make(map[Value]struct{})
	return modified
}

// Contains returns true if the set contains the specified value.
func (collection Set[Value]) Contains(value Value) (contains bool) {
	_, contains = collection[value]
	return contains
}

// ContainsAll returns true if the set contains all of the specified values.
func (collection Set[Value]) ContainsAll(values ...Value) (contains bool) {
	for _, value := range values {
		if _, contains = collection[value]; !contains {
			return false
		}
	}
	return true
}

// Equal compares the set to the specified values for equality.
func (collection Set[Value]) Equal(values ...Value) (equal bool) {
	if len(collection) != len(values) {
		return false
	}
	buffer := make(map[Value]struct{}, len(values))
	for _, value := range values {
		if _, contains := collection[value]; !contains {
			return false
		}
		buffer[value] = struct{}{}
	}
	for value := range collection {
		if _, contains := buffer[value]; !contains {
			return false
		}
	}
	return true
}

// ForEach performs the specified action for each value of the set until all
// values have been processed or the action returns false.
func (collection Set[Value]) ForEach(action func(value Value) (next bool)) {
	for value := range collection {
		if !action(value) {
			return
		}
	}
}

// IsEmpty returns true if the set contains no values.
func (collection Set[Value]) IsEmpty() (empty bool) {
	return len(collection) == 0
}

// MarshalJSON returns a byte representation of the set.
func (collection Set[Value]) MarshalJSON() (values []byte, err error) {
	return json.Marshal(collection.Slice())
}

// Partitions performs the specified action for each partition of the specified
// size over the values of the set.
func (collection Set[Value]) Partitions(size int, action func(partition []Value) (next bool)) {
	if len(collection) > 0 && size > 0 {
		partition := make([]Value, size)
		index := 0
		for element := range collection {
			partition[index] = element
			index = (index + 1) % size
			if index == 0 && !action(partition) {
				return
			}
		}
		if index > 0 {
			action(partition[:index])
		}
	}
}

// Remove removes the specified value from the set.
func (collection Set[Value]) Remove(value Value) (modified bool) {
	_, modified = collection[value]
	delete(collection, value)
	return modified
}

// RemoveAll removes all of the specified values from the set.
func (collection Set[Value]) RemoveAll(values ...Value) (modified bool) {
	for _, value := range values {
		_, contains := collection[value]
		delete(collection, value)
		modified = contains || modified
	}
	return modified
}

// RetainAll removes all values in the set that are not included in the
// specified values.
func (collection Set[Value]) RetainAll(values ...Value) (modified bool) {
	buffer := make(Set[Value], len(values))
	for _, value := range values {
		buffer[value] = struct{}{}
	}
	for value := range collection {
		_, contains := buffer[value]
		if !contains {
			delete(collection, value)
			modified = true
		}
	}
	return modified
}

// Size returns the number of values in the set.
func (collection Set[Value]) Size() (size int) {
	return len(collection)
}

// Slice returns a slice containing all of the values in the set.
func (collection Set[Value]) Slice() (values []Value) {
	values = make([]Value, 0, len(collection))
	for value := range collection {
		values = append(values, value)
	}
	return values
}

// String returns a string representation of the set.
func (collection Set[Value]) String() (values string) {
	return fmt.Sprint(collection.Slice())
}

// UnmarshalJSON replaces all of the set's values with the specified values.
func (collection *Set[Value]) UnmarshalJSON(values []byte) (err error) {
	buffer := make([]Value, 0)
	err = json.Unmarshal(values, &buffer)
	collection.Clear()
	for _, value := range buffer {
		(*collection)[value] = struct{}{}
	}
	return err
}
