package main_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	web "github.com/namikaze-dev/cyoa/cmd/web"
	"github.com/namikaze-dev/cyoa/internal/store"
)

func TestHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()

		handler := web.Handler{
			Store: testStoryStore,
		}
		handler.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assertEqual(t, resp.StatusCode, http.StatusOK)

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("unexpected error from io.ReadAll: %v", err)
		}
		body = bytes.TrimSpace(body)

		want := "Title 1"
		if !bytes.Contains(body, []byte(want)) {
			t.Errorf("want body containing %q, got %q", want, string(body))
		}
		
		want = "<html"
		if !bytes.Contains(body, []byte(want)) {
			t.Errorf("want body containing %q, got %q", want, string(body))
		}
	})

	t.Run("bad request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/bad-arc-name", nil)
		rr := httptest.NewRecorder()

		handler := web.Handler{
			Store: testStoryStore,
		}
		handler.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assertEqual(t, resp.StatusCode, http.StatusBadRequest)
	})
}

func assertEqual(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v want %#v", got, want)
	}
}

var testStoryStore = map[string]store.Story{
	"intro": {
		Title: "Title 1",
		Story: []string{
			"story line 1",
			"story line 2",
			"story line 3",
		},
		Options: []store.StoryOption{
			{
				Text: "text 1",
				Arc:  "valley",
			},
		},
	},
	"valley": {
		Title: "Valley of Code",
		Story: []string{
			"story line 1",
			"story line 2",
		},
		Options: []store.StoryOption{
			{
				Text: "start",
				Arc:  "intro",
			},
			{
				Text: "end",
				Arc:  "end-arc",
			},
		},
	},
}
