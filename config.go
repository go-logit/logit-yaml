// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/11/29 02:26:34

package logityaml

import "github.com/FishGoddess/logit"

// config is the struct of yaml configuration.
type config struct {
	// TODO Complete yaml configuration
	Level string `yaml:"level"`
}

// newConfig returns a new config holder.
func newConfig() *config {
	return new(config)
}

// toLogitOptions returns a slice of logit.Option for creating logit.Logger.
// Returns an error if something wrong happens.
func (c *config) toLogitOptions() ([]logit.Option, error) {
	return nil, nil
}
