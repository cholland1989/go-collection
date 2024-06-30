package collection

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleList() {
	// List can be initialized with make
	values := make(List[int], 0)
	values.AddAll(0, 1, 2, 3)
	// Or cast from a compatible slice
	values = List[int]([]int{0, 1, 2})
	values.Remove(2)
	// And iterated with range
	result := make([]string, 0)
	for index, value := range values {
		result = append(result, fmt.Sprintf("%d=%d", index, value))
	}
	fmt.Println(result)
	// Output: [0=0 1=1]
}

func TestList_Add(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.Add(0))
	require.True(test, collection.Equal(0))
	require.True(test, collection.Add(0))
	require.True(test, collection.Equal(0, 0))
}

func TestList_AddAll(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.AddAll(0, 1))
	require.True(test, collection.Equal(0, 1))
	require.True(test, collection.AddAll(0, 1))
	require.True(test, collection.Equal(0, 1, 0, 1))
}

func TestList_Clear(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.Add(0))
	require.False(test, collection.IsEmpty())
	require.True(test, collection.Clear())
	require.True(test, collection.IsEmpty())
	require.False(test, collection.Clear())
}

func TestList_Contains(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.False(test, collection.Contains(0))
	require.True(test, collection.Add(0))
	require.True(test, collection.Contains(0))
}

func TestList_ContainsAll(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.AddAll(0, 1))
	require.False(test, collection.ContainsAll(0, 2))
	require.True(test, collection.ContainsAll(0, 1))
}

func TestList_Delete(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	previous, err := collection.Delete(0)
	require.Error(test, err)
	require.Equal(test, 0, previous)
	require.True(test, collection.Add(1))

	previous, err = collection.Delete(0)
	require.NoError(test, err)
	require.Equal(test, 1, previous)
}

func TestList_Equal(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.AddAll(0, 1))
	require.False(test, collection.Equal(0))
	require.False(test, collection.Equal(0, 0))
	require.False(test, collection.Equal(0, 2))
	require.True(test, collection.Equal(0, 1))
	require.False(test, collection.Equal(1, 0))
}

func TestList_ForEach(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.Add(0))
	collection.ForEach(func(value int) bool {
		require.Equal(test, 0, value)
		return false
	})
}

func TestList_Get(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	current, err := collection.Get(0)
	require.Error(test, err)
	require.Equal(test, 0, current)
	require.True(test, collection.Add(1))

	current, err = collection.Get(0)
	require.NoError(test, err)
	require.Equal(test, 1, current)
}

func TestList_IndexOf(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.Equal(test, -1, collection.IndexOf(0))
	require.True(test, collection.AddAll(0, 0))
	require.Equal(test, 0, collection.IndexOf(0))
}

func TestList_Insert(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.Error(test, collection.Insert(1, 0))
	require.NoError(test, collection.Insert(0, 0))
	require.NoError(test, collection.Insert(0, 1))
	require.True(test, collection.Equal(1, 0))
}

func TestList_InsertAll(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.Error(test, collection.InsertAll(1, 0, 1))
	require.NoError(test, collection.InsertAll(0, 0, 1))
	require.NoError(test, collection.InsertAll(1, 0, 1))
	require.True(test, collection.Equal(0, 0, 1, 1))
}

func TestList_IsEmpty(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.IsEmpty())
	require.True(test, collection.Add(0))
	require.False(test, collection.IsEmpty())
}

func TestList_LastIndexOf(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.Equal(test, -1, collection.LastIndexOf(0))
	require.True(test, collection.AddAll(0, 0))
	require.Equal(test, 1, collection.LastIndexOf(0))
}

func TestList_MarshalJSON(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.Add(0))

	data, err := json.Marshal(collection)
	require.NoError(test, err)

	expected, err := json.Marshal([]int{0})
	require.NoError(test, err)
	require.Equal(test, expected, data)
}

func TestList_Partitions(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.AddAll(0, 1, 2))

	length := 2
	collection.Partitions(2, func(values []int) bool {
		require.Len(test, values, length)
		length--
		return true
	})

	collection.Partitions(2, func(values []int) bool {
		require.Len(test, values, 2)
		return false
	})
}

func TestList_Remove(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.AddAll(0, 0, 1))
	require.True(test, collection.Remove(0))
	require.True(test, collection.Equal(0, 1))
	require.True(test, collection.Remove(0))
	require.True(test, collection.Equal(1))
	require.False(test, collection.Remove(0))
}

func TestList_RemoveAll(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.AddAll(0, 0, 1))
	require.True(test, collection.RemoveAll(0))
	require.True(test, collection.Equal(1))
	require.False(test, collection.RemoveAll(0))
}

func TestList_RetainAll(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.AddAll(0, 1, 1))
	require.True(test, collection.RetainAll(0))
	require.True(test, collection.Equal(0))
	require.False(test, collection.RetainAll(0))
}

func TestList_Reverse(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.AddAll(0, 1))
	collection.Reverse()
	require.True(test, collection.Equal(1, 0))
}

func TestList_Set(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.Error(test, collection.Set(0, 0))
	require.True(test, collection.Add(0))
	require.NoError(test, collection.Set(0, 0))
}

func TestList_Size(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.Add(0))
	require.Equal(test, 1, collection.Size())
}

func TestList_Slice(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.Add(0))
	require.Len(test, collection.Slice(), 1)
}

func TestList_Sort(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.AddAll(1, 0))
	collection.Sort(func(this int, that int) bool { return this < that })
	require.True(test, collection.Equal(0, 1))
}

func TestList_String(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.Add(0))
	require.Equal(test, fmt.Sprint([]int{0}), fmt.Sprint(collection))
}

func TestList_Swap(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	previous, err := collection.Swap(0, 1)
	require.Error(test, err)
	require.Equal(test, 0, previous)
	require.True(test, collection.Add(1))

	previous, err = collection.Swap(0, 1)
	require.NoError(test, err)
	require.Equal(test, 1, previous)
}

func TestList_UnmarshalJSON(test *testing.T) {
	test.Parallel()

	collection := make(List[int], 0)
	require.True(test, collection.Add(1))

	data, err := json.Marshal([]int{0})
	require.NoError(test, err)

	err = json.Unmarshal(data, &collection)
	require.NoError(test, err)
	require.True(test, collection.Equal(0))
}
