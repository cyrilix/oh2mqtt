package cli

import (
	"os"
	"testing"
)

func TestSetDefaultValueFromEnv(t *testing.T) {
	err := os.Setenv("KEY1", "value1")
	if err != nil {
		t.Errorf("unable to set env value: %v", err)
	}

	cases := []struct {
		key      string
		defValue string
		expected string
	}{
		{"MISSING_KEY", "default", "default"},
		{"KEY1", "bad value", "value1"},
	}

	for _, c := range cases {
		var value = ""
		SetDefaultValueFromEnv(&value, c.key, c.defValue)
		if c.expected != value {
			t.Errorf("SetDefaultValueFromEnv(*value, %v, %v): %v, wants %v", c.key, c.defValue, value, c.expected)
		}
	}
}

func TestSetIntDefaultValueFromEnv(t *testing.T) {
	err := os.Setenv("KEY1", "12")
	err = os.Setenv("BAD_VALUE", "bad value")
	if err != nil {
		t.Errorf("unable to set env value: %v", err)
	}

	cases := []struct {
		key       string
		defValue  int
		expected  int
		withError bool
	}{
		{"MISSING_KEY", 3, 3, false},
		{"KEY1", 5, 12, false},
		{"BAD_VALUE", 5, 0, true},
	}

	for _, c := range cases {
		var value = 0
		err := SetIntDefaultValueFromEnv(&value, c.key, c.defValue)
		if err != nil {
			if !c.withError {
				t.Errorf("SetIntDefaultValueFromEnv(*value, %v, %v): %v", c.key, c.defValue, err)
			}
		}
		if c.withError && err == nil {
			t.Errorf("SetIntDefaultValueFromEnv(*value, %v, %v): %v, wants an error", c.key, c.defValue, value)
		} else if c.expected != value {
			t.Errorf("SetDefaultValueFromEnv(*value, %v, %v): %v, wants %v", c.key, c.defValue, value, c.expected)
		}
	}
}

func TestSetFloat64DefaultValueFromEnv(t *testing.T) {
	err := os.Setenv("KEY1", "12.2")
	err = os.Setenv("KEY2", "12")
	err = os.Setenv("BAD_VALUE", "bad value")
	if err != nil {
		t.Errorf("unable to set env value: %v", err)
	}

	cases := []struct {
		key       string
		defValue  float64
		expected  float64
		withError bool
	}{
		{"MISSING_KEY", 3.3, 3.3, false},
		{"KEY1", 5, 12.2, false},
		{"KEY2", 5, 12., false},
		{"BAD_VALUE", 5, 0, true},
	}

	for _, c := range cases {
		var value = 0.
		err := SetFloat64DefaultValueFromEnv(&value, c.key, c.defValue)
		if err != nil {
			if !c.withError {
				t.Errorf("SetFloat64DefaultValueFromEnv(*value, %v, %v): %v", c.key, c.defValue, err)
			}
		}
		if c.withError && err == nil {
			t.Errorf("SetFloat64DefaultValueFromEnv(*value, %v, %v): %v, wants an error", c.key, c.defValue, value)
		} else if c.expected != value {
			t.Errorf("SetFloat64DefaultValueFromEnv(*value, %v, %v): %v, wants %v", c.key, c.defValue, value, c.expected)
		}
	}
}
