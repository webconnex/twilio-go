package twilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

// setup sets up a test HTTP server along with a github.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Twilio, mux *http.ServeMux, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	url, _ := url.Parse(server.URL)
	// client is the GitHub client being tested and is
	// configured to use test server.
	client = NewClient("AC123", "AuthToken")
	client.defaultbaseURL = String(url.String())

	return client, mux, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()

	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, want url.Values) {
	t.Helper()

	err := r.ParseForm()

	if err != nil {
		fmt.Printf("Err: %v", err)
	}

	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}

func testQueryValues(t *testing.T, r *http.Request, values values) {
	t.Helper()

	want := url.Values{}

	for k, v := range values {
		want.Set(k, v)
	}

	if got := r.URL.Query(); !reflect.DeepEqual(got, want) {
		t.Errorf("Request query: %v, want %v", got, want)
	}
}

// Test whether the marshaling of v produces JSON that corresponds
// to the want string.
func testJSONMarshal(t *testing.T, v interface{}, want string) {
	t.Helper()
	// Unmarshal the wanted JSON, to verify its correctness, and marshal it back
	// to sort the keys.
	u := reflect.New(reflect.TypeOf(v)).Interface()
	if err := json.Unmarshal([]byte(want), &u); err != nil {
		t.Errorf("Unable to unmarshal JSON for %v: %v", want, err)
	}

	w, err := json.Marshal(u)
	if err != nil {
		t.Errorf("Unable to marshal JSON for %#v", u)
	}

	// Marshal the target value.
	j, err := json.Marshal(v)
	if err != nil {
		t.Errorf("Unable to marshal JSON for %#v", v)
	}

	if string(w) != string(j) {
		t.Errorf("json.Marshal(%q) returned %s, want %s", v, j, w)
	}
}
