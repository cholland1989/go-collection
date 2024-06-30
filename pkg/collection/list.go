// Package collection provides generic list, map, and set definitions with
// common utility methods.
package collection

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
)

// ErrIndexOutOfRange indicates that an index was out of range.
var ErrIndexOutOfRange = errors.New("index out of range")

// List represents an ordered collection of values.
type List[Value any] []Value

// Add ensures that the list contains the specified value.
func (collection *List[Value]) Add(value Value) (modified bool) {
	*collection = append(*collection, value)
	return true
}

// AddAll ensures that the list contains all of the specified values.
func (collection *List[Value]) AddAll(values ...Value) (modified bool) {
	*collection = append(*collection, values...)
	return len(values) != 0
}

// Clear removes all of the values from the list.
func (collection *List[Value]) Clear() (modified bool) {
	modified = len(*collection) > 0
	*collection = make([]Value, 0)
	return modified
}

// Contains returns true if the list contains the specified value. This method
// uses reflection to test equality.
func (collection List[Value]) Contains(value Value) (contains bool) {
	for index := range collection {
		if reflect.DeepEqual(collection[index], value) {
			return true
		}
	}
	return false
}

// ContainsAll returns true if the list contains all of the specified values.
// This method uses reflection to test equality.
func (collection List[Value]) ContainsAll(values ...Value) (contains bool) {
OuterLoop:
	for index := range values {
		for jndex := range collection {
			if reflect.DeepEqual(collection[jndex], values[index]) {
				continue OuterLoop
			}
		}
		return false
	}
	return true
}

// Delete removes the value at the specified position in the list, returning
// the previous value.
func (collection *List[Value]) Delete(index int) (previous Value, err error) {
	if index >= 0 && index < len(*collection) {
		var empty Value
		previous = (*collection)[index]
		copy((*collection)[index:], (*collection)[index+1:])
		(*collection)[len(*collection)-1] = empty
		*collection = (*collection)[:len(*collection)-1]
	} else {
		err = ErrIndexOutOfRange
	}
	return previous, err
}

// Equal compares the list to the specified values for equality. This method
// uses reflection to test equality.
func (collection List[Value]) Equal(values ...Value) (equal bool) {
	return reflect.DeepEqual([]Value(collection), values)
}

// ForEach performs the specified action for each value of the list until all
// values have been processed or the action returns false.
func (collection List[Value]) ForEach(action func(value Value) (next bool)) {
	for index := range collection {
		if !action(collection[index]) {
			return
		}
	}
}

// Get returns the value at the specified position in the list.
func (collection List[Value]) Get(index int) (current Value, err error) {
	if index >= 0 && index < len(collection) {
		current = collection[index]
	} else {
		err = ErrIndexOutOfRange
	}
	return current, err
}

// IndexOf returns the index of the first occurrence of the specified value in
// the list, or -1 if the list does not contain the specified value. This
// method uses reflection to test equality.
func (collection List[Value]) IndexOf(value Value) (index int) {
	for index := range collection {
		if reflect.DeepEqual(collection[index], value) {
			return index
		}
	}
	return -1
}

// Insert adds the specified value to the list at the specified position.
func (collection *List[Value]) Insert(index int, value Value) (err error) {
	if index >= 0 && index <= len(*collection) {
		var empty Value
		*collection = append(*collection, empty)
		copy((*collection)[index+1:], (*collection)[index:])
		(*collection)[index] = value
	} else {
		err = ErrIndexOutOfRange
	}
	return err
}

// InsertAll adds all of the specified values to the list at the specified
// position.
func (collection *List[Value]) InsertAll(index int, values ...Value) (err error) {
	if index >= 0 && index <= len(*collection) {
		*collection = append((*collection)[:index], append(values, (*collection)[index:]...)...)
	} else {
		err = ErrIndexOutOfRange
	}
	return err
}

// IsEmpty returns true if the list contains no values.
func (collection List[Value]) IsEmpty() (empty bool) {
	return len(collection) == 0
}

// LastIndexOf returns the index of the last occurrence of the specified value
// in the list, or -1 if the list does not contain the specified value. This
// method uses reflection to test equality.
func (collection List[Value]) LastIndexOf(value Value) (index int) {
	for index = len(collection) - 1; index >= 0; index-- {
		if reflect.DeepEqual(collection[index], value) {
			return index
		}
	}
	return -1
}

