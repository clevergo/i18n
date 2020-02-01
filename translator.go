// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package i18n

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

// Option for appling on translatiors.
type Option func(ts *Translators)

// Fallback is an option to change fallback language of translators.
func Fallback(fallback string) Option {
	return func(ts *Translators) {
		ts.fallback = fallback
	}
}

// Translators is a collections of translator.
type Translators struct {
	*catalog.Builder
	fallback string
}

// New returns a translators.
func New(opts ...Option) *Translators {
	ts := &Translators{
		fallback: "en",
	}
	for _, opt := range opts {
		opt(ts)
	}
	ts.Builder = catalog.NewBuilder(catalog.Fallback(language.Make(ts.fallback)))
	return ts
}

// MatchTranslator returns the matched translator of the given language.
func (ts *Translators) MatchTranslator(langs ...string) *Translator {
	tag, _ := language.MatchStrings(ts.Matcher(), langs...)
	p := message.NewPrinter(tag, message.Catalog(ts))
	return NewTranslator(p)
}

// Translations is a map that mapping from language to translations.
type Translations map[string]Translation

// Translation is a key-value pair.
type Translation map[string]string

// Import imports translations from the given store.
func (ts *Translators) Import(store Store) error {
	translations, err := store.Get()
	if err != nil {
		return err
	}

	for lang, translation := range translations {
		tag := language.Make(lang)
		if tag.IsRoot() {
			return fmt.Errorf("unsupported language code: %s", lang)
		}
		for key, msg := range translation {
			if err = ts.SetString(tag, key, msg); err != nil {
				return err
			}
		}
	}

	return nil
}

// Translator is a wrapper of message.Printer.
type Translator struct {
	*message.Printer
}

// NewTranslator returns a new Translator.
func NewTranslator(printer *message.Printer) *Translator {
	return &Translator{printer}
}
