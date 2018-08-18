package vareval

import (
	"testing"
)

func TestSymbolTablestoreAndretrieve(t *testing.T) {
	tab := newTable()
	key := "meaning of life"
	value := intValue{42}
	tab.store(key, value)
	if result, ok := tab.retrieve(key); !ok {
		t.Errorf("couldn't retrieve value for %s", key)
	} else {
		if !result.equals(value) {
			t.Errorf("got %d, want %d", result, value)
		}
	}
}

func TestSymbolTableretrieveNotSet(t *testing.T) {
	tab := newTable()
	key := "nothing to see"
	if _, ok := tab.retrieve(key); ok {
		t.Errorf("failed to indicate absence of key")
	}
}

func TestSymbolTableOverWrite(t *testing.T) {
	tab := newTable()
	key := "halves in a quarter"
	value := realValue{0.5}
	tab.store(key, value)
	if result, ok := tab.retrieve(key); !ok {
		t.Errorf("couldn't retrieve value for %s", key)
	} else {
		if !result.equals(value) {
			t.Errorf("got %f, want %f", result, value)
		}
	}

	key = "went to moon"
	newValue := boolValue{true}
	tab.store(key, newValue)
	if result, ok := tab.retrieve(key); !ok {
		t.Errorf("couldn't retrieve value for %s", key)
	} else {
		if !result.equals(newValue) {
			t.Errorf("got %t, want %t", result, newValue)
		}
	}
}

func TestSymbolTableShadow(t *testing.T) {
	tab := newTable()
	key := "antidote"
	value := intValue{99}
	tab.store(key, value)
	if result, ok := tab.retrieve(key); !ok {
		t.Errorf("failed to retrieve for key %s", key)
		return
	} else if !result.equals(value) {
		t.Errorf("got %d, want %d", result, value)
	}

	tab.push()
	shadowValue := realValue{103.5}
	tab.store(key, shadowValue)
	if result, ok := tab.retrieve(key); !ok {
		t.Errorf("failed to retrieve for key %s", key)
		return
	} else if !result.equals(shadowValue) {
		t.Errorf("got %f, want %f", result, shadowValue)
	}

	tab.pop()
	if result, ok := tab.retrieve(key); !ok {
		t.Errorf("failed to retrieve for key %s", key)
		return
	} else if !result.equals(value) {
		t.Errorf("got %d, want %d", result, value)
	}
}

func TestSymbolTableGoesOutOfScope(t *testing.T) {
	tab := newTable()
	key := "frantic"
	value := realValue{99.5}

	tab.push()
	tab.store(key, value)
	if result, ok := tab.retrieve(key); !ok {
		t.Errorf("failed to retrieve for key %s", key)
		return
	} else if !result.equals(value) {
		t.Errorf("got %f, want %f", result, value)
	}

	tab.pop()
	if _, ok := tab.retrieve(key); ok {
		t.Errorf("failed to indicate absence of key")
	}
}

func TestSymbolTableResolution(t *testing.T) {
	tab := newTable()
	key := "over the"
	value := stringValue{"top"}

	tab.store(key, value)
	tab.push()
	if result, ok := tab.retrieve(key); !ok {
		t.Errorf("failed to retrieve for key %s", key)
		return
	} else if !result.equals(value) {
		t.Errorf("got %s, want %s", result, value)
	}
}
