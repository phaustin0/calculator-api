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
	var err error
	var operation Operation

	err = json.NewDecoder(r.Body).Decode(&operation)
	if err != nil {
		fmt.Errorf("[ERROR]: unable to decode request body, reason: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	answer := operation.Number1 + operation.Number2
	result := &Result{
		Result: answer,
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		fmt.Errorf("[ERROR]: unable to encode payload, reason: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func subtract(w http.ResponseWriter, r *http.Request) {}

func multiply(w http.ResponseWriter, r *http.Request) {}

func divide(w http.ResponseWriter, r *http.Request) {}

func sum(w http.ResponseWriter, r *http.Request) {}
