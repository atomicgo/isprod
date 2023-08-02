package isprod_test

import (
	. "atomicgo.dev/isprod"
	"testing"
)

func TestConditionCheck(t *testing.T) {
	cond := Condition{
		EnvVarName:     "TEST_VAR",
		AllowAnyValue:  true,
		ExcludedValues: []string{"excluded"},
	}

	t.Setenv("TEST_VAR", "value")

	if !cond.Check() {
		t.Error("Expected true for allowed value, got false")
	}

	t.Setenv("TEST_VAR", "excluded")

	if cond.Check() {
		t.Error("Expected false for excluded value, got true")
	}

	cond = Condition{
		EnvVarName:    "TEST_VAR",
		AllowedValues: []string{"allowed"},
		AllowAnyValue: false,
	}

	t.Setenv("TEST_VAR", "allowed")

	if !cond.Check() {
		t.Error("Expected true for allowed value, got false")
	}

	t.Setenv("TEST_VAR", "other")

	if cond.Check() {
		t.Error("Expected false for not allowed value, got true")
	}
}

func TestConditionsCheck(t *testing.T) {
	conds := Conditions{
		{
			EnvVarName:     "TEST_VAR_1",
			AllowAnyValue:  true,
			ExcludedValues: []string{"excluded"},
		},
		{
			EnvVarName:    "TEST_VAR_2",
			AllowedValues: []string{"allowed"},
			AllowAnyValue: false,
		},
	}

	t.Setenv("TEST_VAR_1", "value")
	t.Setenv("TEST_VAR_2", "other")

	if !conds.Check() {
		t.Error("Expected true when one condition is satisfied, got false")
	}

	t.Setenv("TEST_VAR_1", "excluded")
	t.Setenv("TEST_VAR_2", "other")

	if conds.Check() {
		t.Error("Expected false when no conditions are satisfied, got true")
	}

	t.Setenv("TEST_VAR_1", "value")
	t.Setenv("TEST_VAR_2", "allowed")

	if !conds.Check() {
		t.Error("Expected true when both conditions are satisfied, got false")
	}
}
