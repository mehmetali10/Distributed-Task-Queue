package transporthelper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const token = "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjI3LCJuYW1lIjoiTXVyYXQiLCJzdXJuYW1lIjoiWWFsxLFuYXlhayIsInBob25lIjoiKzkwNTM0NTg4OTgwOCIsImVtYWlsIjoibXVyYXR5YWzEsW5heWFrQG1hcy5jb20iLCJzZXNzaW9uVXVpZCI6IjEyMTdjNDY3LTVmYzEtNGVhMy1iNWQ4LTc3ODUzNWViNzIwOCIsImV4cCI6MTcwMzc5NTY2OH0.BAXLKtqO7S29E54QPd1dfBIJi0sgWZu7JDPVpQVI6Wk"

type BaseTemp struct {
	Server    *httptest.Server
	Url       string
	Want      int
	Method    string
	BodyParam interface{}
}

// TestWithoutToken performs an HTTP request without an authorization token for testing purposes.
// It sends an HTTP request using the provided method, URL, and body parameters.
// If any errors occur during JSON conversion, sending the request, or receiving the response,
// it logs the error and returns nil.
// It also checks if the response status code matches the expected status code (b.Want) and logs any mismatches.
// The function returns the JSON response map or nil if there were errors.
func TestWithoutToken(t *testing.T, b BaseTemp) map[string]interface{} {
	jsonData, err := json.Marshal(b.BodyParam)
	if err != nil {
		t.Errorf("An error occured during the json convertation: %v", err)
		return nil
	}

	client := http.Client{}
	req, _ := http.NewRequest(b.Method, b.Url, bytes.NewReader(jsonData))
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("An error occured during the sending request: %v", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != b.Want {
		t.Errorf("want %d result %d", b.Want, resp.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	var jsonRes map[string]interface{}
	err = json.Unmarshal(body, &jsonRes)

	return jsonRes
}

// TestWithToken performs an HTTP request with an authorization token for testing purposes.
// It follows a similar logic to TestWithoutToken but includes an Authorization header with the token.
func TestWithToken(t *testing.T, b BaseTemp) map[string]interface{} {
	jsonData, err := json.Marshal(b.BodyParam)
	if err != nil {
		t.Errorf("An error occured during the json convertation: %v", err)
		return nil
	}

	client := http.Client{}
	req, _ := http.NewRequest(b.Method, b.Url, bytes.NewReader(jsonData))
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("An error occured during the sending request: %v", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != b.Want {
		t.Errorf("want %d result %d", b.Want, resp.StatusCode)
		return nil
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var jsonRes map[string]interface{}
	err = json.Unmarshal(body, &jsonRes)

	return jsonRes
}
