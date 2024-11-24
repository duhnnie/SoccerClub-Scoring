package expression

type expressionExpType[T ~string] struct {
	expType T
}

func (s *expressionExpType[T]) GetType() string {
	return string(s.expType)
}
