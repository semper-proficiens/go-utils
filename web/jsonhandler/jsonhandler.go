package jsonhandler

import (
	"encoding/json"
	"io"
	"net/http"
)

// UnmarshalJSONResponse takes an HTTP response and a pointer to a struct,
// unmarshals the JSON from the response body into the struct, and returns an error if any.
func UnmarshalJSONResponse(resp *http.Response, result interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, result); err != nil {
		return err
	}
	return nil
}
