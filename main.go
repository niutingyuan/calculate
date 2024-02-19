package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Response struct {
	FactorialA int `json:"a!"`
	FactorialB int `json:"b!"`
}

type result struct {
	value int
	id    int // Use id to distinguish between a! and b!
}

// Middleware to check if a and b exist and are non-negative
func validateInput(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error":"Incorrect input"}`, http.StatusBadRequest)
			return
		}
		if req.A < 0 || req.B < 0 {
			http.Error(w, `{"error":"Incorrect input"}`, http.StatusBadRequest)
			return
		}

		// Use context to pass the validated request to the next handler
		ctx := context.WithValue(r.Context(), "validatedReq", req)
		next(w, r.WithContext(ctx), p)
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func calculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req, ok := r.Context().Value("validatedReq").(Request)
	if !ok {
		http.Error(w, `{"error":"Request validation failed"}`, http.StatusInternalServerError)
		return
	}

	ch := make(chan result, 2)

	go func() {
		ch <- result{value: factorial(req.A), id: 1}
	}()

	go func() {
		ch <- result{value: factorial(req.B), id: 2}
	}()

	res := Response{}
	for i := 0; i < 2; i++ {
		result := <-ch
		if result.id == 1 {
			res.FactorialA = result.value
		} else if result.id == 2 {
			res.FactorialB = result.value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	router := httprouter.New()
	router.POST("/calculate", validateInput(calculateHandler))

	log.Fatal(http.ListenAndServe(":8989", router))
}
