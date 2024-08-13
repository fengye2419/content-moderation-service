package api

import (
	"encoding/json"
	"net/http"

	"github.com/fengye2419/content-moderation-service/internal/service"
)

// ValidateRequest represents the request body for the /validate endpoint.
// @Description ValidateRequest represents the request body for the /validate endpoint.
// @Description It contains a keyword that needs to be validated against sensitive words.
// @Description The request body is a JSON object with one field: "keyword".
// @Description Example: { "keyword": "This is a badword1 example." }
type ValidateRequest struct {
	Keyword string `json:"keyword"`
}

// ValidateResponse represents the response body for the /validate endpoint.
// @Description ValidateResponse represents the response body for the /validate endpoint.
// @Description It contains a list of sensitive words that were matched and the count of matches.
// @Description The response body is a JSON object with two fields: "sensitive_words" and "match_count".
// @Description Example: { "sensitive_words": ["badword1"], "match_count": 1 }
type ValidateResponse struct {
	SensitiveWords []string `json:"sensitive_words"`
	MatchCount     int      `json:"match_count"`
}

// ValidateHandler handles the /validate endpoint.
// @Summary Validate keywords against sensitive words
// @Description ValidateHandler handles the /validate endpoint.
// @Description It receives a request with keywords and checks if they match any sensitive words.
// @Description It returns the matched sensitive words and the count of matches.
// @Tags sensitive-words
// @Accept  json
// @Produce  json
// @Param   request body ValidateRequest true "Request body"
// @Success 200 {object} ValidateResponse
// @Failure 400 {string} string "Bad request"
// @Router /sensitive-words/validate [post]
func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	var req ValidateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sensitiveWords, matchCount := service.CheckSensitiveWords(req.Keyword)

	resp := ValidateResponse{
		SensitiveWords: sensitiveWords,
		MatchCount:     matchCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
