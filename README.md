# 📝 logit-yaml

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/props)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![License](_icons/build.svg)](_icons/build.svg)
[![License](_icons/coverage.svg)](_icons/coverage.svg)

**logit-yaml** 是一个配置文件模块，可以用来创建 logit 实例。

[Read me in English](./README.en.md)

### 🥇 功能特性

* 使用 yaml 配置创建 logit 的 Logger。

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

### 📖 使用手册

```shell
$ go get -u github.com/go-logit/logit-yaml
```

```go
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

	logger.Debug("I am created by yaml creator!").String("yaml", configPath).End()
	logger.Info("I am created by yaml creator!").String("yaml", configPath).End()
	logger.Warn("I am created by yaml creator!").String("yaml", configPath).End()
	logger.Error("I am created by yaml creator!").String("yaml", configPath).End()
	logger.Print("I am created by yaml creator!")
	logger.Println("I am created by yaml creator!")
	logger.Printf("I am created by yaml creator!")
}
```

### 👥 贡献者

如果您觉得 **logit-yaml** 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 logit-yaml 使用的技术

| 项目    | 作者       | 描述                  | 链接                                                                                   |
|-------|----------|---------------------|--------------------------------------------------------------------------------------|
| logit | go-logit | 一个高性能、功能强大且极易上手的日志库 | [码云](https://gitee.com/go-logit/logit) / [GitHub](https://github.com/go-logit/logit) |
| yaml  | go-yaml  | 一个极易上手的 yaml 配置库    | [GitHub](https://gopkg.in/yaml.v2)                                                   |
