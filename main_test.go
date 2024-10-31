package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	url := "http://localhost:8000/add"
	jsonReq := []byte(`{"number1": 4, "number2": 6}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var result Result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, result, Result{Result: 10})
}

func TestSubtraction(t *testing.T) {
	url := "http://localhost:8000/subtract"

	// negative result
	jsonReq := []byte(`{"number1": 4, "number2": 6}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var result Result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, result, Result{Result: -2})

	// zero result
	jsonReq = []byte(`{"number1": 6, "number2": 6}`)
	req, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, result, Result{Result: 0})

	// positive result
	jsonReq = []byte(`{"number1": 100, "number2": 6}`)
	req, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, result, Result{Result: 94})
}

func TestMultiplication(t *testing.T) {
	url := "http://localhost:8000/multiply"

	// negative and positive
	jsonReq := []byte(`{"number1": -4, "number2": 6}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var result Result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, result, Result{Result: -24})

	// negative and negative
	jsonReq = []byte(`{"number1": -6, "number2": -6}`)
	req, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, result, Result{Result: 36})

	// positive and positive
	jsonReq = []byte(`{"number1": 100, "number2": 6}`)
	req, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, result, Result{Result: 600})
}

func TestDivision(t *testing.T) {
	url := "http://localhost:8000/divide"

	// negative
	jsonReq := []byte(`{"number1": -24, "number2": 6}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var result Result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, result, Result{Result: -4})

	// negative and negative
	jsonReq = []byte(`{"number1": -6, "number2": -6}`)
	req, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, result, Result{Result: 1})

	// divide by zero
	jsonReq = []byte(`{"number1": 100, "number2": 0}`)
	req, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	var divErr DivisionByZeroError
	err = json.NewDecoder(resp.Body).Decode(&divErr)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, divErr, DivisionByZeroError{Error: "cannot divide by zero"})
}
