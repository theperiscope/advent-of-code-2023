package lib

// Returns a new slice that has the item at index removed
func Remove[T any](items []T, index int) []T {
	result := make([]T, len(items)-1)
	copy(result, items[:index])
	copy(result[index:], items[index+1:])
	return result
}

// Returns copy of the items.
func Clone[T any](items []T) []T {
	result := make([]T, len(items))
	copy(result, items)
	return result
}

// Returns new slice which has items in reverse order
func Reverse[T any](items []T) []T {
	result := make([]T, len(items))
	for i, item := range items {
		result[len(result)-i-1] = item
	}
	return result

}

// Filter returns entries in items where filter function returns true
func Filter[T any](items []T, filter func(T) bool) []T {
	var result []T
	for _, item := range items {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}

// Returns keys in a map; order is not guaranteed
func Keys[T comparable, V any](items map[T]V) []T {
	keys := make([]T, len(items))

	i := 0
	for key := range items {
		keys[i] = key
		i += 1
	}

	return keys
}

// Returns values in a map; order is not guaranteed
func Values[T comparable, V any](items map[T]V) []V {
	values := make([]V, len(items))

	i := 0
	for _, value := range items {
		values[i] = value
		i += 1
	}

	return values
}

// Converts a slice of items of type T to type V using a mapping function. Stops on first error and returns nil.
func Convert[T, V any](items []T, mapper func(T) (V, error)) ([]V, error) {
	result := make([]V, len(items))

	for i, t := range items {
		v, err := mapper(t)
		if err != nil {
			return nil, err
		}
		result[i] = v
	}

	return result, nil
}
