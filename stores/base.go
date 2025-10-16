package stores

type CreateItemRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ItemStore interface {
	GetAllItems() ([]Item, error)
	GetItem(ID string) (*Item, error)
	CreateItem(request CreateItemRequest) (*Item, error)
	UpdateItem(ID string, request CreateItemRequest) (*Item, error)
	DeleteItem(ID string) error
}
