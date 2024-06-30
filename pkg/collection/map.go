package collection

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Map represents an unordered collection that maps keys to values.
type Map[Key comparable, Value any] map[Key]Value

// Clear removes all of the elements from the map.
func (collection *Map[Key, Value]) Clear() (modified bool) {
	modified = len(*collection) > 0
	*collection = make(map[Key]Value)
	return modified
}

// ContainsAll returns true if the map contains all of the specified elements.
// This method uses reflection to test equality.
func (collection Map[Key, Value]) ContainsAll(elements map[Key]Value) (contains bool) {
	for key, value := range elements {
		if _, exists := collection[key]; !exists {
			return false
		} else if !reflect.DeepEqual(collection[key], value) {
			return false
		}
	}
	return true
}

// ContainsKey returns true if the map contains the specified key.
func (collection Map[Key, Value]) ContainsKey(key Key) (contains bool) {
	_, contains = collection[key]
	return contains
}

// ContainsValue returns true if the map contains the specified value. This
// method uses reflection to test equality.
func (collection Map[Key, Value]) ContainsValue(value Value) (contains bool) {
	for key := range collection {
		if reflect.DeepEqual(collection[key], value) {
			return true
		}
	}
	return false
}

// Equal compares the map to the specified elements for equality. This method
// uses reflection to test equality.
func (collection Map[Key, Value]) Equal(elements map[Key]Value) (equal bool) {
	return reflect.DeepEqual(map[Key]Value(collection), elements)
}

// ForEach performs the specified action for each element of the map until all
// elements have been processed or the action returns false.
func (collection Map[Key, Value]) ForEach(action func(key Key, value Value) (next bool)) {
	for key, value := range collection {
		if !action(key, value) {
			return
		}
	}
}

// Get returns the value associated with the specified key, or the zero value
// if the map does not contain the specified key.
func (collection Map[Key, Value]) Get(key Key) (current Value) {
	return collection[key]
}

// GetOrDefault returns the value associated with the specified key, or the
// specified value if the map does not contain the specified key.
func (collection Map[Key, Value]) GetOrDefault(key Key, value Value) (current Value) {
	current = value
	if value, contains := collection[key]; contains {
		current = value
	}
	return current
}

// IsEmpty returns true if the map contains no elements.
func (collection Map[Key, Value]) IsEmpty() (empty bool) {
	return len(collection) == 0
}

// Keys returns the keys contained in the map.
func (collection Map[Key, Value]) Keys() (keys []Key) {
	keys = make([]Key, 0, len(collection))
	for key := range collection {
		keys = append(keys, key)
	}
	return keys
}

// Map returns a map containing all of the elements in the map.
func (collection Map[Key, Value]) Map() (elements map[Key]Value) {
	elements = make(map[Key]Value, len(collection))
	for key, value := range collection {
		elements[key] = value
	}
	return elements
}

// MarshalJSON returns a byte representation of the map.
func (collection Map[Key, Value]) MarshalJSON() (elements []byte, err error) {
	return json.Marshal(map[Key]Value(collection))
}

// Put associates the specified value with the specified key in the map.
func (collection Map[Key, Value]) Put(key Key, value Value) {
	collection[key] = value
}

// PutAll associates all of the specified values with the specified keys in the
// map.
func (collection Map[Key, Value]) PutAll(elements map[Key]Value) {
	for key, value := range elements {
		collection[key] = value
	}
}

// Remove removes the specified key from the map, returning the previous value.
func (collection Map[Key, Value]) Remove(key Key) (previous Value) {
	previous = collection[key]
	delete(collection, key)
	return previous
}

// Size returns the number of elements in the map.
func (collection Map[Key, Value]) Size() (size int) {
	return len(collection)
}

// String returns a string representation of the map.
func (collection Map[Key, Value]) String() (elements string) {
	return fmt.Sprint(map[Key]Value(collection))
}

// Swap associates the specified value with the specified key in the map,
// returning the previous value.
func (collection Map[Key, Value]) Swap(key Key, value Value) (previous Value) {
	previous = collection[key]
	collection[key] = value
	return previous
}

// UnmarshalJSON replaces all of the map's elements with the specified elements.
func (collection *Map[Key, Value]) UnmarshalJSON(elements []byte) (err error) {
	collection.Clear()
	err = json.Unmarshal(elements, (*map[Key]Value)(collection))
	return err
}

// Values returns the values contained in this map.
func (collection Map[Key, Value]) Values() (values []Value) {
	values = make([]Value, 0, len(collection))
	for _, value := range collection {
		values = append(values, value)
	}
	return values
}
