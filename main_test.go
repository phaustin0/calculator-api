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
