package array

func intersectSlice[T comparable](a, b []T) []T {
	var intersect []T

	for _, v := range a {
		if Include(b, v) {
			intersect = append(intersect, v)
		}
	}

	return intersect
}

func Intersect[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}

	intersect := slices[0]

	for _, slice := range slices[1:] {
		intersect = intersectSlice(intersect, slice)
	}

	return intersect
}