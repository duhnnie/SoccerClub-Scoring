package scoring

import "fmt"

type ScoringItemNotFoundError string

func (e ScoringItemNotFoundError) Error() string {
	return fmt.Sprintf("no \"%s\" scoring item found", string(e))
}

type NotSupporteEngineError string

func (e NotSupporteEngineError) Error() string {
	return fmt.Sprintf("engine \"%s\" not supported", string(e))
}

type UnmarshalRepositoryError struct {
	index int
	err   error
}

func (e *UnmarshalRepositoryError) Error() string {
	return fmt.Sprintf("at item %d: %s", e.index, e.err)
}

type DoesntResolveToBooleanError string

func (e DoesntResolveToBooleanError) Error() string {
	return fmt.Sprintf("expression \"%s\" doesn't resolve to boolean", string(e))
}

type JExpItemError struct {
	Path string
	Err  error
}

func (e *JExpItemError) Error() string {
	return fmt.Sprintf("%s: %s", e.Path, e.Err)
}
