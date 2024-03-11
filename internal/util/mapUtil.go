package util

func mergeMap(map1 map[string]bool, map2 map[string]bool) map[string]bool {
	combinedMap := make(map[string]bool)
	for key := range map1 {
		if map2[key] {
			combinedMap[key] = true
		}
	}
	return combinedMap
}
