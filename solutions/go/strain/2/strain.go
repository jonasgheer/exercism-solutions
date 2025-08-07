package strain

func Keep[E any](coll []E, predicate func(e E) bool) []E {
	result := []E{}
	for _, e := range coll {
		if predicate(e) {
			result = append(result, e)
		}
	}
	return result
}

func Discard[E any](coll []E, predicate func(e E) bool) []E {
	return Keep(coll, func(e E) bool { return !predicate(e) })
}
