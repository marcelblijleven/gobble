package pkg

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

const (
	maxBytes = int64(1048576)
)

// ReadRequestJSON decodes the incoming request body into the provided
// target, ensuring the request size doesn't exceed the maximum
// allowed bytes
func ReadRequestJSON(w http.ResponseWriter, r *http.Request, target any) error {
	// Ensure request bodies aren't too long
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(target); err != nil {
		return err
	}

	return nil
}

// ReadResponseJSON decodes the incoming response body into the provided target
func ReadResponseJSON(r *http.Response, target any) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(target); err != nil {
		return err
	}

	return nil
}

// ReadResponseXML decodes the incoming response body into the provided target
func ReadResponseXML(r *http.Response, target any) error {
	defer r.Body.Close()
	decoder := xml.NewDecoder(r.Body)

	if err := decoder.Decode(target); err != nil {
		return err
	}

	return nil
}
