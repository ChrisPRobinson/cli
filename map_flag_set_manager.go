// Context is a type that is passed through to
package cli

import "time"

type MapFlagSetManager struct {
	wrappedFsm FlagSetManager
	valueMap   map[string]interface{}
}

// Determines if the flag was actually set
func (fsm *MapFlagSetManager) HasFlag(name string) bool {
	return fsm.wrappedFsm.HasFlag(name)
}

func (fsm *MapFlagSetManager) IsDefaultValueSet(name string) bool {
	return fsm.wrappedFsm.IsDefaultValueSet(name)
}

func (fsm *MapFlagSetManager) IsEnvVarSet(name string) bool {
	return fsm.wrappedFsm.IsEnvVarSet(name)
}

// Determines if the flag was actually set
func (fsm *MapFlagSetManager) IsSet(name string) bool {
	if fsm.wrappedFsm.IsSet(name) {
		return true
	}

	_, exists := fsm.valueMap[name]
	return exists
}

// Returns the number of flags set
func (fsm *MapFlagSetManager) NumFlags() int {
	return fsm.wrappedFsm.NumFlags()
}

// Returns the command line arguments associated with the context.
func (fsm *MapFlagSetManager) Args() Args {
	return fsm.wrappedFsm.Args()
}

var intValue = 1

func (fsm *MapFlagSetManager) Int(name string) int {
	value := fsm.wrappedFsm.Int(name)
	if !fsm.IsEnvVarSet(name) && (value == 0 || fsm.wrappedFsm.IsDefaultValueSet(name)) {
		otherValue, isType := fsm.valueMap[name].(int)
		if isType {
			return otherValue
		}
	}

	return value
}

func (fsm *MapFlagSetManager) Duration(name string) time.Duration {
	value := fsm.wrappedFsm.Duration(name)
	if !fsm.IsEnvVarSet(name) && (value == 0 || fsm.wrappedFsm.IsDefaultValueSet(name)) {
		otherValue, isType := fsm.valueMap[name].(time.Duration)
		if isType {
			return otherValue
		}
	}

	return value
}

func (fsm *MapFlagSetManager) Float64(name string) float64 {
	value := fsm.wrappedFsm.Float64(name)
	if !fsm.IsEnvVarSet(name) && (value == 0 || fsm.wrappedFsm.IsDefaultValueSet(name)) {
		otherValue, isType := fsm.valueMap[name].(float64)
		if isType {
			return otherValue
		}
	}

	return value
}

func (fsm *MapFlagSetManager) String(name string) string {
	value := fsm.wrappedFsm.String(name)
	if !fsm.IsEnvVarSet(name) && (value == "" || fsm.wrappedFsm.IsDefaultValueSet(name)) {
		otherValue, isType := fsm.valueMap[name].(string)
		if isType {
			return otherValue
		}
	}

	return value
}

func (fsm *MapFlagSetManager) StringSlice(name string) []string {
	value := fsm.wrappedFsm.StringSlice(name)
	if !fsm.IsEnvVarSet(name) && (value == nil || fsm.wrappedFsm.IsDefaultValueSet(name)) {
		otherValue, isType := fsm.valueMap[name].([]string)
		if isType {
			return otherValue
		}
	}

	return value
}

func (fsm *MapFlagSetManager) IntSlice(name string) []int {
	value := fsm.wrappedFsm.IntSlice(name)
	if !fsm.IsEnvVarSet(name) && (value == nil || fsm.wrappedFsm.IsDefaultValueSet(name)) {
		otherValue, isType := fsm.valueMap[name].([]int)
		if isType {
			return otherValue
		}
	}

	return value
}

func (fsm *MapFlagSetManager) Generic(name string) interface{} {
	value := fsm.wrappedFsm.Generic(name)
	if !fsm.IsEnvVarSet(name) && (value == nil || fsm.wrappedFsm.IsDefaultValueSet(name)) {
		return fsm.valueMap[name].(time.Duration)
	}

	return value
}

func (fsm *MapFlagSetManager) Bool(name string) bool {
	value := fsm.wrappedFsm.Bool(name)
	if !fsm.IsEnvVarSet(name) && !value {
		otherValue, isType := fsm.valueMap[name].(bool)
		if isType {
			return otherValue
		}
	}

	return value
}

func (fsm *MapFlagSetManager) BoolT(name string) bool {
	value := fsm.wrappedFsm.BoolT(name)
	if !fsm.IsEnvVarSet(name) && value {
		otherValue, isType := fsm.valueMap[name].(bool)
		if isType {
			return otherValue
		}
	}

	return value
}
