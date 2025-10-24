package stores

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialState(t *testing.T) {
	m := NewMemoryItemStore()

	items, err := m.GetAllItems()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(items), "Store should initially have no items")
}

func TestCreateItem(t *testing.T) {
	m := NewMemoryItemStore()

	item, err := m.CreateItem(CreateItemRequest{
		Name:        "Oil",
		Description: "Some beautiful oil.",
	})
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Greater(t, len(item.ID), 6, "ID must be longer than 6 characters.")
	assert.Equal(t, "Oil", item.Name)
	assert.Equal(t, "Some beautiful oil.", item.Description)
}

func TestCreateAndGetItem(t *testing.T) {
	m := NewMemoryItemStore()

	item, _ := m.CreateItem(CreateItemRequest{
		Name:        "Milk",
		Description: "2 liters",
	})

	loadedItem, err := m.GetItem(item.ID)
	assert.NoError(t, err)
	assert.NotNil(t, loadedItem)
	assert.Equal(t, item, loadedItem)
}

func TestGetItem_NotFound(t *testing.T) {
	m := NewMemoryItemStore()

	item, err := m.GetItem("non-existent-id")
	assert.Error(t, err)
	assert.Nil(t, item)
}

func TestUpdateItem(t *testing.T) {
	m := NewMemoryItemStore()

	// Create an item first
	item, _ := m.CreateItem(CreateItemRequest{
		Name:        "Bread",
		Description: "White bread",
	})

	// Update it
	updateReq := CreateItemRequest{
		Name:        "Brown Bread",
		Description: "Whole grain",
	}
	updatedItem, err := m.UpdateItem(item.ID, updateReq)
	assert.NoError(t, err)
	assert.Equal(t, "Brown Bread", updatedItem.Name)
	assert.Equal(t, "Whole grain", updatedItem.Description)

	// Fetch and verify
	got, err := m.GetItem(item.ID)
	assert.NoError(t, err)
	assert.Equal(t, updatedItem, got)
}

func TestUpdateItem_NotFound(t *testing.T) {
	m := NewMemoryItemStore()

	updateReq := CreateItemRequest{
		Name:        "Test",
		Description: "Test desc",
	}
	item, err := m.UpdateItem("non-existent-id", updateReq)
	assert.Error(t, err)
	assert.Nil(t, item)
}

func TestDeleteItem(t *testing.T) {
	m := NewMemoryItemStore()

	// Create an item first
	item, _ := m.CreateItem(CreateItemRequest{
		Name:        "Juice",
		Description: "Apple juice",
	})

	// Delete it
	err := m.DeleteItem(item.ID)
	assert.NoError(t, err)

	// Verify deletion
	got, err := m.GetItem(item.ID)
	assert.Error(t, err)
	assert.Nil(t, got)
}

func TestDeleteItem_NotFound(t *testing.T) {
	m := NewMemoryItemStore()

	err := m.DeleteItem("non-existent-id")
	assert.Error(t, err)
}

func TestGetAllItems(t *testing.T) {
	m := NewMemoryItemStore()

	// Add multiple items
	item1, _ := m.CreateItem(CreateItemRequest{Name: "A", Description: "First"})
	item2, _ := m.CreateItem(CreateItemRequest{Name: "B", Description: "Second"})

	items, err := m.GetAllItems()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(items), 2)
	assert.Contains(t, items, *item1)
	assert.Contains(t, items, *item2)
}
