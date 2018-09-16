package cfg

// BodyList represents a list of production bodies.
type BodyList []Body

// HasEmpty checks if a body list contains a body consisting solely
// of an empty component.
func (l BodyList) HasEmpty() bool {
	for _, body := range l {
		if body.IsEmpty() {
			return true
		}
	}
	return false
}