// MarshalJSON returns a byte representation of the list.
func (collection List[Value]) MarshalJSON() (values []byte, err error) {
	return json.Marshal([]Value(collection))
}

// Partitions performs the specified action for each partition of the specified
// size over the values of the list.
func (collection List[Value]) Partitions(size int, action func(values []Value) (next bool)) {
	if len(collection) > 0 && size > 0 {
		index := 0
		for index < len(collection)-size {
			if !action(collection[index : index+size]) {
				return
			}
			index += size
		}
		action(collection[index:])
	}
}

// Remove removes a single instance of the specified value from the list. This
// method uses reflection to test equality.
func (collection *List[Value]) Remove(value Value) (modified bool) {
	for index := range *collection {
		if !reflect.DeepEqual((*collection)[index], value) {
			continue
		}
		var empty Value
		copy((*collection)[index:], (*collection)[index+1:])
		(*collection)[len(*collection)-1] = empty
		*collection = (*collection)[:len(*collection)-1]
		return true
	}
	return false
}

// RemoveAll removes all instances of the specified values from the list. This
// method uses reflection to test equality.
func (collection *List[Value]) RemoveAll(values ...Value) (modified bool) {
	index := 0
OuterLoop:
	for jndex := range *collection {
		for kndex := range values {
			if reflect.DeepEqual((*collection)[jndex], values[kndex]) {
				continue OuterLoop
			}
		}
		(*collection)[index] = (*collection)[jndex]
		index++
	}
	modified = index != len(*collection)
	copy((*collection)[index:], make([]Value, len(*collection)-index))
	*collection = (*collection)[:index]
	return modified
}

// RetainAll removes all values in the list that are not included in the
// specified values. This method uses reflection to test equality.
func (collection *List[Value]) RetainAll(values ...Value) (modified bool) {
	index := 0
OuterLoop:
	for jndex := range *collection {
		for kndex := range values {
			if !reflect.DeepEqual((*collection)[jndex], values[kndex]) {
				continue OuterLoop
			}
		}
		(*collection)[index] = (*collection)[jndex]
		index++
	}
	modified = index != len(*collection)
	copy((*collection)[index:], make([]Value, len(*collection)-index))
	*collection = (*collection)[:index]
	return modified
}

// Reverse reverses the order of the values in the list.
func (collection List[Value]) Reverse() {
	for index, jndex := 0, len(collection)-1; index < jndex; index, jndex = index+1, jndex-1 {
		collection[index], collection[jndex] = collection[jndex], collection[index]
	}
}

// Set replaces the value at the specified position in the list with the
// specified value.
func (collection List[Value]) Set(index int, value Value) (err error) {
	if index >= 0 && index < len(collection) {
		collection[index] = value
	} else {
		err = ErrIndexOutOfRange
	}
	return err
}

// Size returns the number of values in the list.
func (collection List[Value]) Size() (size int) {
	return len(collection)
}

// Slice returns a slice containing all of the values in the list.
func (collection List[Value]) Slice() (values []Value) {
	return append(make([]Value, 0, len(collection)), collection...)
}

// Sort reorders the list according to the order induced by the specified
// comparator.
func (collection List[Value]) Sort(comparator func(this Value, that Value) (swap bool)) {
	sort.Slice(collection, func(index, jndex int) bool {
		return comparator(collection[index], collection[jndex])
	})
}

// String returns a string representation of the list.
func (collection List[Value]) String() (values string) {
	return fmt.Sprint([]Value(collection))
}

// Swap replaces the value at the specified position in the list with the
// specified value, returning the previous value.
func (collection List[Value]) Swap(index int, value Value) (previous Value, err error) {
	if index >= 0 && index < len(collection) {
		previous = collection[index]
		collection[index] = value
	} else {
		err = ErrIndexOutOfRange
	}
	return previous, err
}

// UnmarshalJSON replaces all of the list's values with the specified values.
func (collection *List[Value]) UnmarshalJSON(values []byte) (err error) {
	collection.Clear()
	err = json.Unmarshal(values, (*[]Value)(collection))
	return err
}
