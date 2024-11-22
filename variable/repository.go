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
		return nil, ErrorNoValueFound
	}

	return Resolve(targetObject, path[1:])
}

func (r *Repository) GetInt(variableName string) (int, error) {
	if resolvedValue, err := r.get(variableName); err != nil {
		return 0, err
	} else if intValue, ok := resolvedValue.(int); !ok {
		return 0, ErrorCantResolveToType
	} else {
		return intValue, nil
	}
}

func (r *Repository) GetFloat32(variableName string) (float32, error) {
	if resolvedValue, err := r.get(variableName); err != nil {
		return 0, err
	} else if floatValue, ok := resolvedValue.(float32); !ok {
		return 0, ErrorCantResolveToType
	} else {
		return floatValue, nil
	}
}

func (r *Repository) GetBool(variableName string) (bool, error) {
	if resolvedValue, err := r.get(variableName); err != nil {
		return false, err
	} else if boolValue, ok := resolvedValue.(bool); !ok {
		return false, ErrorCantResolveToType
	} else {
		return boolValue, nil
	}
}

func (r *Repository) GetString(variableName string) (string, error) {
	if resolvedValue, err := r.get(variableName); err != nil {
		return "", err
	} else if stringValue, ok := resolvedValue.(string); !ok {
		return "", ErrorCantResolveToType
	} else {
		return stringValue, nil
	}
}
