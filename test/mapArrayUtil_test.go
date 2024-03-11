package test

import (
	"govtech-onecv/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapToArray(t *testing.T) {
	myMap := map[string]bool{"test": true}
	myArray := util.MapToArray(myMap)
	assert.Equal(t, myArray, []string{"test"})
}

func TestArrayToMap(t *testing.T) {
	myArray := []string{"test"}
	myMap := util.ArrayToMap(myArray)
	assert.Equal(t, myMap, map[string]bool{"test": true})
}

func TestMergeMap(t *testing.T) {
	map1 := map[string]bool{"test1": true}
	map2 := map[string]bool{"test2": true}
	combinedMap := util.MergeMap(map1, map2)
	assert.Equal(t, combinedMap, map[string]bool{"test1": true, "test2": true})
}

func TestMergeArrayWithoutDuplicate(t *testing.T) {
	array1 := []string{"test1"}
	array2 := []string{"test2"}

	expected := []string{}
	expected = append(expected, "test1")
	expected = append(expected, "test2")

	combinedArray := util.MergeArrayWithoutDuplicate(array1, array2)
	assert.Equal(t, combinedArray, expected)
}
