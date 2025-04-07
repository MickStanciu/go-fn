package fn

// DeduplicateList - based on a collection type T and a function that returns the unique KEY.
// Last element that is duplicated will be stored
func DeduplicateList[T any](elements []*T, pkFun func(element *T) string) []*T {
	var dMap = map[string]*T{}
	var filteredValues []*T

	if len(elements) == 0 {
		return filteredValues
	}

	// use a map to de-dupe
	for _, row := range elements {
		mapPk := pkFun(row)
		dMap[mapPk] = row
	}

	// convert back to list
	for _, row := range dMap {
		filteredValues = append(filteredValues, row)
	}
	return filteredValues
}
