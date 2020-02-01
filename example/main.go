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
	translators = i18n.New()
	store := i18n.NewFileStore("./translations", i18n.JSONFileDecoder{})
	err := translators.Import(store)
	if err != nil {
		panic(err)
	}

	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.HandleFunc("/hello", hello)
	parsers := []i18n.LanguageParser{i18n.NewURLLanguageParser("lang"), i18n.NewCookieLanguageParser("lang"), i18n.HeaderLanguageParser{}}
	handler := i18n.Handler(translators, mux, parsers...)
	http.ListenAndServe(":1234", handler)
}
