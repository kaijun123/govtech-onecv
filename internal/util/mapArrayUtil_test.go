package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapToArray(t *testing.T) {
	myMap := map[string]bool{"test": true}
	myArray := MapToArray(myMap)
	assert.Equal(t, myArray, []string{"test"})
}

func TestArrayToMap(t *testing.T) {
	myArray := []string{"test"}
	myMap := ArrayToMap(myArray)
	assert.Equal(t, myMap, map[string]bool{"test": true})
}

func TestMergeMap(t *testing.T) {
	map1 := map[string]bool{"test1": true}
	map2 := map[string]bool{"test2": true}
	combinedMap := MergeMap(map1, map2)
	assert.Equal(t, combinedMap, map[string]bool{"test1": true, "test2": true})
}

func TestMergeArrayWithoutDuplicate(t *testing.T) {
	array1 := []string{"test1"}
	array2 := []string{"test2"}

	combinedArray := MergeArrayWithoutDuplicate(array1, array2)
	assert.Equal(t, combinedArray, []string{"test1", "test2"})
}
