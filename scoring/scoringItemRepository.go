package scoring

import (
	"encoding/json"
	"fmt"
)

type ItemRepository struct {
	items map[string]*Item
}

func NewItemRepository(items map[string]*Item) *ItemRepository {
	return &ItemRepository{
		items: items,
	}
}

func NewItemRepositoryFromData(data []byte) (*ItemRepository, error) {
	items := make(map[string]*Item)
	dictionary := make(map[string]scoringItemBridge)

	if err := json.Unmarshal(data, &dictionary); err != nil {
		return nil, err
	}

	for key, value := range dictionary {
		items[key] = value.toScoringItem()
	}

	return NewItemRepository(items), nil
}

func (r *ItemRepository) ExecuteItem(itemID string) (bool, error) {
	if _, exists := r.items[itemID]; !exists {
		return false, fmt.Errorf("no \"%s\" item found", itemID)
	} else {
		// TODO: call item.resolve()
		return true, nil
	}
}
