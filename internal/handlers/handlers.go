package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/niutingyuan/calculate"
	"github.com/niutingyuan/calculate/internal/service"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req calculate.RequestFactorial
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid input"}`, http.StatusBadRequest)
		return
	}

	factorialA, factorialB := service.CalculateFactorials(req.A, req.B)
	res := calculate.ResponseFactorial{
		FactorialA: factorialA,
		FactorialB: factorialB,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
