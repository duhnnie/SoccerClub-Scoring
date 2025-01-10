package scoring

import "github.com/duhnnie/soccerclub-scoring/types"

type Item interface {
	GetID() string
	Resolve(variables types.VariableContainer) (bool, error)
}
