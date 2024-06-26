package integration

import (
	"testing"
)

// Unit test
func TestSendPostRequestPersonal(t *testing.T) {
	// Run the function with the URL of the mock server
	body, err := sendPostRequest("http://localhost:8080/admin/deductions/personal", `{"amount": 60000}`)
	if err != nil {
		t.Fatal(err)
	}

	expectedBody := `{"personalDeduction":60000}`
	if body != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, body)
	}
}
