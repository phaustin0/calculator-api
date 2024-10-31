package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Operation struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type Result struct {
	Result int `json:"result"`
}

func add(w http.ResponseWriter, r *http.Request) {
	operation := new(Operation)
	decodeRequestBody(w, r, operation)

	answer := operation.Number1 + operation.Number2
	result := &Result{
		Result: answer,
	}

	encodeRequestBody(w, r, result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func subtract(w http.ResponseWriter, r *http.Request) {
	operation := new(Operation)
	decodeRequestBody(w, r, operation)

	answer := operation.Number1 - operation.Number2
	result := &Result{
		Result: answer,
	}

	encodeRequestBody(w, r, result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func multiply(w http.ResponseWriter, r *http.Request) {}

func divide(w http.ResponseWriter, r *http.Request) {}

func sum(w http.ResponseWriter, r *http.Request) {}

func decodeRequestBody(w http.ResponseWriter, r *http.Request, v any) {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		fmt.Errorf("[ERROR]: unable to decode request body, reason: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func encodeRequestBody(w http.ResponseWriter, r *http.Request, v any) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		fmt.Errorf("[ERROR]: unable to encode payload, reason: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
