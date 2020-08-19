package lowercase

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"
	"regexp"
	"net/http"
	"text/template"
)

// Config the plugin configuration.
type Config struct {
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
	}
}

// Create a lowercase plugin.
type Lowercase struct {
	next     http.Handler
	name     string
}

// New created a new lowercase plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Lowercase{
		next:     next,
		name:     name,
	}, nil
}

func (a *Lowercase) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var uppercaseChars = regexp.MustCompile(`.*[A-Z].*`)
	var path = req.URL.Path
	if (uppercaseChars.MatchString(path)) {
    log.Printf("Lowercase Middleware fixing uppercase characters: %s", strings.ToLower(path))
		http.Redirect(rw, req, strings.ToLower(path) , http.StatusMovedPermanently)
	} else {
		a.next.ServeHTTP(rw, req)
	}
}
