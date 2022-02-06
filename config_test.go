// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/11/29 02:35:07

package logityaml

import "testing"

// go test -v -cover -run =^TestConfigToOptions$
func TestConfigToOptions(t *testing.T) {
	c := newDefaultConfig()

	options, err := c.ToOptions()
	if err != nil {
		t.Error(err)
	}

	for _, _ = range options {

	}
}
