package traefik_template_plugin

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"text/template"

	"github.com/xabinapal/traefik-template-plugin/internal/config"
)

// CreateConfig creates the default plugin configuration.
func CreateConfig() *config.Config {
	return config.New()
}

// Plugin is the plugin instance.
type Plugin struct {
	name     string
	next     http.Handler
	headers  map[string]string
	template *template.Template
}

// New creates a new plugin instance.
func New(ctx context.Context, next http.Handler, config *config.Config, name string) (http.Handler, error) {
	if len(config.Headers) == 0 {
		return nil, fmt.Errorf("headers cannot be empty")
	}

	return &Plugin{
		name:     name,
		next:     next,
		headers:  config.Headers,
		template: template.New("demo").Delims("[[", "]]"),
	}, nil
}

// ServeHTTP is the main plugin function that will be called by Traefik.
func (p *Plugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for key, value := range p.headers {
		tmpl, err := p.template.Parse(value)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		writer := &bytes.Buffer{}

		err = tmpl.Execute(writer, req)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		req.Header.Set(key, writer.String())
	}

	p.next.ServeHTTP(rw, req)
}
