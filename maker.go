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

	"github.com/FishGoddess/logit"
	"gopkg.in/yaml.v2"
)

func init() {
	err := logit.RegisterLoggerMaker("yaml", newYamlLoggerMaker())
	if err != nil {
		panic(err)
	}
}

// yamlLoggerMaker makes logit.Logger from yaml configuration.
type yamlLoggerMaker struct{}

// newYamlLoggerMaker 创建一个 yaml 日志配置初始化器
func newYamlLoggerMaker() logit.LoggerMaker {
	return new(yamlLoggerMaker)
}

// yamlToOptions reads yaml and transfers yaml to []logit.Option.
func (ylm *yamlLoggerMaker) yamlToOptions(yamlFile string) ([]logit.Option, error) {
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
func (ylm *yamlLoggerMaker) MakeLogger(ctx context.Context, params ...interface{}) (*logit.Logger, error) {
	if len(params) < 1 {
		return nil, errors.New("YamlLoggerMaker: need the path of yaml configuration")
	}

	yamlFile, ok := params[0].(string)
	if !ok {
		return nil, errors.New("YamlLoggerMaker: params[0] must be the path of yaml configuration")
	}

	options, err := ylm.yamlToOptions(yamlFile)
	if err != nil {
		return nil, err
	}

	return logit.NewLogger(options...), nil
}
