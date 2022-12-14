package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/steps798/gorestapi/pkg/mocks"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	// run LoopingArrayExamples func
	loop()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Articles)
}
