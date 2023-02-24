package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"weehongayden/finance-flow-app/internal/config"
)

func TestServe(t *testing.T) {
	mockServerConfig := config.Server{
		Address: "localhost",
		Port:    12345,
	}

	// Create a mock config with a random port number
	mockConfig := config.Config{
		Server: mockServerConfig,
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	go serve(mockConfig)

	time.Sleep(1 * time.Second) // Give the server time to start up.

	resp, err := http.Get(fmt.Sprintf("http://%s/api/server-health", server.Listener.Addr()))
	if err != nil {
		t.Fatalf("Failed to make GET request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
}
