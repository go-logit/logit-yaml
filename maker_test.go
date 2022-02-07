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
		loggerConfig{
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
