package match

import "encoding/json"

type Match struct {
	id    string
	title string
	facts json.RawMessage
}

func New(id, title string, facts json.RawMessage) *Match {
	return &Match{id, title, facts}
}

func (m *Match) GetID() string {
	return m.id
}

func (m *Match) GetFacts() json.RawMessage {
	return m.facts
}
