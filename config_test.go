// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logityaml

import "testing"

// go test -v -cover -run =^TestConfigToOptions$
func TestConfigToOptions(t *testing.T) {
	c := newDefaultConfig()

	options, err := c.ToOptions()
	if err != nil {
		t.Error(err)
	}

	for range options {

	}
}
