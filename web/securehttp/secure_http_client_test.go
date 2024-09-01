package securehttp

import (
	"context"
	"crypto/x509"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// mockSystemCertPool is a mock implementation of x509.SystemCertPool.
func mockSystemCertPool() (*x509.CertPool, error) {
	return nil, errors.New("mock error")
}

// init replaces the real SystemCertPool with the mock implementation.
func init() {
	systemCertPool = mockSystemCertPool
}

// MockTransport is custom transport for testing purposes.
type MockTransport struct {
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req)
}

func TestNewSecureHTTPClient(t *testing.T) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := NewSecureHTTPClient(ctx)
	if err != nil {
		t.Fatalf("Failed to create secure HTTP client: %v", err)
	}

	if client == nil {
		t.Fatal("Expected non-nil client")
	}
}

func TestSecureHTTPClient_Get_Success(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"message": "success"}`)); err != nil {
			t.Fatalf("Failed to write response: %v", err)
		}
	}))
	defer ts.Close()

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a new secure HTTP client
	client, err := NewSecureHTTPClient(ctx)
	if err != nil {
		t.Fatalf("Failed to create secure HTTP client: %v", err)
	}

	// Make a GET request
	resp, err := client.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			t.Fatalf("Failed to close response body: %v", err)
		}
	}()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestSecureHTTPClient_Get_RequestError(t *testing.T) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a new secure HTTP client
	client, err := NewSecureHTTPClient(ctx)
	if err != nil {
		t.Fatalf("Failed to create secure HTTP client: %v", err)
	}

	// Make a GET request with an invalid URL
	_, err = client.Get("http://invalid-url")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestSecureHTTPClient_Get_DoError(t *testing.T) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a new secure HTTP client with a mock transport
	client, err := NewSecureHTTPClient(ctx)
	if err != nil {
		t.Fatalf("Failed to create secure HTTP client: %v", err)
	}

	client.client.Transport = &MockTransport{
		RoundTripFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("mock error")
		},
	}

	// Make a GET request
	_, err = client.Get("http://example.com")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestSecureHTTPClient_Get_Non2xxStatusCode(t *testing.T) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	// Create a new secure HTTP client
	client, err := NewSecureHTTPClient(ctx)
	if err != nil {
		t.Fatalf("Failed to create secure HTTP client: %v", err)
	}

	// Make a GET request
	resp, err := client.Get(ts.URL)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if resp != nil {
		defer func() {
			if err = resp.Body.Close(); err != nil {
				t.Fatalf("Failed to close response body: %v", err)
			}
		}()
	}
}

func TestNewSecureHTTPClient_SystemCertPoolError(t *testing.T) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// This test will use the mock implementation of systemCertPool
	client, err := NewSecureHTTPClient(ctx)
	if err != nil {
		t.Fatalf("Failed to create secure HTTP client: %v", err)
	}

	if client == nil {
		t.Fatal("Expected non-nil client")
	}
}

func TestSecureHTTPClient_Get_InvalidURL(t *testing.T) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a new secure HTTP client
	client, err := NewSecureHTTPClient(ctx)
	if err != nil {
		t.Fatalf("Failed to create secure HTTP client: %v", err)
	}

	// Make a GET request with an invalid URL format
	_, err = client.Get("http://%41:8080/") // Invalid URL
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if err.Error() != "error creating new HTTP request: parse \"http://%41:8080/\": invalid URL escape \"%41\"" {
		t.Fatalf("Unexpected error message: %v", err)
	}
}
