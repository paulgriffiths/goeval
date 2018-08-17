package symtab

import (
	"testing"
)

func TestSymbolTableStoreAndRetrieve(t *testing.T) {
	tab := NewTable()
	key := "meaning of life"
	value := 42
	tab.Store(key, NewInt(value))
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("couldn't retrieve value for %s", key)
	} else {
		if result.IntValue() != value {
			t.Errorf("got %d, want %d", result.IntValue(), value)
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
	value := 0.5
	tab.Store(key, NewReal(value))
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("couldn't retrieve value for %s", key)
	} else {
		if result.RealValue() != value {
			t.Errorf("got %f, want %f", result.RealValue(), value)
		}
	}

	key = "went to moon"
	newValue := true
	tab.Store(key, NewBool(newValue))
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("couldn't retrieve value for %s", key)
	} else {
		if result.BoolValue() != newValue {
			t.Errorf("got %t, want %t", result.BoolValue(), newValue)
		}
	}
}

func TestSymbolTableShadow(t *testing.T) {
	tab := NewTable()
	key := "antidote"
	value := 99
	tab.Store(key, NewInt(value))
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("failed to retrieve for key %s", key)
		return
	} else if testValue := result.IntValue(); testValue != value {
		t.Errorf("got %d, want %d", testValue, value)
	}

	tab.Push()
	shadowValue := 103.5
	tab.Store(key, NewReal(shadowValue))
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("failed to retrieve for key %s", key)
		return
	} else if testValue := result.RealValue(); testValue != shadowValue {
		t.Errorf("got %f, want %f", testValue, shadowValue)
	}

	tab.Pop()
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("failed to retrieve for key %s", key)
		return
	} else if testValue := result.IntValue(); testValue != value {
		t.Errorf("got %d, want %d", testValue, value)
	}
}

func TestSymbolTableGoesOutOfScope(t *testing.T) {
	tab := NewTable()
	key := "frantic"
	value := 99.5

	tab.Push()
	tab.Store(key, NewReal(value))
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("failed to retrieve for key %s", key)
		return
	} else if testValue := result.RealValue(); testValue != value {
		t.Errorf("got %f, want %f", testValue, value)
	}

	tab.Pop()
	if _, ok := tab.Retrieve(key); ok {
		t.Errorf("failed to indicate absence of key")
	}
}

func TestSymbolTableResolution(t *testing.T) {
	tab := NewTable()
	key := "over the"
	value := "top"

	tab.Store(key, NewString(value))
	tab.Push()
	if result, ok := tab.Retrieve(key); !ok {
		t.Errorf("failed to retrieve for key %s", key)
		return
	} else if testValue := result.StringValue(); testValue != value {
		t.Errorf("got %s, want %s", testValue, value)
	}
}
