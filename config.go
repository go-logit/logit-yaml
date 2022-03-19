// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logityaml

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/go-logit/logit"
	"github.com/go-logit/logit/core/appender"
	"github.com/go-logit/logit/pkg"
)

// config is the struct of yaml configuration.
type config struct {
	Logger struct {
		Level         string `json:"level" yaml:"level"`                   // The level of a logger.
		NeedPid       bool   `json:"need_pid" yaml:"need_pid"`             // Logs will carry pid if needPid is true.
		NeedCaller    bool   `json:"need_caller" yaml:"need_caller"`       // Logs will carry caller information if needCaller is true.
		MsgKey        string `json:"msg_key" yaml:"msg_key"`               // The key of message in a log.
		TimeKey       string `json:"time_key" yaml:"time_key"`             // The key of time in a log.
		LevelKey      string `json:"level_key" yaml:"level_key"`           // The key of level in a log.
		PidKey        string `json:"pid_key" yaml:"pid_key"`               // The key of pid in a log.
		FileKey       string `json:"file_key" yaml:"file_key"`             // The key of caller's file in a log.
		LineKey       string `json:"line_key" yaml:"line_key"`             // The key of caller's line in a log.
		TimeFormat    string `json:"time_format" yaml:"time_format"`       // The format of time in a log.
		CallerDepth   int    `json:"caller_depth" yaml:"caller_depth"`     // The depth of caller.
		AutoFlush     string `json:"auto_flush" yaml:"auto_flush"`         // The frequency between two flush operations. See time.ParseDuration.
		Appender      string `json:"appender" yaml:"appender"`             // The appender of logger.
		DebugAppender string `json:"debug_appender" yaml:"debug_appender"` // The debug appender of logger.
		InfoAppender  string `json:"info_appender" yaml:"info_appender"`   // The info appender of logger.
		WarnAppender  string `json:"warn_appender" yaml:"warn_appender"`   // The warning appender of logger.
		ErrorAppender string `json:"error_appender" yaml:"error_appender"` // The error appender of logger.
		PrintAppender string `json:"print_appender" yaml:"print_appender"` // The print appender of logger.
		Writer        struct {
			Output   string `json:"output" yaml:"output"`     // The output of writer. You can set stdout or stderr or a file path.
			Buffered bool   `json:"buffered" yaml:"buffered"` // The output mode of writer. Use buffered writer if true.
		} `json:"writer" yaml:"writer"` // The writer of logger.
		DebugWriter struct {
			Output   string `json:"output" yaml:"output"`     // The output of writer. You can set stdout or stderr or a file path.
			Buffered bool   `json:"buffered" yaml:"buffered"` // The output mode of writer. Use buffered writer if true.
		} `json:"debug_writer" yaml:"debug_writer"` // The debug writer of logger.
		InfoWriter struct {
			Output   string `json:"output" yaml:"output"`     // The output of writer. You can set stdout or stderr or a file path.
			Buffered bool   `json:"buffered" yaml:"buffered"` // The output mode of writer. Use buffered writer if true.
		} `json:"info_writer" yaml:"info_writer"` // The info writer of logger.
		WarnWriter struct {
			Output   string `json:"output" yaml:"output"`     // The output of writer. You can set stdout or stderr or a file path.
			Buffered bool   `json:"buffered" yaml:"buffered"` // The output mode of writer. Use buffered writer if true.
		} `json:"warn_writer" yaml:"warn_writer"` // The warning writer of logger.
		ErrorWriter struct {
			Output   string `json:"output" yaml:"output"`     // The output of writer. You can set stdout or stderr or a file path.
			Buffered bool   `json:"buffered" yaml:"buffered"` // The output mode of writer. Use buffered writer if true.
		} `json:"error_writer" yaml:"error_writer"` // The error writer of logger.
		PrintWriter struct {
			Output   string `json:"output" yaml:"output"`     // The output of writer. You can set stdout or stderr or a file path.
			Buffered bool   `json:"buffered" yaml:"buffered"` // The output mode of writer. Use buffered writer if true.
		} `json:"print_writer" yaml:"print_writer"` // The print writer of logger.
	} `json:"logger" yaml:"logger"` // Wrap with logger, so we can use the same config file.
}

// newDefaultConfig returns a default config.
func newDefaultConfig() *config {
	return new(config)
}

// parseLevel returns the level option of c.
func (c *config) parseLevel(level string) (logit.Option, error) {
	switch strings.ToLower(level) {
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
		return nil, fmt.Errorf("logit-yaml: level %s mismatches", level)
	}
}

