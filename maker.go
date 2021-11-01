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

	"github.com/FishGoddess/logit"
)

// YamlLoggerMaker makes logit.Logger from yaml configuration.
type YamlLoggerMaker struct{}

// yamlToOptions reads yaml and transfers yaml to []logit.Option.
func (ylm *YamlLoggerMaker) yamlToOptions(yamlFile string) ([]logit.Option, error) {
	// TODO Read yaml and transfer yaml to []logit.Option
	return []logit.Option{}, nil
}

// MakeLogger makes a new logger using params and returns an error if something wrong happens.
func (ylm *YamlLoggerMaker) MakeLogger(ctx context.Context, params interface{}) (*logit.Logger, error) {
	yamlFile := params.(string)

	options, err := ylm.yamlToOptions(yamlFile)
	if err != nil {
		return nil, err
	}

	return logit.NewLogger(options...), nil
}
