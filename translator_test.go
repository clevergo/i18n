// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package i18n

import "testing"

func TestFallback(t *testing.T) {
	tests := []string{"de", "en", "zh"}
	for _, test := range tests {
		ts := New(Fallback(test))
		if ts.fallback != test {
			t.Errorf("expected fallback %q, got %q", test, ts.fallback)
		}
	}
}

func TestTranslators_Import(t *testing.T) {
	ts := New()
	fs := NewFileStore(tempDir, JSONFileDecoder{})
	err := ts.Import(fs)
	if err != nil {
		t.Errorf("failed to import translations: %s", err)
	}
}
