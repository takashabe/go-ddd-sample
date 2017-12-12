package interfaces

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	fixture "github.com/takashabe/go-fixture"
)

func loadFixture(t *testing.T, conn *sql.DB, file string) {
	fixture, err := fixture.NewFixture(conn, "mysql")
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	err = fixture.Load(file)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
}

func prepareServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(Routes())
}

func sendRequest(t *testing.T, method, url string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	return res
}

func TestGetUser(t *testing.T) {
	ts := prepareServer(t)
	defer ts.Close()

	cases := []struct {
		input      int
		expectJSON []byte
		expectCode int
	}{
		{
			1,
			[]byte(`{"id":1,"name":"satoshi"}`),
			http.StatusOK,
		},
		{
			0,
			nil,
			http.StatusNotFound,
		},
	}
	for i, c := range cases {
		url := fmt.Sprintf("%s/user/%d", ts.URL, c.input)
		res := sendRequest(t, "GET", url, nil)
		defer res.Body.Close()

		if c.expectCode != res.StatusCode {
			t.Errorf("#%d: want %d, got %d", i, c.expectCode, res.StatusCode)
		}
		if res.StatusCode != http.StatusOK {
			continue
		}

		payload, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("#%d: want non error, got %#v", i, err)
		}
		if !reflect.DeepEqual(c.expectJSON, payload) {
			t.Errorf("#%d: want %s, got %s", i, c.expectJSON, payload)
		}
	}
}

func TestGetUsers(t *testing.T) {
	ts := prepareServer(t)
	defer ts.Close()

	cases := []struct {
		expectJSON []byte
		expectCode int
	}{
		{
			[]byte(`[{"id":1,"name":"satoshi"},{"id":2,"name":"kasumi"}]`),
			http.StatusOK,
		},
	}
	for i, c := range cases {
		url := fmt.Sprintf("%s/users", ts.URL)
		res := sendRequest(t, "GET", url, nil)
		defer res.Body.Close()

		if c.expectCode != res.StatusCode {
			t.Errorf("#%d: want %d, got %d", i, c.expectCode, res.StatusCode)
		}
		if res.StatusCode != http.StatusOK {
			continue
		}

		payload, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("#%d: want non error, got %#v", i, err)
		}
		if !reflect.DeepEqual(c.expectJSON, payload) {
			t.Errorf("#%d: want %s, got %s", i, c.expectJSON, payload)
		}
	}
}
