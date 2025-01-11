package match

type Match struct {
	id    string
	title string
	facts map[string]interface{}
}

func New(id, title string, facts map[string]interface{}) *Match {
	return &Match{id, title, facts}
}

func (m *Match) GetFacts() map[string]interface{} {
	return m.facts
}
