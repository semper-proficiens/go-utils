package securehttp

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//// MockCertPoolLoader is a mock implementation of CertPoolLoader.
//type MockCertPoolLoader struct {
//	LoadSystemCertPoolFunc func() (*x509.CertPool, error)
//}
//
//func (m *MockCertPoolLoader) LoadSystemCertPool() (*x509.CertPool, error) {
//	return m.LoadSystemCertPoolFunc()
//}

// MockTransport is custom transport for testing purposes.
type MockTransport struct {
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req)
}

func TestNewSecureHTTPClient(t *testing.T) {
	client, err := NewSecureHTTPClient()
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

	// Create a new secure HTTP client
	client, err := NewSecureHTTPClient()
	if err != nil {
		t.Fatalf("Failed to create secure HTTP client: %v", err)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Make a GET request
	resp, err := client.Get(ctx, ts.URL)
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
	// Create a new secure HTTP client
	client, err := NewSecureHTTPClient()
	if err != nil {
		t.Fatalf("Failed to create secure HTTP client: %v", err)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Make a GET request with an invalid URL
	_, err = client.Get(ctx, "http://invalid-url")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestSecureHTTPClient_Get_DoError(t *testing.T) {
	// Create a new secure HTTP client with a mock transport
	client, err := NewSecureHTTPClient()
	if err != nil {
		t.Fatalf("Failed to create secure HTTP client: %v", err)
	}

	client.client.Transport = &MockTransport{
		RoundTripFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("mock error")
		},
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Make a GET request
	_, err = client.Get(ctx, "http://example.com")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestSecureHTTPClient_Get_Non2xxStatusCode(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	// Create a new secure HTTP client
	client, err := NewSecureHTTPClient()
	if err != nil {
		t.Fatalf("Failed to create secure HTTP client: %v", err)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Make a GET request
	resp, err := client.Get(ctx, ts.URL)
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

//func TestNewSecureHTTPClient_SystemCertPoolError(t *testing.T) {
//	// Mock CertPoolLoader to return an error
//	loader := &MockCertPoolLoader{
//		LoadSystemCertPoolFunc: func() (*x509.CertPool, error) {
//			return nil, errors.New("mock error")
//		},
//	}
//
//	client, err := NewSecureHTTPClient()
//	if err != nil {
//		t.Fatalf("Failed to create secure HTTP client: %v", err)
//	}
//
//	if client == nil {
//		t.Fatal("Expected non-nil client")
//	}
//}
