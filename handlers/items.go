package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hs-heilbronn-devsecops/acetlisto/stores"
)

type ItemHandler struct {
	store stores.ItemStore
}

func NewItemHandler(store stores.ItemStore) *ItemHandler {
	return &ItemHandler{
		store: store,
	}
}

func (i *ItemHandler) listItems(w http.ResponseWriter, r *http.Request) {
	items, err := i.store.GetAllItems()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not load items")
	}
	respondWithJSON(w, http.StatusOK, items)
}

func (i *ItemHandler) createItem(w http.ResponseWriter, r *http.Request) {
	var request stores.CreateItemRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer deferredClose(r.Body)

	item, err := i.store.CreateItem(request)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not store item")
	}

	respondWithJSON(w, http.StatusCreated, item)
}

func (i *ItemHandler) getItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item, err := i.store.GetItem(vars["ID"])

	if err != nil {
		respondWithError(w, http.StatusNotFound, "")
		return
	}

	respondWithJSON(w, http.StatusOK, item)
}

func (i *ItemHandler) updateItem(w http.ResponseWriter, r *http.Request) {
	var request stores.CreateItemRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer deferredClose(r.Body)

	vars := mux.Vars(r)
	item, err := i.store.UpdateItem(vars["ID"], request)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "")
		return
	}

	respondWithJSON(w, http.StatusOK, item)
}

func (i *ItemHandler) deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := i.store.DeleteItem(vars["ID"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "")
	}
}
