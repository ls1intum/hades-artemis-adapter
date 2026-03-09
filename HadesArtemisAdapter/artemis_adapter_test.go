package main

import (
	"context"
	"testing"
)

func TestNewAdapter(t *testing.T) {
	cfg := AdapterConfig{
		APIPort:           "8082",
		ArtemisBaseURL:    "http://localhost:8080",
		NewResultEndpoint: "api/test",
		ArtemisAuthToken:  "test-token",
	}

	ctx := context.Background()
	adapter := NewAdapter(ctx, cfg)

	if adapter == nil {
		t.Fatal("NewAdapter returned nil")
	}

	if adapter.cfg.APIPort != "8082" {
		t.Errorf("Expected port 8082, got %s", adapter.cfg.APIPort)
	}

	if adapter.httpClient == nil {
		t.Error("HTTP client should not be nil")
	}
}

func TestAdapterStoreResults(t *testing.T) {
	cfg := AdapterConfig{
		APIPort:           "8082",
		ArtemisBaseURL:    "http://localhost:8080",
		NewResultEndpoint: "api/test",
		ArtemisAuthToken:  "test-token",
	}

	ctx := context.Background()
	adapter := NewAdapter(ctx, cfg)

	testResult := ResultDTO{
		ResultMetadata: ResultMetadata{
			JobName:                  "test-job",
			UUID:                     "test-job-123",
			AssignmentRepoBranchName: "main",
			IsBuildSuccessful:        true,
			Passed:                   5,
		},
		Results: []TestSuiteDTO{},
	}

	// Store results (will not send since no logs exist yet)
	err := adapter.StoreResults("test-job-123", testResult)
	// Should not error, just won't send yet
	if err != nil {
		t.Errorf("StoreResults should not error when only results exist: %v", err)
	}

	// Verify results were stored
	val, ok := adapter.results.Load("test-job-123")
	if !ok {
		t.Fatal("Results were not stored")
	}

	results := val.(ResultDTO)
	if results.JobName != "test-job" {
		t.Errorf("Expected job name 'test-job', got '%s'", results.JobName)
	}

	if results.Passed != 5 {
		t.Errorf("Expected 5 passed tests, got %d", results.Passed)
	}
}
