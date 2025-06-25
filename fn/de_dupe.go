package fn

// DeduplicateList - based on a collection type T and a function that returns the unique KEY.
// The last element that is duplicated will be stored
func DeduplicateList[T any](elements []*T, pkFun func(element *T) string) []*T {
	if len(elements) == 0 {
		return []*T{}
	}

	// Use a map to de-dupe
	dMap := make(map[string]*T, len(elements))
	for _, row := range elements {
		mapPk := pkFun(row)
		dMap[mapPk] = row
	}

	// Pre-allocate the result slice with the exact size needed
	filteredValues := make([]*T, 0, len(dMap))
	for _, row := range dMap {
		filteredValues = append(filteredValues, row)
	}
	return filteredValues
}

// DeduplicateOrderedList - based on a collection type T and a function that returns the unique KEY.
// The first element that is duplicated will be maintained, preserving the order
func DeduplicateOrderedList[T any](elements []*T, pkFun func(element *T) string) []*T {
	if len(elements) == 0 {
		return []*T{}
	}

	// Pre-allocate with maximum possible size
	filteredValues := make([]*T, 0, len(elements))
	seen := make(map[string]struct{}, len(elements))

	// Single pass deduplication while preserving order
	for _, row := range elements {
		mapPk := pkFun(row)
		if _, ok := seen[mapPk]; !ok {
			// Only add the first occurrence of each key
			seen[mapPk] = struct{}{}
			filteredValues = append(filteredValues, row)
		}
	}

	return filteredValues
}
