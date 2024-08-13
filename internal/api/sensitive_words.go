package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fengye2419/content-moderation-service/internal/model"
	"github.com/go-chi/chi/v5"
)

// AddSensitiveWordRequest represents the request body for adding a sensitive word.
type AddSensitiveWordRequest struct {
	Word string `json:"word"`
}

// AddSensitiveWordHandler handles the addition of a new sensitive word.
// @Summary Add a new sensitive word
// @Description AddSensitiveWordHandler handles the addition of a new sensitive word.
// @Description It receives a request with a word that needs to be added to the list of sensitive words.
// @Description The request body is a JSON object with one field: "word".
// @Description Example: { "word": "badword1" }
// @Tags sensitive-words
// @Accept  json
// @Produce  json
// @Param   request body AddSensitiveWordRequest true "Request body"
// @Success 201 {string} string "Created"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /sensitive-words [post]
func AddSensitiveWordHandler(w http.ResponseWriter, r *http.Request) {
	var request AddSensitiveWordRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := model.AddSensitiveWord(request.Word); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DeleteSensitiveWordHandler handles the deletion of a sensitive word by ID.
// @Summary Delete a sensitive word by ID
// @Description DeleteSensitiveWordHandler handles the deletion of a sensitive word by ID.
// @Description It receives a request with the ID of the sensitive word that needs to be deleted.
// @Tags sensitive-words
// @Param   id path int true "Sensitive word ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /sensitive-words/{id} [delete]
func DeleteSensitiveWordHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := model.DeleteSensitiveWord(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
