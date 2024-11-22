package variable

type Error string

func (err Error) Error() string {
	return string(err)
}

const (
	ErrorResolveInvalidParams     = Error("second argument needs to be a string or a slice of strings.")
	ErrorResolveInvalidFirstParam = Error("\"target\" parameter should be a \"map[string]interface{}\" when second argument is not an empty string nor empty string slice")
	ErrorCantResolveToType        = Error("can't resolve to specified type")
	ErrorNoValueFound             = Error("no value for variable was found")
)
