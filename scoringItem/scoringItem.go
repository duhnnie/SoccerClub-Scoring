package scoringItem

import "encoding/json"

type ScoringItem struct {
	id          string
	name        string
	description string
	expression  map[string]interface{}
}

func New(id, name, description string, expression map[string]interface{}) *ScoringItem {
	return &ScoringItem{
		id:          id,
		name:        name,
		description: description,
		expression:  expression,
	}
}

func NewFromData(id string, data []byte) (*ScoringItem, error) {
	var scoringItem ScoringItem
	bridge := scoringItemBridge{}

	err := json.Unmarshal(data, &bridge)
	bridge.ID = id
	bridge.toScoringItem(&scoringItem)

	return &scoringItem, err
}

func NewFromString(id, data string) (*ScoringItem, error) {
	return NewFromData(id, []byte(data))
}

func (s *ScoringItem) ToJSON() ([]byte, error) {
	bridge := bridgeFromScoreItem(*s)
	return json.Marshal(bridge)
}

func (s *ScoringItem) GetID() string {
	return s.id
}

func (s *ScoringItem) GetName() string {
	return s.name
}

func (s *ScoringItem) GetDescription() string {
	return s.description
}

func (s *ScoringItem) GetExpression() map[string]interface{} {
	return s.expression
}
