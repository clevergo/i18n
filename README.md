# Internationalization
[![Build Status](https://img.shields.io/travis/clevergo/i18n?style=for-the-badge)](https://travis-ci.org/clevergo/i18n)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/i18n?style=for-the-badge)](https://coveralls.io/github/clevergo/i18n)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/i18n?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/i18n?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/i18n)
[![Release](https://img.shields.io/github/release/clevergo/i18n.svg?style=for-the-badge)](https://github.com/clevergo/i18n/releases)

This package is built on top of [text/language](https://pkg.go.dev/golang.org/x/text/language) and [text/message](https://pkg.go.dev/golang.org/x/text/message).

## Usage

Checkout [example](https://github.com/clevergo/examples/tree/master/i18n) for details.

## Integrate with other frameworks

It is easy to integrate with other frameworks by [Handler](https://pkg.go.dev/github.com/clevergo/i18n#Handler) or [Middleware](https://pkg.go.dev/github.com/clevergo/i18n#Middleware), and then retrieves translator in handler by [GetTranslator](https://pkg.go.dev/github.com/clevergo/i18n#GetTranslator).
