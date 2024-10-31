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
	// addition
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
