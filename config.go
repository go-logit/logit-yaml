// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/11/29 02:26:34

package logityaml

import (
	"fmt"
	"github.com/go-logit/logit"
)

// loggerConfig stores all configurations of logger.
type loggerConfig struct {
	Level       string `json:"level" yaml:"level"`               // The level of a logger.
	NeedPid     bool   `json:"need_pid" yaml:"need_pid"`         // Logs will carry pid if needPid is true.
	NeedCaller  bool   `json:"need_caller" yaml:"need_caller"`   // Logs will carry caller information if needCaller is true.
	MsgKey      string `json:"msg_key" yaml:"msg_key"`           // The key of message in a log.
	TimeKey     string `json:"time_key" yaml:"time_key"`         // The key of time in a log.
	LevelKey    string `json:"level_key" yaml:"level_key"`       // The key of level in a log.
	PIDKey      string `json:"pid_key" yaml:"pid_key"`           // The key of pid in a log.
	FileKey     string `json:"file_key" yaml:"file_key"`         // The key of caller's file in a log.
	LineKey     string `json:"line_key" yaml:"line_key"`         // The key of caller's line in a log.
	TimeFormat  string `json:"time_format" yaml:"time_format"`   // The format of time in a log.
	CallerDepth int    `json:"caller_depth" yaml:"caller_depth"` // The depth of caller.
}

// config is the struct of yaml configuration.
type config struct {
	Logger loggerConfig `json:"logger" yaml:"logger"` // Wrap with logger so we can use the same config file.
}

// newDefaultConfig returns a default config.
func newDefaultConfig() *config {
	return &config{
		loggerConfig{
			Level:       "debug",
			NeedPid:     false,
			NeedCaller:  false,
			MsgKey:      "log.msg",
			TimeKey:     "log.time",
			LevelKey:    "log.level",
			PIDKey:      "log.pid",
			FileKey:     "log.file",
			LineKey:     "log.line",
			TimeFormat:  "2006-01-02 15:04:05",
			CallerDepth: 3,
		},
	}
}

// levelOption returns the level option of c.
func (c *config) levelOption() (logit.Option, error) {
	switch c.Logger.Level {
	case "debug":
		return logit.Options().WithDebugLevel(), nil
	case "info":
		return logit.Options().WithInfoLevel(), nil
	case "warn":
		return logit.Options().WithWarnLevel(), nil
	case "error":
		return logit.Options().WithErrorLevel(), nil
	case "off":
		return logit.Options().WithOffLevel(), nil
	default:
		return nil, fmt.Errorf("level %s mismatches", c.Logger.Level)
	}
}

// Options returns a slice of logit.Option for creating logit.Logger.
// Returns an error if something wrong happens.
func (c *config) Options() ([]logit.Option, error) {
	if c == nil {
		return nil, nil
	}

	levelOption, err := c.levelOption()
	if err != nil {
		return nil, err
	}

	options := logit.Options()
	result := []logit.Option{levelOption}

	cfg := c.Logger
	if cfg.NeedPid {
		result = append(result, options.WithPid())
	}

	if cfg.NeedCaller {
		result = append(result, options.WithCaller())
	}

	return result, nil
}
