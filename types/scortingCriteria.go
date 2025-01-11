package types

import "errors"

// TODO: consider using only int64 instead of generic
type ScoringCriteria struct {
	criteria map[string]float64
}

func NewScoringCriteria() *ScoringCriteria {
	return &ScoringCriteria{}
}

func (c *ScoringCriteria) Set(name string, points float64) {
	c.criteria[name] = points
}

func (c *ScoringCriteria) GetScore(name string) (float64, error) {
	if v, exists := c.criteria[name]; !exists {
		// TODO: improve error
		return 0.0, errors.New("not exists")
	} else {
		return v, nil
	}
}

func (c *ScoringCriteria) Get(name string) (interface{}, error) {
	return c.GetScore(name)
}

func (c *ScoringCriteria) TopMap() map[string]interface{} {
	res := map[string]interface{}{}

	for k, v := range c.criteria {
		res[k] = v
	}

	return res
}
