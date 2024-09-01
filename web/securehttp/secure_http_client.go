package securehttp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"time"
)

// systemCertPool is a variable that points to x509.SystemCertPool by default.
var systemCertPool = x509.SystemCertPool

// CustomHTTPClientInterface defines the methods for the secure HTTP client.
type CustomHTTPClientInterface interface {
	Get(url string) (*http.Response, error)
}

// CustomHTTPClient is a struct that holds the HTTP client.
type CustomHTTPClient struct {
	client *http.Client
	ctx    context.Context
}

// NewSecureHTTPClient creates a new HTTP client with secure settings based on a given context
func NewSecureHTTPClient(ctx context.Context) (*CustomHTTPClient, error) {
	// Load system CA certificates
	rootCAs, err := systemCertPool()
	if err != nil {
		rootCAs = x509.NewCertPool()
	}

	// Create a custom TLS configuration
	tlsConfig := &tls.Config{
		RootCAs:          rootCAs,
		MinVersion:       tls.VersionTLS12, // Use TLS 1.2 or higher
		CurvePreferences: []tls.CurveID{tls.CurveP256, tls.X25519},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_CHACHA20_POLY1305_SHA256,
		},
	}

	// Create an HTTP transport with the custom TLS configuration
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
		// Enable HTTP/2
		ForceAttemptHTTP2: true,
		// Set other transport settings
		MaxIdleConns:       100,
		IdleConnTimeout:    90 * time.Second,
		DisableCompression: false,
	}

	// Create and return the HTTP client with the custom transport
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second, // Set a timeout for the client
	}

	return &CustomHTTPClient{
		client: client,
		ctx:    ctx,
	}, nil
}

// Get is a utility function for making secure HTTP requests. This assumes "Content-Type" is json.
func (sc *CustomHTTPClient) Get(url string) (*http.Response, error) {
	// Create a new HTTP request with context
	req, err := http.NewRequestWithContext(sc.ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating new HTTP request: %w", err)
	}

	// Set common headers (if any)
	req.Header.Set("Content-Type", "application/json")

	// Make the HTTP request
	resp, err := sc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	return resp, nil
}
