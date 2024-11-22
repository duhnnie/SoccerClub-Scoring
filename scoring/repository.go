package scoring

import (
	"encoding/json"
)

type Repository struct {
	items map[string]*Item
}

func NewRepository(items map[string]*Item) *Repository {
	return &Repository{
		items: items,
	}
}

func NewRepositoryFromData(data []byte) (*Repository, error) {
	items := make(map[string]*Item)
	dictionary := make(map[string]itemBridge)

	if err := json.Unmarshal(data, &dictionary); err != nil {
		return nil, err
	}

	for key, value := range dictionary {
		items[key] = value.toScoringItem()
	}

	return NewRepository(items), nil
}

func (r *Repository) ExecuteItem(itemID string) (bool, error) {
	if _, exists := r.items[itemID]; !exists {
		return false, NoScoringItemFound
	} else {
		// TODO: call item.resolve()
		return true, nil
	}
}
