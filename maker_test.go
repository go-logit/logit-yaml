// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/11/02 00:22:27

package logityaml

import (
	"io/ioutil"
	"testing"
)

var (
	testConfig = &config{
		struct {
			Level         string `json:"level" yaml:"level"`
			NeedPid       bool   `json:"need_pid" yaml:"need_pid"`
			NeedCaller    bool   `json:"need_caller" yaml:"need_caller"`
			MsgKey        string `json:"msg_key" yaml:"msg_key"`
			TimeKey       string `json:"time_key" yaml:"time_key"`
			LevelKey      string `json:"level_key" yaml:"level_key"`
			PidKey        string `json:"pid_key" yaml:"pid_key"`
			FileKey       string `json:"file_key" yaml:"file_key"`
			LineKey       string `json:"line_key" yaml:"line_key"`
			TimeFormat    string `json:"time_format" yaml:"time_format"`
			CallerDepth   int    `json:"caller_depth" yaml:"caller_depth"`
			AutoFlush     string `json:"auto_flush" yaml:"auto_flush"`
			Appender      string `json:"appender" yaml:"appender"`
			DebugAppender string `json:"debug_appender" yaml:"debug_appender"`
			InfoAppender  string `json:"info_appender" yaml:"info_appender"`
			WarnAppender  string `json:"warn_appender" yaml:"warn_appender"`
			ErrorAppender string `json:"error_appender" yaml:"error_appender"`
			PrintAppender string `json:"print_appender" yaml:"print_appender"`
			Writer        struct {
				Output   string `json:"output" yaml:"output"`
				Buffered bool   `json:"buffered" yaml:"buffered"`
			} `json:"writer" yaml:"writer"`
		}{
			Level:       "info",
			NeedPid:     true,
			NeedCaller:  true,
			MsgKey:      "msg",
			TimeKey:     "time",
			LevelKey:    "level",
			PidKey:      "pid",
			FileKey:     "file",
			LineKey:     "line",
			TimeFormat:  "20060102150405",
			CallerDepth: 1,
		},
	}

	testConfigBytes = []byte(`
logger:
  level: "info"
  need_pid: true
  need_caller: true
  msg_key: "msg"
  time_key: "time"
  level_key: "level"
  pid_key: "pid"
  file_key: "file"
  line_key: "line"
  time_format: "20060102150405"
  caller_depth: 1
`)
)

func createConfig(dir string, pattern string) (string, error) {
	file, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(testConfigBytes)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

// go test -v -cover -run=^TestYamlMakerGetConfig$
func TestYamlMakerGetConfig(t *testing.T) {
	maker := newYamlMaker()

	_, err := maker.getConfig()
	if err == nil {
		t.Error("err == nil")
	}

	configPath, err := createConfig(t.TempDir(), t.Name())
	if err != nil {
		t.Error(err)
	}

	t.Log(configPath)
	cfg, err := maker.getConfig(configPath)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", cfg)

	if *cfg != *testConfig {
		t.Errorf("cfg %+v != testConfig %+v", cfg, testConfig)
	}
}
