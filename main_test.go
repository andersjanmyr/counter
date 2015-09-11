package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCounter(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/Counter", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	counter(w, req)

	fmt.Printf("%d - %s", w.Code, w.Body.String())
	actualCode := w.Code
	expectedCode := 200
	if actualCode != expectedCode {
		t.Errorf("w.Code: %v, expected %v", actualCode, expectedCode)
	}
	actualBody := w.Body.String()
	expectedBody := "counter"
	if actualBody != expectedBody {
		t.Errorf("w.Body: %v, expected %v", actualBody, expectedBody)
	}
}
