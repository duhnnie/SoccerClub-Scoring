package scoring

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	NoScoringItemFound = Error("no specified scoring item found")
)