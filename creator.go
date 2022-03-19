// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logityaml

import (
	"errors"
	"io/ioutil"

	"github.com/go-logit/logit"
	"gopkg.in/yaml.v2"
)

func init() {
	err := logit.RegisterLoggerCreator("yaml", newYamlCreator())
	if err != nil {
		panic(err)
	}
}

// yamlCreator creates logit.Logger from yaml configuration.
type yamlCreator struct{}

// newYamlCreator creates a yaml logger creator.
func newYamlCreator() *yamlCreator {
	return new(yamlCreator)
}

func (yc *yamlCreator) getConfig(params ...interface{}) (*config, error) {
	if len(params) < 1 {
		return nil, errors.New("logit-yaml: need the path of yaml configuration")
	}

	configPath, ok := params[0].(string)
	if !ok {
		return nil, errors.New("logit-yaml: params[0] must be the path of yaml configuration")
	}

	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	cfg := newDefaultConfig()
	return cfg, yaml.Unmarshal(configBytes, cfg)
}

func (yc *yamlCreator) getInterceptors(params ...interface{}) []logit.Interceptor {
	interceptors := make([]logit.Interceptor, 0, len(params)+2)

	for _, param := range params {
		if interceptor, ok := param.(logit.Interceptor); ok {
			interceptors = append(interceptors, interceptor)
		}
	}

	return interceptors
}

// CreateLogger creates a new logger using params and returns an error if failed.
func (yc *yamlCreator) CreateLogger(params ...interface{}) (*logit.Logger, error) {
	cfg, err := yc.getConfig(params...)
	if err != nil {
		return nil, err
	}

	options, err := cfg.ToOptions()
	if err != nil {
		return nil, err
	}

	// Design a Param struct to do this thing?
	interceptors := yc.getInterceptors(params...)
	if len(interceptors) > 0 {
		options = append(options, logit.Options().WithInterceptors(interceptors...))
	}

	return logit.NewLogger(options...), nil
}
