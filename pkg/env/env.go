package env

import (
	"errors"
)

type VarType byte

const (
	STRING VarType = iota
	BOOL
	INT
)

type Var struct {
	Name         string
	Description  string
	DefaultValue string
	Type         VarType
}

var envVars = make(map[string]Var)

func NewStringVar(name, description, defaultValue string) error {
	return register(name, description, defaultValue, STRING)
}

func NewBoolVar(name, description, defaultValue string) error {
	return register(name, description, defaultValue, BOOL)
}

func NewIntVar(name, description, defaultValue string) error {
	return register(name, description, defaultValue, INT)
}

func ForceNewStringVar(name, description, defaultValue string) {
	forceRegister(name, description, defaultValue, STRING)
}

func ForceNewBoolVar(name, description, defaultValue string) {
	forceRegister(name, description, defaultValue, BOOL)
}

func ForceNewIntVar(name, description, defaultValue string) {
	forceRegister(name, description, defaultValue, INT)
}

func register(name, description, defaultValue string, varType VarType) error {
	if _, ok := envVars[name]; ok {
		return errors.New("already allocated env name")
	}
	envVars[name] = Var{
		Name:         name,
		Description:  description,
		DefaultValue: defaultValue,
		Type:         varType,
	}
	return nil
}

func forceRegister(name, description, defaultValue string, varType VarType) {
	envVars[name] = Var{
		Name:         name,
		Description:  description,
		DefaultValue: defaultValue,
		Type:         varType,
	}
}

func Get(name string) Var {
	return envVars[name]
}
