package lowercase

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"regexp"
	"net/http"
	"text/template"
)

type Config struct {
}


func CreateConfig() *Config {
	return &Config{
	}
}

type Lowercase struct {
	next     http.Handler
	name     string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Lowercase{
		next:     next,
		name:     name,
	}, nil
}

func (a *Lowercase) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var uppercaseChars = regexp.MustCompile(`.*[A-Z].*`)
	if (uppercaseChars.MatchString(req.URL.Path)) {
		http.Redirect(rw, req, strings.ToLower(req.URL.Path) , http.StatusMovedPermanently)
	} else {
		a.next.ServeHTTP(rw, req)
	}
}
