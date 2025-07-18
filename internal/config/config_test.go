package config_test

import (
	"testing"

	"github.com/xabinapal/traefik-template-plugin/internal/config"
)

func TestNew(t *testing.T) {
	cfg := config.New()

	if cfg == nil {
		t.Fatal("expected config to be not nil")
	}

	if len(cfg.Headers) != 0 {
		t.Fatal("expected headers to be empty")
	}
}
