package cfg

// BodyList represents a list of production bodies.
type BodyList []Body

// IsEmpty checks is a body list is empty.
func (l BodyList) IsEmpty() bool {
	return len(l) == 0
}

// HasEmpty checks if a body list contains a body consisting solely
// of an empty component.
func (l BodyList) HasEmpty() bool {
	for _, body := range l {
		if body.IsEmptyBody() {
			return true
		}
	}
	return false
}
