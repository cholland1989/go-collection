package collection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleMap() {
	// Map can be initialized with make
	values := make(Map[int, int])
	values.PutAll(map[int]int{0: 0, 1: 1, 2: 2, 3: 3})
	// Or cast from a compatible map
	values = Map[int, int](map[int]int{0: 0, 1: 1, 2: 2})
	values.Remove(2)
	// And iterated with range
	result := make([]string, 0)
	for key, value := range values {
		result = append(result, fmt.Sprintf("%d=%d", key, value))
	}
	// Iteration order is not guaranteed
	slices.Sort(result)
	fmt.Println(result)
	// Output: [0=0 1=1]
}

func TestMap_Clear(test *testing.T) {
	test.Parallel()

	collection := make(Map[int, int])
	collection.Put(0, 0)
	require.False(test, collection.IsEmpty())
	require.True(test, collection.Clear())
	require.True(test, collection.IsEmpty())
	require.False(test, collection.Clear())
}

func TestMap_ContainsAll(test *testing.T) {
	test.Parallel()

	collection := make(Map[int, int])
	require.False(test, collection.ContainsAll(map[int]int{0: 0}))
	collection.Put(0, 1)
	require.False(test, collection.ContainsAll(map[int]int{0: 0}))
	collection.Put(0, 0)
	require.True(test, collection.ContainsAll(map[int]int{0: 0}))
}

func TestMap_ContainsKey(test *testing.T) {
	test.Parallel()

	collection := make(Map[int, int])
	require.False(test, collection.ContainsKey(0))
	collection.Put(0, 0)
	require.True(test, collection.ContainsKey(0))
}

func TestMap_ContainsValue(test *testing.T) {
	test.Parallel()

	collection := make(Map[int, int])
	require.False(test, collection.ContainsValue(0))
	collection.Put(0, 0)
	require.True(test, collection.ContainsValue(0))
}

func TestMap_Equal(test *testing.T) {
	test.Parallel()

	collection := make(Map[int, int])
	collection.PutAll(map[int]int{0: 0, 1: 1})
	require.False(test, collection.Equal(map[int]int{0: 0}))
	require.False(test, collection.Equal(map[int]int{0: 0, 2: 2}))
	require.True(test, collection.Equal(map[int]int{0: 0, 1: 1}))
}

func TestMap_ForEach(test *testing.T) {
	test.Parallel()

	collection := make(Map[int, int])
	collection.Put(0, 0)
	collection.ForEach(func(key int, value int) bool {
		require.Equal(test, 0, key)
		require.Equal(test, 0, value)
		return false
	})
}

func TestMap_Get(test *testing.T) {
	test.Parallel()

	collection := make(Map[int, int])
	require.Equal(test, 0, collection.Get(0))
	collection.Put(1, 1)
	require.Equal(test, 1, collection.Get(1))
}

func TestMap_GetOrDefault(test *testing.T) {
	test.Parallel()

	collection := make(Map[int, int])
	require.Equal(test, 1, collection.GetOrDefault(0, 1))
	collection.Put(0, 0)
	require.Equal(test, 0, collection.GetOrDefault(0, 1))
}

func TestMap_IsEmpty(test *testing.T) {
	test.Parallel()

	collection := make(Map[int, int])
	require.True(test, collection.IsEmpty())
	collection.Put(0, 0)
	require.False(test, collection.IsEmpty())
}

func TestMap_Keys(test *testing.T) {
	test.Parallel()

	collection := make(Map[int, int])
	require.Len(test, collection.Keys(), 0)
	collection.Put(0, 0)
	require.Len(test, collection.Keys(), 1)
}

func TestMap_Map(test *testing.T) {
	test.Parallel()

	collection := make(Map[int, int])
	require.Len(test, collection.Map(), 0)
	collection.Put(0, 0)
	require.Len(test, collection.Map(), 1)
}

func TestMap_MarshalJSON(test *testing.T) {
	test.Parallel()
	data, err := json.Marshal(map[int]int{0: 0})
	if err != nil {
		test.Fatal(err)
	}
	collection := make(Map[int, int])
	collection.Put(0, 0)
	elements, err := json.Marshal(collection)
	if err != nil {
		test.Fatal(err)
	}
	if !bytes.Equal(elements, data) {
		test.Fatal("method should return elements in map as bytes")
	}
}

func TestMap_Put(test *testing.T) {
	test.Parallel()
	collection := make(Map[int, int])
	collection.Put(0, 0)
	if !collection.Equal(map[int]int{0: 0}) {
		test.Fatal("method should add element to map")
	}
}

func TestMap_PutAll(test *testing.T) {
	test.Parallel()
	collection := make(Map[int, int])
	collection.PutAll(map[int]int{0: 0, 1: 1})
	if !collection.Equal(map[int]int{0: 0, 1: 1}) {
		test.Fatal("method should add elements to map")
	}
}

func TestMap_Remove(test *testing.T) {
	test.Parallel()
	collection := make(Map[int, int])
	collection.Put(0, 1)
	if collection.Remove(0) != 1 {
		test.Fatal("method should return previous value when key exists")
	}
	if collection.Remove(0) != 0 {
		test.Fatal("method should return zero value when key does not exist")
	}
}

func TestMap_Size(test *testing.T) {
	test.Parallel()
	collection := make(Map[int, int])
	collection.Put(0, 0)
	if collection.Size() != 1 {
		test.Fatal("method should return number of elements in map")
	}
}

func TestMap_String(test *testing.T) {
	test.Parallel()
	collection := make(Map[int, int])
	collection.Put(0, 0)
	if fmt.Sprint(collection) != fmt.Sprint(map[int]int{0: 0}) {
		test.Fatal("method should return elements in map as string")
	}
}

func TestMap_Swap(test *testing.T) {
	test.Parallel()
	collection := make(Map[int, int])
	if collection.Swap(0, 1) != 0 {
		test.Fatal("method should return zero value when key does not exist")
	}
	if !collection.Equal(map[int]int{0: 1}) {
		test.Fatal("method should add element to map")
	}
	if collection.Swap(0, 0) != 1 {
		test.Fatal("method should return previous value when key exists")
	}
}

func TestMap_UnmarshalJSON(test *testing.T) {
	test.Parallel()
	data, err := json.Marshal(map[int]int{0: 0})
	if err != nil {
		test.Fatal(err)
	}
	collection := make(Map[int, int])
	collection.Put(1, 1)
	if err := json.Unmarshal(data, &collection); err != nil {
		test.Fatal(err)
	}
	if !collection.Equal(map[int]int{0: 0}) {
		test.Fatal("method should replace elements in map")
	}
}

func TestMap_Values(test *testing.T) {
	test.Parallel()
	collection := make(Map[int, int])
	if len(collection.Values()) != 0 {
		test.Fatal("method should return an empty slice when map is empty")
	}
	collection.Put(0, 0)
	if len(collection.Values()) != 1 {
		test.Fatal("method should return a slice of values when map is not empty")
	}
}
