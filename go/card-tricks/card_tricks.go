package cards

// IsInBounds takes a supplied slice and index and returns true if the index is in bounds.
// Otherwise it returns false.
func IsInBounds[T any](slice []T, index int) bool {
	return index >= 0 && index < len(slice)
}

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    if IsInBounds(slice, index) {
    	return slice[index]
    } else {
		return -1
    }
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if IsInBounds(slice, index) {
        slice[index] = value
    } else {
    	slice = append(slice, value)
    }

    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    if IsInBounds(slice, index) {
		return append(slice[:index], slice[index+1:]...)
    } else {
    	return slice
    }
}
