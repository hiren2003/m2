package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hs-heilbronn-devsecops/acetlisto/stores"
)

func New() *mux.Router {
//
	r := mux.NewRouter()

	s := stores.NewMemoryItemStore()

	c := NewItemHandler(s)

	// Routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/items/", c.listItems).Methods("GET")
	r.HandleFunc("/items/", c.createItem).Methods("POST")
	r.HandleFunc("/items/{ID}", c.getItem).Methods("GET")
	r.HandleFunc("/items/{ID}", c.updateItem).Methods("PUT")
	r.HandleFunc("/items/{ID}", c.deleteItem).Methods("DELETE")

	return r
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/item/", http.StatusTemporaryRedirect)
}
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Println("Error writing response", err)
	}
}

func deferredClose(c io.ReadCloser) {
	err := c.Close()
	if err != nil {
		log.Println("Error closing request", err)
		}
}
		

