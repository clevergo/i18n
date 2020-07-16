// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package i18n

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/text/language"
)

func TestCookieLanguageParser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	name := "lang"
	parser := NewCookieLanguageParser(name)
	lang := parser.Parse(req)
	if lang != "" {
		t.Errorf("expected empty language, got %q", lang)
	}

	req.AddCookie(&http.Cookie{Name: name, Value: "en"})
	lang = parser.Parse(req)
	if lang != "en" {
		t.Errorf("expected language %q, got %q", "en", lang)
	}
}

func TestURLLanguageParser(t *testing.T) {
	name := "lang"
	parser := NewURLLanguageParser(name)
	tests := []string{"de", "en", "zh"}
	for _, test := range tests {
		req := httptest.NewRequest(http.MethodGet, "/?"+name+"="+test, nil)
		lang := parser.Parse(req)
		if lang != test {
			t.Errorf("expected language %q, got %q", test, lang)
		}
	}
}

func TestHeaderLanguageParser(t *testing.T) {
	tests := []string{"de", "en", "zh"}
	parser := HeaderLanguageParser{}
	for _, test := range tests {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Accept-Language", test)
		lang := parser.Parse(req)
		if lang != test {
			t.Errorf("expected language %q, got %q", test, lang)
		}
	}
}

func TestMiddleware(t *testing.T) {
	var translator *Translator
	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		translator = GetTranslator(r)
	})

	ts := New()
	ts.SetString(language.English, "hello", "Hello")
	ts.SetString(language.Chinese, "hello", "你好")

	handler = Middleware(ts, NewURLLanguageParser("lang"))(handler)

	tests := []struct {
		lang     string
		expected string
	}{
		{"", "Hello"},
		{"en", "Hello"},
		{"zh", "你好"},
	}

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodGet, "/?lang="+test.lang, nil)
		handler.ServeHTTP(nil, req)
		if translator == nil {
			t.Error("expected a translator, got nil")
		}
		if translator.Sprintf("%m", "hello") != test.expected {
			t.Errorf("expected %q, got %q", test.expected, translator.Sprintf("%m", "hello"))
		}

	}
}
