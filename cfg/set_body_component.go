package cfg

// SetBodyComp implements a set of integers.
type SetBodyComp map[BodyComp]bool

// NewSetBodyComp creates a new set of integers with optional initial elements.
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

// Contains returns true if the set contains the specified integer.
func (s SetBodyComp) Contains(n BodyComp) bool {
	return s[n]
}

// Insert inserts an integer into a set if it isn't already in the set.
func (s *SetBodyComp) Insert(n BodyComp) {
	(*s)[n] = true
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
