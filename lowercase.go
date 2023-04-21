package lowercase

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

var uppercaseChars = regexp.MustCompile(`.*[A-Z].*`)

type Config struct {
}

func CreateConfig() *Config {
	return &Config{}
}

type Lowercase struct {
	next http.Handler
	name string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Lowercase{
		next: next,
		name: name,
	}, nil
}

func (a *Lowercase) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if uppercaseChars.MatchString(req.URL.Path) {
		http.Redirect(rw, req, strings.ToLower(req.URL.Path) + "?" + req.URL.RawQuery, http.StatusMovedPermanently)
	} else {
		a.next.ServeHTTP(rw, req)
	}
}
