package util

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
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

// ErrorFromResponse is a helper function which creates an error from a response
// for which the status code is not in the 2xx range.
// The status code should be checked before calling this function, but to be sure this
// function also does the check and returns nil if the response has a status code in the
// 2xx range.
func ErrorFromResponse(r *http.Response) error {
	if r.StatusCode/100 == 2 {
		return nil
	}

	return fmt.Errorf("request to %+v unsuccessful: %d %s", r.Request.URL.String(), r.StatusCode, r.Status)
}

// debugResponseBody shouldn't be used in other cases than during development
func debugResponseBody(res *http.Response) {
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}
