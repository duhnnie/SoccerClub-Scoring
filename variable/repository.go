package variable

import (
	"strings"
)

type Repository struct {
	variables map[string]interface{}
}

func NewRepo() *Repository {
	return &Repository{
		variables: make(map[string]interface{}),
	}
}

func Resolve(target, path interface{}) (interface{}, error) {
	var newPath []string

	if target == nil {
		return nil, ErrorTargetIsNIL
	}

	if pathArray, ok := path.([]string); !ok {
		if pathString, ok := path.(string); !ok {
			return nil, ErrorResolveInvalidParams
		} else if len(pathString) == 0 {
			return target, nil
		} else {
			newPath = strings.Split(pathString, ".")
		}
	} else if len(pathArray) == 0 {
		return target, nil
	} else {
		newPath = pathArray
	}

	if targetMap, ok := target.(map[string]interface{}); !ok {
		return nil, ErrorResolveInvalidFirstParam
	} else {
		newTarget := targetMap[newPath[0]]
		return Resolve(newTarget, newPath[1:])
	}
}

func (r *Repository) Set(name string, value interface{}) {
	r.variables[name] = value
}

func (r *Repository) get(variableName string) (interface{}, error) {
	path := strings.Split(variableName, ".")
	name := path[0]
	targetObject := r.variables[name]

	if targetObject == nil {
		return nil, ErrorNoVariableFound(variableName)
	}

	res, err := Resolve(targetObject, path[1:])

	if err != nil && err == ErrorTargetIsNIL {
		return nil, ErrorNoVariableFound(variableName)
	}

	return res, err
}

func (r *Repository) GetInt64(variableName string) (int64, error) {
	if resolvedValue, err := r.GetFloat64(variableName); err == nil {
		return int64(resolvedValue), nil
	} else if resolvedValue, err := r.get(variableName); err != nil {
		return 0, err
	} else if intValue, ok := resolvedValue.(int64); !ok {
		return 0, ErrorCantResolveToType{"int64", variableName}
	} else {
		return intValue, nil
	}
}

func (r *Repository) GetFloat64(variableName string) (float64, error) {
	if resolvedValue, err := r.get(variableName); err != nil {
		return 0, err
	} else if floatValue, ok := resolvedValue.(float64); !ok {
		return 0, ErrorCantResolveToType{"float64", variableName}
	} else {
		return floatValue, nil
	}
}

func (r *Repository) GetBool(variableName string) (bool, error) {
	if resolvedValue, err := r.get(variableName); err != nil {
		return false, err
	} else if boolValue, ok := resolvedValue.(bool); !ok {
		return false, ErrorCantResolveToType{"bool", variableName}
	} else {
		return boolValue, nil
	}
}

func (r *Repository) GetString(variableName string) (string, error) {
	if resolvedValue, err := r.get(variableName); err != nil {
		return "", err
	} else if stringValue, ok := resolvedValue.(string); !ok {
		return "", ErrorCantResolveToType{"string", variableName}
	} else {
		return stringValue, nil
	}
}
