package collection

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleSet() {
	// Set can be initialized with make
	values := make(Set[int], 0)
	values.AddAll(0, 1, 2, 3)
	// Or cast from a compatible map
	values = Set[int](map[int]struct{}{0: {}, 1: {}, 2: {}})
	values.Remove(2)
	// And iterated with range
	result := make([]int, 0)
	for value := range values {
		result = append(result, value)
	}
	// Iteration order is not guaranteed
	sort.Ints(result)
	fmt.Println(result)
	// Output: [0 1]
}

func TestSet_Add(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.Add(0))
	require.True(test, collection.Equal(0))
	require.False(test, collection.Add(0))
	require.True(test, collection.Equal(0))
}

func TestSet_AddAll(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.AddAll(0, 1))
	require.True(test, collection.Equal(0, 1))
	require.False(test, collection.AddAll(0, 1))
	require.True(test, collection.Equal(0, 1))
}

func TestSet_Clear(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.Add(0))
	require.False(test, collection.IsEmpty())
	require.True(test, collection.Clear())
	require.True(test, collection.IsEmpty())
	require.False(test, collection.Clear())
}

func TestSet_Contains(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.False(test, collection.Contains(0))
	require.True(test, collection.Add(0))
	require.True(test, collection.Contains(0))
}

func TestSet_ContainsAll(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.AddAll(0, 1))
	require.False(test, collection.ContainsAll(0, 2))
	require.True(test, collection.ContainsAll(0, 1))
}

func TestSet_Equal(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.AddAll(0, 1))
	require.False(test, collection.Equal(0))
	require.False(test, collection.Equal(0, 0))
	require.False(test, collection.Equal(0, 2))
	require.True(test, collection.Equal(0, 1))
	require.True(test, collection.Equal(1, 0))
}

func TestSet_ForEach(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.Add(0))
	collection.ForEach(func(value int) (next bool) {
		require.Equal(test, 0, value)
		return false
	})
}

func TestSet_IsEmpty(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.IsEmpty())
	require.True(test, collection.Add(0))
	require.False(test, collection.IsEmpty())
}

func TestSet_MarshalJSON(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.Add(0))

	data, err := json.Marshal(collection)
	require.NoError(test, err)

	expected, err := json.Marshal([]int{0})
	require.NoError(test, err)
	require.Equal(test, expected, data)
}

func TestSet_Partitions(test *testing.T) {
	test.Parallel()

	collection := make(Set[int], 0)
	require.True(test, collection.AddAll(0, 1, 2))

	length := 2
	collection.Partitions(2, func(partition []int) bool {
		require.Len(test, partition, length)
		length--
		return true
	})

	collection.Partitions(2, func(partition []int) bool {
		require.Len(test, partition, 2)
		return false
	})
}

func TestSet_Remove(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.AddAll(0, 1))
	require.True(test, collection.Remove(0))
	require.True(test, collection.Equal(1))
	require.False(test, collection.Remove(0))
}

func TestSet_RemoveAll(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.AddAll(0, 1, 2))
	require.True(test, collection.RemoveAll(0, 1))
	require.True(test, collection.Equal(2))
	require.False(test, collection.RemoveAll(0, 1))
}

func TestSet_RetainAll(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.AddAll(0, 1, 2))
	require.True(test, collection.RetainAll(0, 1))
	require.True(test, collection.Equal(0, 1))
	require.False(test, collection.RetainAll(0, 1))
}

func TestSet_Size(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.Add(0))
	require.Equal(test, 1, collection.Size())
}

func TestSet_Slice(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.Add(0))
	require.Len(test, collection.Slice(), 1)
}

func TestSet_String(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.Add(0))
	require.Equal(test, fmt.Sprint([]int{0}), fmt.Sprint(collection))
}

func TestSet_UnmarshalJSON(test *testing.T) {
	test.Parallel()

	collection := make(Set[int])
	require.True(test, collection.Add(1))

	data, err := json.Marshal([]int{0})
	require.NoError(test, err)

	err = json.Unmarshal(data, &collection)
	require.NoError(test, err)
	require.True(test, collection.Equal(0))
}
