// Copyright 2022 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2022/02/06 21:59:22

package main

import (
	"context"

	"github.com/go-logit/logit"
	_ "github.com/go-logit/logit-yaml" // Register yaml maker to logit.
)

func main() {
	configPath := "./config.yml"

	logger, err := logit.NewLoggerFromMaker(context.Background(), "yaml", configPath)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Debug("I am created by yaml maker!").String("yaml", configPath).End()
	logger.Info("I am created by yaml maker!").String("yaml", configPath).End()
	logger.Warn("I am created by yaml maker!").String("yaml", configPath).End()
	logger.Error("I am created by yaml maker!").String("yaml", configPath).End()
	logger.Print("I am created by yaml maker!")
	logger.Println("I am created by yaml maker!")
	logger.Printf("I am created by yaml maker!")
}
