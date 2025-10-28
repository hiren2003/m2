package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/hs-heilbronn-devsecops/acetlisto/stores"
	"github.com/stretchr/testify/assert"
)

// helper to run HTTP requests
func performRequest(r *mux.Router, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	return rec
}

// --- Test: CreateItem ---
func TestCreateItem(t *testing.T) {
	store := stores.NewMemoryItemStore()
	router := New(store)

	body := []byte(`{"Name":"Oil","Description":"Should be nice!"}`)
	resp := performRequest(router, "POST", "/items/", body)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal create response: %v", err)
	}

	assert.Contains(t, result, "id")
	assert.Equal(t, "Oil", result["name"])
	assert.Equal(t, "Should be nice!", result["description"])
}

// --- Test: ListItems ---
func TestListItems(t *testing.T) {
	store := stores.NewMemoryItemStore()
	router := New(store)

	body := []byte(`{"Name":"Milk","Description":"2 liters"}`)
	createResp := performRequest(router, "POST", "/items/", body)
	assert.Equal(t, http.StatusCreated, createResp.Code)

	listResp := performRequest(router, "GET", "/items/", nil)
	assert.Equal(t, http.StatusOK, listResp.Code)

	var items []map[string]interface{}
	if err := json.Unmarshal(listResp.Body.Bytes(), &items); err != nil {
		t.Fatalf("Failed to unmarshal list response: %v", err)
	}
	assert.GreaterOrEqual(t, len(items), 1)
}

// --- Test: GetItem ---
func TestGetItem_Handler(t *testing.T) {
	store := stores.NewMemoryItemStore()
	router := New(store)

	createBody := []byte(`{"Name":"Eggs","Description":"Free range"}`)
	createResp := performRequest(router, "POST", "/items/", createBody)
	assert.Equal(t, http.StatusCreated, createResp.Code)

	var created map[string]interface{}
	if err := json.Unmarshal(createResp.Body.Bytes(), &created); err != nil {
		t.Fatalf("Failed to unmarshal create response: %v", err)
	}
	id := created["id"].(string)

	getResp := performRequest(router, "GET", "/items/"+id, nil)
	assert.Equal(t, http.StatusOK, getResp.Code)

	var fetched map[string]interface{}
	if err := json.Unmarshal(getResp.Body.Bytes(), &fetched); err != nil {
		t.Fatalf("Failed to unmarshal get response: %v", err)
	}

	assert.Equal(t, created["id"], fetched["id"])
	assert.Equal(t, created["name"], fetched["name"])
	assert.Equal(t, created["description"], fetched["description"])
}

// --- Test: UpdateItem ---
func TestUpdateItem(t *testing.T) {
	store := stores.NewMemoryItemStore()
	router := New(store)

	createBody := []byte(`{"Name":"Bread","Description":"White bread"}`)
	createResp := performRequest(router, "POST", "/items/", createBody)
	assert.Equal(t, http.StatusCreated, createResp.Code)

	var created map[string]interface{}
	if err := json.Unmarshal(createResp.Body.Bytes(), &created); err != nil {
		t.Fatalf("Failed to unmarshal create response: %v", err)
	}
	id := created["id"].(string)

	updateBody := []byte(`{"Name":"Brown Bread","Description":"Whole grain bread"}`)
	updateResp := performRequest(router, "PUT", "/items/"+id, updateBody)
	assert.Equal(t, http.StatusOK, updateResp.Code)

	var updated map[string]interface{}
	if err := json.Unmarshal(updateResp.Body.Bytes(), &updated); err != nil {
		t.Fatalf("Failed to unmarshal update response: %v", err)
	}

	assert.Equal(t, "Brown Bread", updated["name"])
	assert.Equal(t, "Whole grain bread", updated["description"])
}

// --- Test: DeleteItem ---
func TestDeleteItem(t *testing.T) {
	store := stores.NewMemoryItemStore()
	router := New(store)

	createBody := []byte(`{"Name":"Juice","Description":"Apple juice"}`)
	createResp := performRequest(router, "POST", "/items/", createBody)
	assert.Equal(t, http.StatusCreated, createResp.Code)

	var created map[string]interface{}
	if err := json.Unmarshal(createResp.Body.Bytes(), &created); err != nil {
		t.Fatalf("Failed to unmarshal create response: %v", err)
	}
	id := created["id"].(string)

	deleteResp := performRequest(router, "DELETE", "/items/"+id, nil)
	assert.Equal(t, http.StatusOK, deleteResp.Code)

	getResp := performRequest(router, "GET", "/items/"+id, nil)
	assert.NotEqual(t, http.StatusOK, getResp.Code)
}

// --- Test: Invalid ID ---
func TestGetItem_NotFound_Handler(t *testing.T) {
	store := stores.NewMemoryItemStore()
	router := New(store)

	resp := performRequest(router, "GET", "/items/nonexistent-id", nil)
	assert.NotEqual(t, http.StatusOK, resp.Code)
}
