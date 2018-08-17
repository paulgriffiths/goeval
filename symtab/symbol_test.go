package symtab

import (
	"testing"
)

func TestBasicSymbolInt(t *testing.T) {
	testValue := 6
	sym := NewInt(testValue)
	if result := sym.IntValue(); result != testValue {
		t.Errorf("Got %d, want %d", result, testValue)
	}
}

func TestBasicSymbolReal(t *testing.T) {
	testValue := 7.5
	sym := NewReal(testValue)
	if result := sym.RealValue(); result != testValue {
		t.Errorf("Got %f, want %f", result, testValue)
	}
}

func TestBasicSymbolBool(t *testing.T) {
	testValue := true
	sym := NewBool(testValue)
	if result := sym.BoolValue(); result != testValue {
		t.Errorf("Got %t, want %t", result, testValue)
	}
}

func TestBasicSymbolString(t *testing.T) {
	testValue := "Capital!"
	sym := NewString(testValue)
	if result := sym.StringValue(); result != testValue {
		t.Errorf("Got %s, want %s", result, testValue)
	}
}

func TestBasicSetSymbolInt(t *testing.T) {
	startValue, testValue := 4, 13
	sym := NewInt(startValue)
	if result := sym.IntValue(); result != startValue {
		t.Errorf("Got %d, want %d", result, startValue)
	}

	sym.SetInt(testValue)
	if result := sym.IntValue(); result != testValue {
		t.Errorf("Got %d, want %d", result, testValue)
	}
}

func TestBasicSetSymbolReal(t *testing.T) {
	startValue, testValue := 9.5, 22.5
	sym := NewReal(startValue)
	if result := sym.RealValue(); result != startValue {
		t.Errorf("Got %f, want %f", result, startValue)
	}

	sym.SetReal(testValue)
	if result := sym.RealValue(); result != testValue {
		t.Errorf("Got %f, want %f", result, testValue)
	}
}

func TestBasicSetSymbolBool(t *testing.T) {
	startValue, testValue := true, false
	sym := NewBool(startValue)
	if result := sym.BoolValue(); result != startValue {
		t.Errorf("Got %t, want %t", result, startValue)
	}

	sym.SetBool(testValue)
	if result := sym.BoolValue(); result != testValue {
		t.Errorf("Got %t, want %t", result, testValue)
	}
}

func TestBasicSetSymbolString(t *testing.T) {
	startValue, testValue := "Hello", "World!"
	sym := NewString(startValue)
	if result := sym.StringValue(); result != startValue {
		t.Errorf("Got %s, want %s", result, startValue)
	}

	sym.SetString(testValue)
	if result := sym.StringValue(); result != testValue {
		t.Errorf("Got %s, want %s", result, testValue)
	}
}

func TestSetSymbolIntFromOtherType(t *testing.T) {
	startValue, testValue := "turbot", 13
	sym := NewString(startValue)
	if result := sym.StringValue(); result != startValue {
		t.Errorf("Got %s, want %s", result, startValue)
	}

	sym.SetInt(testValue)
	if result := sym.IntValue(); result != testValue {
		t.Errorf("Got %d, want %d", result, testValue)
	}
}

func TestSetSymbolRealFromOtherType(t *testing.T) {
	startValue, testValue := 15, 93.5
	sym := NewInt(startValue)
	if result := sym.IntValue(); result != startValue {
		t.Errorf("Got %d, want %d", result, startValue)
	}

	sym.SetReal(testValue)
	if result := sym.RealValue(); result != testValue {
		t.Errorf("Got %f, want %f", result, testValue)
	}
}

func TestSetSymbolBoolFromOtherType(t *testing.T) {
	startValue, testValue := 47.25, false
	sym := NewReal(startValue)
	if result := sym.RealValue(); result != startValue {
		t.Errorf("Got %f, want %f", result, startValue)
	}

	sym.SetBool(testValue)
	if result := sym.BoolValue(); result != testValue {
		t.Errorf("Got %t, want %t", result, testValue)
	}
}

func TestSetSymbolStringFromOtherType(t *testing.T) {
	startValue, testValue := true, "World!"
	sym := NewBool(startValue)
	if result := sym.BoolValue(); result != startValue {
		t.Errorf("Got %t, want %t", result, startValue)
	}

	sym.SetString(testValue)
	if result := sym.StringValue(); result != testValue {
		t.Errorf("Got %s, want %s", result, testValue)
	}
}
