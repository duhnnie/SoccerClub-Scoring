package expression

type expressionExpType struct {
	expType string
}

func (s *expressionExpType) GetType() string {
	return string(s.expType)
}
