package util

func MergeMap(map1 map[string]bool, map2 map[string]bool) map[string]bool {
	combinedMap := make(map[string]bool)
	for key := range map1 {
		combinedMap[key] = true
	}
	for key := range map2 {
		combinedMap[key] = true
	}
	return combinedMap
}

func MapToArray(m map[string]bool) []string {
	array := []string{}
	for key := range m {
		array = append(array, key)
	}
	return array
}

func ArrayToMap(array []string) map[string]bool {
	myMap := make(map[string]bool)
	for _, str := range array {
		if _, exists := myMap[str]; !exists {
			myMap[str] = true
		}
	}
	return myMap
}

func MergeArrayWithoutDuplicate(array1 []string, array2 []string) []string {
	map1 := ArrayToMap(array1)
	map2 := ArrayToMap(array2)
	combinedMap := MergeMap(map1, map2)

	return MapToArray(combinedMap)
}

func AppendArrayWithoutDuplicate(array1 []string, element string) []string {
	myMap := ArrayToMap(array1)
	myMap[element] = true
	return MapToArray(myMap)
}
