package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteAlert(t *testing.T) {
	// Create a request with sample JSON data for the alert
	reqData := `{
        "alert_id": "b950482e9911ec7e41f7ca5e5d9a424f",
        "service_id": "my_test_service_id",
        "service_name": "my_test_service",
        "model": "my_test_model",
        "alert_type": "anomaly",
        "alert_ts": "1695644160",
        "severity": "warning",
        "team_slack": "slack_ch"
    }`
	reqBody := strings.NewReader(reqData)
	req, err := http.NewRequest("POST", "/alerts", reqBody)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	// Call the handler function
	WriteAlert(recorder, req)

	// Check the HTTP response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	// You can add more assertions to check the response body or verify the database state.
}

func TestReadAlerts(t *testing.T) {
	// Create a request with URL parameters
	req, err := http.NewRequest("GET", "/alerts?service_id=my_test_service_id&alert_ts=1695644060&alert_end_ts=1695644160", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	// Call the handler function
	ReadAlerts(recorder, req)

	// Check the HTTP response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	// You can add more assertions to check the response body or verify the database state.
}
