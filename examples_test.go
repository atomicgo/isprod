package isprod_test

import (
	"atomicgo.dev/isprod"
	"fmt"
	"os"
)

func ExampleCheck() {
	os.Setenv("PRODUCTION", "true") // Many common names are supported. See DefaultConditions.
	fmt.Println(isprod.Check())
	// Output: true
}

func ExampleCondition_Check() {
	os.Setenv("MY_ENV_VAR", "live")

	cond := isprod.Condition{
		EnvVarName:     "MY_ENV_VAR",
		AllowAnyValue:  true,
		ExcludedValues: []string{"false"},
	}

	fmt.Println(cond.Check())
	// Output: true
}

func ExampleConditions_Add() {
	var conds isprod.Conditions

	cond1 := isprod.Condition{
		EnvVarName:    "ENV_VAR_1",
		AllowAnyValue: true,
	}
	cond2 := isprod.Condition{
		EnvVarName:    "ENV_VAR_2",
		AllowAnyValue: true,
	}

	conds.Add(cond1)
	conds.Add(cond2)

	fmt.Println(len(conds))
	// Output: 2
}

func ExampleConditions_Check() {
	os.Setenv("ENV_VAR_1", "true")

	var conds isprod.Conditions

	cond1 := isprod.Condition{
		EnvVarName:    "ENV_VAR_1",
		AllowAnyValue: true,
	}
	cond2 := isprod.Condition{
		EnvVarName:    "ENV_VAR_2",
		AllowAnyValue: true,
	}

	conds.Add(cond1)
	conds.Add(cond2)

	fmt.Println(conds.Check())
	// Output: true
}
