package securehttp

import (
	"io"
	"log"
)

// ResponseBodyCloser is a helper function that closes a body that implements io.ReadCloser.
//
// The intent is to instead of using defer body.Close() and handling the error everytime, that
// we call this function deferred everytime like:
//
// e.g. => defer ResponseBodyCloser(resp.Body)
func ResponseBodyCloser(body io.ReadCloser) {
	if err := body.Close(); err != nil {
		log.Println("Error closing response body:", err)
	}
}
