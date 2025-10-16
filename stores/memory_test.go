package stores

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialState(t *testing.T) {
	m := NewMemoryItemStore()

	items, _ := m.GetAllItems()
	assert.Equal(t, 0, len(items), "Store should initially have no items")

}

func TestCreate(t *testing.T) {
	m := NewMemoryItemStore()

	item, err := m.CreateItem(CreateItemRequest{
		Name:        "Oil",
		Description: "Some beautiful oil.",
	})
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Greater(t, len(item.ID), 6, "ID must be longer than 6 characters.")
}

func TestCreateAndLoad(t *testing.T) {
	m := NewMemoryItemStore()

	item, _ := m.CreateItem(CreateItemRequest{
		Name:        "Oil",
		Description: "Some beautiful oil.",
	})

	loadedItem, err := m.GetItem(item.ID)
	assert.NoError(t, err)
	assert.NotNil(t, loadedItem)
	assert.Equal(t, loadedItem, item)
}
