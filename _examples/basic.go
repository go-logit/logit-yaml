// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/go-logit/logit"
	_ "github.com/go-logit/logit-yaml" // Register yaml creator to logit.
)

func main() {
	configPath := "./config.yml"

	logger, err := logit.NewLoggerFromCreator("yaml", configPath)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Info("I am created by yaml creator!").String("yaml", configPath).End()
}
