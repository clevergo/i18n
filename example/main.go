package main

import (
	"net/http"

	"github.com/clevergo/i18n"
)

var (
	translators *i18n.Translators
)

func index(w http.ResponseWriter, r *http.Request) {
	translator := i18n.GetTranslator(r)
	translator.Fprintf(w, "%m", "home")
}

func hello(w http.ResponseWriter, r *http.Request) {
	translator := i18n.GetTranslator(r)
	name := r.URL.Query().Get("name")
	translator.Fprintf(w, "hello %s", name)
}

func main() {
	translators = i18n.New(
	// i18n.Fallback("en"), // fallback language, default to English.
	)
	store := i18n.NewFileStore("./translations", i18n.JSONFileDecoder{})
	err := translators.Import(store)
	if err != nil {
		panic(err)
	}

	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.HandleFunc("/hello", hello)
	parsers := []i18n.LanguageParser{
		i18n.NewURLLanguageParser("lang"),    // from URL query
		i18n.NewCookieLanguageParser("lang"), // from cookie
		i18n.HeaderLanguageParser{},          // from Accept-Language header
	}
	handler := i18n.Handler(translators, mux, parsers...)
	http.ListenAndServe(":1234", handler)
}
