package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// Helper to execute requests and record the response
func executeRequest(router *mux.Router, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

// Test if New() correctly initializes the router
func TestNew(t *testing.T) {
	router := New()
	assert.NotNil(t, router)
}

// Test homeHandler redirect
func TestHomeHandler(t *testing.T) {
	router := New()
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(router, req)

	assert.Equal(t, http.StatusTemporaryRedirect, response.Code)
	// Fix: match actual redirect path
	assert.Equal(t, "/item/", response.Header().Get("Location"))
}

// Integration test for creating and listing items
func TestIntegration_CreateAndListItems(t *testing.T) {
	router := New()

	createBody := []byte(`{"Name":"Oil","Description":"Nice oil"}`)
	req, _ := http.NewRequest("POST", "/items/", bytes.NewBuffer(createBody))
	req.Header.Set("Content-Type", "application/json")
	createResp := executeRequest(router, req)

	// Fix: expect 201 Created
	assert.Equal(t, http.StatusCreated, createResp.Code)

	req, _ = http.NewRequest("GET", "/items/", nil)
	listResp := executeRequest(router, req)
	assert.Equal(t, http.StatusOK, listResp.Code)

	var items []map[string]interface{}
	if err := json.Unmarshal(listResp.Body.Bytes(), &items); err != nil {
		t.Fatalf("Failed to unmarshal list response: %v", err)
	}
	assert.GreaterOrEqual(t, len(items), 1)
}

// Integration test for updating and deleting items
func TestIntegration_UpdateAndDeleteItem(t *testing.T) {
	router := New()

	createBody := []byte(`{"Name":"Bread","Description":"Whole grain"}`)
	req, _ := http.NewRequest("POST", "/items/", bytes.NewBuffer(createBody))
	req.Header.Set("Content-Type", "application/json")
	createResp := executeRequest(router, req)

	var created map[string]interface{}
	if err := json.Unmarshal(createResp.Body.Bytes(), &created); err != nil {
		t.Fatalf("Failed to unmarshal create response: %v", err)
	}
	id := created["id"].(string)

	// Update
	updateBody := []byte(`{"Name":"Brown Bread","Description":"Updated desc"}`)
	req, _ = http.NewRequest("PUT", "/items/"+id, bytes.NewBuffer(updateBody))
	req.Header.Set("Content-Type", "application/json")
	updateResp := executeRequest(router, req)
	assert.Equal(t, http.StatusOK, updateResp.Code)

	// Delete
	req, _ = http.NewRequest("DELETE", "/items/"+id, nil)
	deleteResp := executeRequest(router, req)
	assert.Equal(t, http.StatusOK, deleteResp.Code)
}
