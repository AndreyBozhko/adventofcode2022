package utils

// Counter is a map for counting comparable keys.
type Counter[T comparable] struct {
	data map[T]int
}

// NewCounter creates a counter with a given key type.
func NewCounter[T comparable]() Counter[T] {
	return Counter[T]{data: make(map[T]int)}
}

// Add increases the count for the element that is inserted.
func (c *Counter[T]) Add(el T) {
	v := c.data[el]
	c.data[el] = v + 1
}

// Get returns the count for a given element.
func (c *Counter[T]) Get(el T) (cnt int, ok bool) {
	cnt, ok = c.data[el]
	return
}

// Len returns the number of keys added to the counter.
func (c *Counter[T]) Len() int {
	return len(c.data)
}

// Remove decreases the count for the element.
// If the count becomes zero, the element is removed.
func (c *Counter[T]) Remove(el T) {
	if v, ok := c.data[el]; ok {
		if v == 1 {
			delete(c.data, el)
		} else {
			c.data[el] = v - 1
		}
	}
}
