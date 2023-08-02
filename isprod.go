package isprod

import (
	"fmt"
	"os"
	"strings"
)

// DefaultConditions is a list of conditions that are used by default.
// It's initialized at package init.
// Rules:
// If environment variable 'prod' is set and its value is not one of [false], consider it as production environment.
// If environment variable 'production' is set and its value is not one of [false], consider it as production environment.
// If environment variable 'staging' is set and its value is not one of [false], consider it as production environment.
// If environment variable 'live' is set and its value is not one of [false], consider it as production environment.
// If environment variable 'ci' is set and its value is not one of [false], consider it as production environment.
// If environment variable 'PROD' is set and its value is not one of [false], consider it as production environment.
// If environment variable 'PRODUCTION' is set and its value is not one of [false], consider it as production environment.
// If environment variable 'STAGING' is set and its value is not one of [false], consider it as production environment.
// If environment variable 'LIVE' is set and its value is not one of [false], consider it as production environment.
// If environment variable 'CI' is set and its value is not one of [false], consider it as production environment.
// If environment variable 'env' is set and its value is one of [prod, production, staging, live, ci, PROD, PRODUCTION, STAGING, LIVE, CI], consider it as production environment.
// If environment variable 'environment' is set and its value is one of [prod, production, staging, live, ci, PROD, PRODUCTION, STAGING, LIVE, CI], consider it as production environment.
// If environment variable 'mode' is set and its value is one of [prod, production, staging, live, ci, PROD, PRODUCTION, STAGING, LIVE, CI], consider it as production environment.
// If environment variable 'ENV' is set and its value is one of [prod, production, staging, live, ci, PROD, PRODUCTION, STAGING, LIVE, CI], consider it as production environment.
// If environment variable 'ENVIRONMENT' is set and its value is one of [prod, production, staging, live, ci, PROD, PRODUCTION, STAGING, LIVE, CI], consider it as production environment.
// If environment variable 'MODE' is set and its value is one of [prod, production, staging, live, ci, PROD, PRODUCTION, STAGING, LIVE, CI], consider it as production environment.
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

func (c Condition) Check() bool {
	value, exists := os.LookupEnv(c.EnvVarName)
	if !exists {
		return false
	}
	value = strings.ToLower(value)

	if c.AllowAnyValue {
		if c.ExcludedValues != nil {
			for _, excludedValue := range c.ExcludedValues {
				if value == strings.ToLower(excludedValue) {
					return false
				}
			}
		}
		return true
	}

	if c.AllowedValues != nil {
		for _, allowedValue := range c.AllowedValues {
			if value == strings.ToLower(allowedValue) {
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

type Conditions []Condition

func (c *Conditions) Add(condition Condition) {
	*c = append(*c, condition)
}

func (c Conditions) Check() bool {
	for _, condition := range c {
		if condition.Check() {
			return true
		}
	}

	return false
}

func (c Conditions) String() string {
	var conditionStrings []string
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
