package isprod

import (
	"fmt"
	"os"
	"strings"
)

// DefaultConditions is a list of conditions that are used by default.
// It's initialized at package init.
var DefaultConditions Conditions

func init() {
	commonProdNames := []string{
		"prod",
		"production",
		"staging",
		"live",
		"ci",
	}

	commonEnvVariableNames := []string{
		"env",
		"environment",
		"mode",
	}

	// Add UPPERCASE versions of the common names.
	commonProdNamesUpper := make([]string, len(commonProdNames))
	for i, name := range commonProdNames {
		commonProdNamesUpper[i] = strings.ToUpper(name)
	}

	commonProdNames = append(commonProdNames, commonProdNamesUpper...)

	for _, name := range commonProdNames {
		DefaultConditions.Add(Condition{
			EnvVarName:     name,
			AllowedValues:  nil,
			AllowAnyValue:  true,
			ExcludedValues: []string{"false"},
		})
	}

	// Add UPPERCASE versions of the common names.
	commonEnvVariableNamesUpper := make([]string, len(commonEnvVariableNames))
	for i, name := range commonEnvVariableNames {
		commonEnvVariableNamesUpper[i] = strings.ToUpper(name)
	}

	commonEnvVariableNames = append(commonEnvVariableNames, commonEnvVariableNamesUpper...)

	for _, name := range commonEnvVariableNames {
		DefaultConditions.Add(Condition{
			EnvVarName:     name,
			AllowedValues:  commonProdNames,
			AllowAnyValue:  false,
			ExcludedValues: nil,
		})
	}
}

// Condition is a condition that checks if the environment is production.
type Condition struct {
	// EnvVarName is the name of the environment variable to check.
	EnvVarName string
	// AllowedValues is a list of values that are considered valid for the environment variable.
	AllowedValues []string
	// AllowAnyValue can be set to true if any value for the environment variable is allowed.
	AllowAnyValue bool
	// ExcludedValues is a list of values that are specifically not allowed, even if AllowAnyValue is set to true.
	ExcludedValues []string
}

// Check checks if the condition is met.
func (c Condition) Check() bool {
	value, exists := os.LookupEnv(c.EnvVarName)
	if !exists {
		return false
	}

	value = strings.ToLower(value)

	if c.AllowAnyValue {
		if c.ExcludedValues != nil {
			for _, excludedValue := range c.ExcludedValues {
				if strings.EqualFold(value, excludedValue) {
					return false
				}
			}
		}

		return true
	}

	if c.AllowedValues != nil {
		for _, allowedValue := range c.AllowedValues {
			if strings.EqualFold(value, allowedValue) {
				return true
			}
		}
	}

	return false
}

func (c Condition) String() string {
	if c.AllowAnyValue {
		return fmt.Sprintf("If environment variable '%s' is set and its value is not one of [%s], consider it as production environment.",
			c.EnvVarName, strings.Join(c.ExcludedValues, ", "))
	}

	return fmt.Sprintf("If environment variable '%s' is set and its value is one of [%s], consider it as production environment.",
		c.EnvVarName, strings.Join(c.AllowedValues, ", "))
}

// Conditions is a list of conditions.
type Conditions []Condition

// Add adds a condition to the list.
func (c *Conditions) Add(condition Condition) {
	*c = append(*c, condition)
}

// Check checks if any of the conditions is true.
func (c Conditions) Check() bool {
	for _, condition := range c {
		if condition.Check() {
			return true
		}
	}

	return false
}

// String returns a string representation of the conditions in plain english.
func (c Conditions) String() string {
	conditionStrings := make([]string, 0, len(c))

	for _, condition := range c {
		conditionStrings = append(conditionStrings, condition.String())
	}

	return strings.Join(conditionStrings, "\n")
}

// Check checks if the application is running in production or not.
// It uses the DefaultConditions.
// If you want to use your own conditions, use the Conditions.Check() method.
func Check() bool {
	return DefaultConditions.Check()
}
