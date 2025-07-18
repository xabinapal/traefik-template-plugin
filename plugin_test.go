package traefik_template_plugin_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	plugin "github.com/xabinapal/traefik-template-plugin"
)

func TestNew(t *testing.T) {
	t.Run("empty headers", func(t *testing.T) {
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

		cfg := plugin.CreateConfig()
		_, err := plugin.New(context.Background(), next, cfg, "traefik-template-plugin")
		if err == nil {
			t.Fatal("expected error to be not nil")
		}
	})

	t.Run("valid headers", func(t *testing.T) {
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

		cfg := plugin.CreateConfig()
		cfg.Headers = map[string]string{
			"X-Demo": "test",
			"X-Url":  "http://localhost",
		}

		p, err := plugin.New(context.Background(), next, cfg, "traefik-template-plugin")
		if err != nil {
			t.Fatal(err)
		}

		if p == nil {
			t.Fatal("expected plugin to be not nil")
		}
	})
}

func TestServeHTTP(t *testing.T) {
	nextCalled := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true

		if r.Header.Get("X-Demo") != "test" {
			t.Fatal("expected X-Demo header to be set")
		}

		if r.Header.Get("X-Url") != "http://localhost" {
			t.Fatal("expected X-Url header to be set")
		}

		w.WriteHeader(http.StatusOK)
	})

	cfg := plugin.CreateConfig()
	cfg.Headers = map[string]string{
		"X-Demo": "test",
		"X-Url":  "http://localhost",
	}

	p, err := plugin.New(context.Background(), next, cfg, "traefik-template-plugin")
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://localhost", nil)

	p.ServeHTTP(rw, req)

	if !nextCalled {
		t.Fatal("expected next to be called")
	}
}
