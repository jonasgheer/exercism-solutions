package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

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
	result := []E{}
	for _, e := range coll {
		if !predicate(e) {
			result = append(result, e)
		}
	}
	return result
}
