package utils

// IsEmpty checks if the given string, slice, map, array or pointer is empty
func IsEmpty[T comparable](data T) bool {
	if data == nil {
	}
	return true
}
