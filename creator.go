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

// parseOptions reads yaml and transfers yaml to []logit.Option.
func (ym *yamlCreator) parseOptions(yamlFile string) ([]logit.Option, error) {
	fileBytes, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}

	cfg := newDefaultConfig()
	err = yaml.Unmarshal(fileBytes, cfg)
	if err != nil {
		return nil, err
	}

	return cfg.ToOptions()
}

func (ym *yamlCreator) getConfig(params ...interface{}) (*config, error) {
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

// CreateLogger creates a new logger using params and returns an error if failed.
func (ym *yamlCreator) CreateLogger(params ...interface{}) (*logit.Logger, error) {
	cfg, err := ym.getConfig(params...)
	if err != nil {
		return nil, err
	}

	options, err := cfg.ToOptions()
	if err != nil {
		return nil, err
	}

	return logit.NewLogger(options...), nil
}
