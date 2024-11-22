package variable

import (
	"errors"
	"fmt"
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
			return nil, errors.New("second argument needs to be a slice of strings or string")
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
		return nil, errors.New("\"target\" parameter should be a \"map[string]interface{}\" when second argument is not an empty string nor empty string slice")
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
		// TODO: define our own errors somewhere
		return nil, fmt.Errorf("No value for variable \"%s\"", variableName)
	}

	return Resolve(targetObject, path[1:])
}

func (r *Repository) GetInt(variableName string) (int, error) {
	if resolvedValue, err := r.get(variableName); err != nil {
		return 0, err
	} else if intValue, ok := resolvedValue.(int); !ok {
		return 0, errors.New("can't resolve to specified type")
	} else {
		return intValue, nil
	}
}

func (r *Repository) GetFloat32(variableName string) (float32, error) {
	if resolvedValue, err := r.get(variableName); err != nil {
		return 0, err
	} else if floatValue, ok := resolvedValue.(float32); !ok {
		return 0, errors.New("can't resolve to specified type")
	} else {
		return floatValue, nil
	}
}

func (r *Repository) GetBool(variableName string) (bool, error) {
	if resolvedValue, err := r.get(variableName); err != nil {
		return false, err
	} else if boolValue, ok := resolvedValue.(bool); !ok {
		return false, errors.New("can't resolve to specified type")
	} else {
		return boolValue, nil
	}
}

func (r *Repository) GetString(variableName string) (string, error) {
	if resolvedValue, err := r.get(variableName); err != nil {
		return "", err
	} else if stringValue, ok := resolvedValue.(string); !ok {
		return "", errors.New("can't resolve to specified type")
	} else {
		return stringValue, nil
	}
}
