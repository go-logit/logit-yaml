// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/11/02 00:10:11

package logityaml

import (
	"context"
	"errors"
	"io/ioutil"

	"github.com/go-logit/logit"
	"gopkg.in/yaml.v2"
)

func init() {
	err := logit.RegisterLoggerMaker("yaml", newYamlMaker())
	if err != nil {
		panic(err)
	}
}

// yamlMaker makes logit.Logger from yaml configuration.
type yamlMaker struct{}

// newYamlMaker creates a yaml logger maker.
func newYamlMaker() logit.LoggerMaker {
	return new(yamlMaker)
}

// parseOptionsFrom reads yaml and transfers yaml to []logit.Option.
func (ylm *yamlMaker) parseOptionsFrom(yamlFile string) ([]logit.Option, error) {
	fileBytes, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}

	cfg := newConfig()
	err = yaml.Unmarshal(fileBytes, cfg)
	if err != nil {
		return nil, err
	}

	return cfg.toLogitOptions()
}

// MakeLogger makes a new logger using params and returns an error if something wrong happens.
func (ylm *yamlMaker) MakeLogger(ctx context.Context, params ...interface{}) (*logit.Logger, error) {
	if len(params) < 1 {
		return nil, errors.New("YamlLoggerMaker: need the path of yaml configuration")
	}

	yamlFile, ok := params[0].(string)
	if !ok {
		return nil, errors.New("YamlLoggerMaker: params[0] must be the path of yaml configuration")
	}

	options, err := ylm.parseOptionsFrom(yamlFile)
	if err != nil {
		return nil, err
	}

	return logit.NewLogger(options...), nil
}
