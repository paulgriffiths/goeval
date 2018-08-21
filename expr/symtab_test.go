package expr

import (
	"testing"
)

func TestSymbolTableStoreAndRetrieve(t *testing.T) {
	tab := NewTable()
	key := "meaning of life"
	value := intValue{42}
	tab.Store(key, value)
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("couldn't Retrieve value for %s", key)
	} else {
		if !result.Equals(value) {
			t.Errorf("got %d, want %d", result, value)
		}
	}
}

func TestSymbolTableRetrieveNotSet(t *testing.T) {
	tab := NewTable()
	key := "nothing to see"
	if _, ok := tab.Retrieve(key); ok {
		t.Errorf("failed to indicate absence of key")
	}
}

func TestSymbolTableOverWrite(t *testing.T) {
	tab := NewTable()
	key := "halves in a quarter"
	value := realValue{0.5}
	tab.Store(key, value)
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("couldn't Retrieve value for %s", key)
	} else {
		if !result.Equals(value) {
			t.Errorf("got %f, want %f", result, value)
		}
	}

	key = "went to moon"
	newValue := boolValue{true}
	tab.Store(key, newValue)
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("couldn't Retrieve value for %s", key)
	} else {
		if !result.Equals(newValue) {
			t.Errorf("got %t, want %t", result, newValue)
		}
	}
}

func TestSymbolTableShadow(t *testing.T) {
	tab := NewTable()
	key := "antidote"
	value := intValue{99}
	tab.Store(key, value)
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("failed to Retrieve for key %s", key)
		return
	} else if !result.Equals(value) {
		t.Errorf("got %d, want %d", result, value)
	}

	tab.Push()
	shadowValue := realValue{103.5}
	tab.Store(key, shadowValue)
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("failed to Retrieve for key %s", key)
		return
	} else if !result.Equals(shadowValue) {
		t.Errorf("got %f, want %f", result, shadowValue)
	}

	tab.Pop()
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("failed to Retrieve for key %s", key)
		return
	} else if !result.Equals(value) {
		t.Errorf("got %d, want %d", result, value)
	}
}

func TestSymbolTableGoesOutOfScope(t *testing.T) {
	tab := NewTable()
	key := "frantic"
	value := realValue{99.5}

	tab.Push()
	tab.Store(key, value)
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("failed to Retrieve for key %s", key)
		return
	} else if !result.Equals(value) {
		t.Errorf("got %f, want %f", result, value)
	}

	tab.Pop()
	if _, ok := tab.Retrieve(key); ok {
		t.Errorf("failed to indicate absence of key")
	}
}

func TestSymbolTableResolution(t *testing.T) {
	tab := NewTable()
	key := "over the"
	value := stringValue{"top"}

	tab.Store(key, value)
	tab.Push()
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("failed to Retrieve for key %s", key)
		return
	} else if !result.Equals(value) {
		t.Errorf("got %s, want %s", result, value)
	}
}
