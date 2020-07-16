// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package i18n

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"testing"
)

var (
	tempDir string
)

func TestMain(m *testing.M) {
	var err error
	tempDir, err = ioutil.TempDir(os.TempDir(), "translations")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	enTempDir := path.Join(tempDir, "en")
	err = os.Mkdir(enTempDir, 0744)
	if err != nil {
		log.Fatal(err)
	}
	file, err := ioutil.TempFile(enTempDir, "test.json")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write([]byte(`{"hello": "Hello"}`))
	if err != nil {
		log.Fatal(err)
	}

	zhTempDir := path.Join(tempDir, "zh")
	err = os.Mkdir(zhTempDir, 0744)
	if err != nil {
		log.Fatal(err)
	}
	file, err = ioutil.TempFile(zhTempDir, "test.json")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write([]byte(`{"hello": "你好"}`))
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestFileStore(t *testing.T) {
	tests := []struct {
		dir          string
		shouldError  bool
		translations Translations
	}{
		{"", true, nil},
		{tempDir, false, Translations{
			"en": {
				"hello": "Hello",
			},
			"zh": {
				"hello": "你好",
			},
		}},
	}

	for _, test := range tests {
		fs := NewFileStore(test.dir, JSONFileDecoder{})
		translations, err := fs.Get()
		if test.shouldError {
			if err == nil {
				t.Error("expected an error, got nil")
			}
			continue
		}
		if !reflect.DeepEqual(test.translations, translations) {
			t.Errorf("expected translations %+v, got %+v", test.translations, translations)
		}
	}

}

func TestJSONDecoder(t *testing.T) {
}
