# Internationalization
[![Build Status](https://travis-ci.org/clevergo/i18n.svg?branch=master)](https://travis-ci.org/clevergo/i18n)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/i18n/badge.svg?branch=master)](https://coveralls.io/github/clevergo/i18n?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/i18n)](https://goreportcard.com/report/github.com/clevergo/i18n)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/clevergo/i18n)
[![Sourcegraph](https://sourcegraph.com/github.com/clevergo/i18n/-/badge.svg)](https://sourcegraph.com/github.com/clevergo/i18n?badge)
[![Release](https://img.shields.io/github/release/clevergo/i18n.svg?style=flat-square)](https://github.com/clevergo/i18n/releases)

This package is built on top of [text/language](https://pkg.go.dev/golang.org/x/text/language) and [text/message](https://pkg.go.dev/golang.org/x/text/message).

## Usage

Please take a look of the following [example](example):

```shell
$ cd example
$ go run main.go
```

```
## fallback language(default to English)
$ curl "http://localhost:1234"
Home

## retrieve prefered language from URL query
$ curl "http://localhost:1234?lang=zh"
主页

$ curl "http://localhost:1234?lang=zh-TW"
主頁

$ curl "http://localhost:1234?lang=zh-HK"
主頁

## retrieve prefered language Cookie
$ curl -b "lang=zh-Hant" "http://localhost:1234"
主頁

## retrieve prefered language from header
$ curl -H "Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-US;q=0.7,zh-TW;q=0.6,pt;q=0.5" "http://localhost:1234/hello?name=foo"
你好，foo
```

## Integrate with other frameworks

It is easy to integrate with other frameworks by [Handler](https://pkg.go.dev/github.com/clevergo/i18n#Handler) or [Middleware](https://pkg.go.dev/github.com/clevergo/i18n#Middleware), and then retrieves translator in handler by [GetTranslator](https://pkg.go.dev/github.com/clevergo/i18n#GetTranslator).
