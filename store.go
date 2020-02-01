// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package i18n

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

// Store is where translations located.
type Store interface {
	Get() (Translations, error)
}

// FileStore is a file store.
type FileStore struct {
	Directory string
	Decoder   FileDecoder
}

// NewFileStore returns a file store with the given directory and decoder.
func NewFileStore(directory string, decoder FileDecoder) *FileStore {
	return &FileStore{
		Directory: directory,
		Decoder:   decoder,
	}
}

// Get implements Store.Get.
func (s *FileStore) Get() (Translations, error) {
	translations := make(Translations)
	dirs, err := ioutil.ReadDir(s.Directory)
	if err != nil {
		return translations, err
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		files, err := ioutil.ReadDir(path.Join(s.Directory, dir.Name()))
		if err != nil {
			return translations, err
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			f, err := os.Open(path.Join(s.Directory, dir.Name(), file.Name()))
			if err != nil {
				return translations, err
			}
			defer f.Close()

			data, err := ioutil.ReadAll(f)
			if err != nil {
				return translations, err
			}
			var v Translation
			if err = s.Decoder.Decode(data, &v); err != nil {
				return translations, err
			}
			translations[dir.Name()] = v
		}
	}

	return translations, nil
}

// FileDecoder is a function that decode file content to tranlation.
type FileDecoder interface {
	Decode([]byte, *Translation) error
}

// JSONFileDecoder is a JSON file decoder.
type JSONFileDecoder struct {
}

// Decode implements FileDecoder.Decode
func (d JSONFileDecoder) Decode(data []byte, translation *Translation) error {
	return json.Unmarshal(data, translation)
}
