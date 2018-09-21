package cfg

// SetBodyComp implements a set of body components.
type SetBodyComp map[BodyComp]bool

// NewSetBodyComp creates a new set of body components with
// optional initial elements.
func NewSetBodyComp(values ...BodyComp) SetBodyComp {
	newSet := make(map[BodyComp]bool)
	for _, value := range values {
		newSet[value] = true
	}
	return newSet
}

// IsEmpty returns true if a set is the empty set.
func (s SetBodyComp) IsEmpty() bool {
	return len(s) == 0
}

// Length returns the number of elements in the set.
func (s SetBodyComp) Length() int {
	return len(s)
}

// Elements returns an array of the elements in the set.
func (s SetBodyComp) Elements() []BodyComp {
	list := make([]BodyComp, 0, len(s))
	for key := range s {
		list = append(list, key)
	}
	return list
}

// Equals tests if two sets contain the same members
func (s SetBodyComp) Equals(other SetBodyComp) bool {
	if len(s) != len(other) || len(s) != len(s.Union(other)) {
		return false
	}
	return true
}

// Contains returns true if the set contains the specified body component.
func (s SetBodyComp) Contains(n BodyComp) bool {
	return s[n]
}

// ContainsEmpty returns true if the set contains an empty body component.
func (s SetBodyComp) ContainsEmpty() bool {
	return s[BodyComp{BodyEmpty, 0}]
}

// Insert inserts an body component into a set if it isn't already
// in the set.
func (s *SetBodyComp) Insert(n BodyComp) {
	(*s)[n] = true
}

// InsertEmpty inserts an empty body component into a set if it
// isn't already in the set.
func (s *SetBodyComp) InsertEmpty() {
	(*s)[BodyComp{BodyEmpty, 0}] = true
}

// Merge inserts into a set the elements from another set.
func (s *SetBodyComp) Merge(other SetBodyComp) {
	for key, value := range other {
		if value {
			(*s)[key] = true
		}
	}
}

// Copy returns a copy of the set.
func (s SetBodyComp) Copy() SetBodyComp {
	c := NewSetBodyComp()
	for key := range s {
		c[key] = true
	}
	return c
}

// Delete deletes an body component from a set.
func (s *SetBodyComp) Delete(n BodyComp) {
	delete(*s, n)
}

// DeleteEmpty deletes an empty body component from a set.
func (s *SetBodyComp) DeleteEmpty() {
	delete(*s, BodyComp{BodyEmpty, 0})
}

// Intersection returns the intersection of two sets.
func (s SetBodyComp) Intersection(other SetBodyComp) SetBodyComp {
	inter := NewSetBodyComp()
	for key := range s {
		if other[key] {
			inter[key] = true
		}
	}
	return inter
}

// Union returns the union of two sets.
func (s SetBodyComp) Union(other SetBodyComp) SetBodyComp {
	return NewSetBodyComp(append(s.Elements(), other.Elements()...)...)
}
