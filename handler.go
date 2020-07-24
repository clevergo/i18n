// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package i18n

import (
	"context"
	"net/http"
)

type contextKey int

const (
	translatorKey contextKey = iota
)

// GetTranslator returns a translator from context.
func GetTranslator(ctx context.Context) *Translator {
	t, _ := ctx.Value(translatorKey).(*Translator)
	return t
}

// Handler is a HTTP handler that store the matched translator in request context.
func Handler(ts *Translators, next http.Handler, parsers ...LanguageParser) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		langs := []string{}

		for _, parser := range parsers {
			if v := parser.Parse(r); v != "" {
				langs = append(langs, v)
			}
		}
		ctx := context.WithValue(r.Context(), translatorKey, ts.MatchTranslator(langs...))
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// Middleware is a HTTP middleware that store the matched translator in request context.
func Middleware(ts *Translators, parsers ...LanguageParser) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return Handler(ts, next, parsers...)
	}
}

// LanguageParser is an interface that defined how to retrieve language from a request.
type LanguageParser interface {
	Parse(*http.Request) string
}

// CookieLanguageParser returns the language value from cookie.
type CookieLanguageParser struct {
	name string
}

// NewCookieLanguageParser returns a CookieLanguageParser.
func NewCookieLanguageParser(name string) CookieLanguageParser {
	return CookieLanguageParser{name: name}
}

// Parse implements LanguageParser.Parse.
func (p CookieLanguageParser) Parse(r *http.Request) string {
	cookie, err := r.Cookie(p.name)
	if err != nil {
		return ""
	}

	return cookie.Value
}

// URLLanguageParser returns the language value from request URL.
type URLLanguageParser struct {
	name string
}

// NewURLLanguageParser returns a URLLanguageParser.
func NewURLLanguageParser(name string) URLLanguageParser {
	return URLLanguageParser{name: name}
}

// Parse implements LanguageParser.Parse.
func (p URLLanguageParser) Parse(r *http.Request) string {
	return r.URL.Query().Get(p.name)
}

// HeaderLanguageParser returns the language value from request header.
type HeaderLanguageParser struct {
}

// Parse implements LanguageParser.Parse.
func (HeaderLanguageParser) Parse(r *http.Request) string {
	return r.Header.Get("Accept-Language")
}