// parseAppender returns the appender of appenderName.
func (c *config) parseAppender(appenderName string) (appender.Appender, error) {
	switch strings.ToLower(appenderName) {
	case "text":
		return appender.Text(), nil
	case "json":
		return appender.Json(), nil
	default:
		return nil, fmt.Errorf("logit-yaml: appender %s mismatches", appenderName)
	}
}

// parseWriter returns the writer of writerName.
func (c *config) parseWriter(writerName string) (io.Writer, error) {
	switch strings.ToLower(writerName) {
	case "stdout":
		return os.Stdout, nil
	case "stderr":
		return os.Stderr, nil
	default:
		return pkg.NewFile(writerName)
	}
}

// ToOptions returns a slice of logit.Option for creating logit.Logger.
// Returns an error if something wrong happens.
func (c *config) ToOptions() ([]logit.Option, error) {
	if c == nil {
		return nil, nil
	}

	options := logit.Options()
	result := make([]logit.Option, 0, 16)

	cfg := c.Logger
	if cfg.Level != "" {
		levelOption, err := c.parseLevel(cfg.Level)
		if err != nil {
			return nil, err
		}

		result = append(result, levelOption)
	}

	if cfg.NeedPid {
		result = append(result, options.WithPid())
	}

	if cfg.NeedCaller {
		result = append(result, options.WithCaller())
	}

	if cfg.MsgKey != "" {
		result = append(result, options.WithMsgKey(cfg.MsgKey))
	}

	if cfg.TimeKey != "" {
		result = append(result, options.WithTimeKey(cfg.TimeKey))
	}

	if cfg.LevelKey != "" {
		result = append(result, options.WithLevelKey(cfg.LevelKey))
	}

	if cfg.PidKey != "" {
		result = append(result, options.WithPidKey(cfg.PidKey))
	}

	if cfg.FileKey != "" {
		result = append(result, options.WithFileKey(cfg.FileKey))
	}

	if cfg.LineKey != "" {
		result = append(result, options.WithLineKey(cfg.LineKey))
	}

	result = append(result, options.WithTimeFormat(cfg.TimeFormat))
	if cfg.CallerDepth > 0 {
		result = append(result, options.WithCallerDepth(cfg.CallerDepth))
	}

	if cfg.AutoFlush != "" {
		frequency, err := time.ParseDuration(cfg.AutoFlush)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithAutoFlush(frequency))
	}

	if cfg.Appender != "" {
		apdr, err := c.parseAppender(cfg.Appender)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithAppender(apdr))
	}

	if cfg.DebugAppender != "" {
		apdr, err := c.parseAppender(cfg.DebugAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithDebugAppender(apdr))
	}

	if cfg.InfoAppender != "" {
		apdr, err := c.parseAppender(cfg.InfoAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithInfoAppender(apdr))
	}

	if cfg.WarnAppender != "" {
		apdr, err := c.parseAppender(cfg.WarnAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithWarnAppender(apdr))
	}

	if cfg.ErrorAppender != "" {
		apdr, err := c.parseAppender(cfg.ErrorAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithErrorAppender(apdr))
	}

	if cfg.PrintAppender != "" {
		apdr, err := c.parseAppender(cfg.PrintAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithPrintAppender(apdr))
	}

	if cfg.Writer.Output != "" {
		writer, err := c.parseWriter(cfg.Writer.Output)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithWriter(writer, cfg.Writer.Buffered))
	}

	if cfg.DebugWriter.Output != "" {
		writer, err := c.parseWriter(cfg.DebugWriter.Output)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithDebugWriter(writer, cfg.DebugWriter.Buffered))
	}

	if cfg.InfoWriter.Output != "" {
		writer, err := c.parseWriter(cfg.InfoWriter.Output)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithInfoWriter(writer, cfg.InfoWriter.Buffered))
	}

	if cfg.WarnWriter.Output != "" {
		writer, err := c.parseWriter(cfg.WarnWriter.Output)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithWarnWriter(writer, cfg.WarnWriter.Buffered))
	}

	if cfg.ErrorWriter.Output != "" {
		writer, err := c.parseWriter(cfg.ErrorWriter.Output)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithErrorWriter(writer, cfg.ErrorWriter.Buffered))
	}

	if cfg.PrintWriter.Output != "" {
		writer, err := c.parseWriter(cfg.PrintWriter.Output)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithPrintWriter(writer, cfg.PrintWriter.Buffered))
	}

	return result, nil
}
