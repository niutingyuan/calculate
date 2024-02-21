package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/niutingyuan/calculate/internal/service"
)

type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Response struct {
	FactorialA int `json:"a!"`
	FactorialB int `json:"b!"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid input"}`, http.StatusBadRequest)
	}

	result := service.CalculateFactorials(req.A, req.B)
	res := Response{
		FactorialA: result.FactorialA,
		FactorialB: result.FactorialB,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
