package stores

import (
	"errors"
	"slices"

	"github.com/google/uuid"
)

type MemoryItemStore struct {
	items []Item
}

func NewMemoryItemStore() ItemStore {
	return &MemoryItemStore{
		items: []Item{},
	}
}

func (m *MemoryItemStore) GetAllItems() ([]Item, error) {
	return m.items, nil
}

func (m *MemoryItemStore) GetItem(ID string) (*Item, error) {
	idx := slices.IndexFunc(m.items, func(i Item) bool { return i.ID == ID })
	if idx < 0 {
		return nil, errors.New("ID not found")
	}
	return &m.items[idx], nil
}

func (m *MemoryItemStore) CreateItem(request CreateItemRequest) (*Item, error) {
	item := Item{ID: uuid.NewString(), Name: request.Name, Description: request.Description}
	m.items = append(m.items, item)
	return &item, nil
}

func (m *MemoryItemStore) UpdateItem(ID string, request CreateItemRequest) (*Item, error) {
	idx := slices.IndexFunc(m.items, func(i Item) bool { return i.ID == ID })
	if idx < 0 {
		return nil, errors.New("ID not found")
	}

	m.items[idx] = Item{
		ID:          ID,
		Name:        request.Name,
		Description: request.Description,
	}

	return &m.items[idx], nil
}
func (m *MemoryItemStore) DeleteItem(ID string) error {
	idx := slices.IndexFunc(m.items, func(i Item) bool { return i.ID == ID })
	if idx < 0 {
		return errors.New("ID not found")
	}
	m.items = slices.Delete(m.items, idx, 1)

	return nil
}
