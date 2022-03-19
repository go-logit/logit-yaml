# 📝 logit-yaml

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/props)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![License](_icons/build.svg)](_icons/build.svg)
[![License](_icons/coverage.svg)](_icons/coverage.svg)

**logit-yaml** is a configuration module for logit which can use yaml to create a logit instance.

[阅读中文版的 Read me](./README.md)

### 🥇 Features

* Create logit.Logger from yaml configuration.

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

### 📖 Guides

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

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects logit-yaml used

| Project | Author   | Description                                           | link                                                                                    |
|---------|----------|-------------------------------------------------------|-----------------------------------------------------------------------------------------|
| logit   | go-logit | A high-performance and easy-to-use logging foundation | [Gitee](https://gitee.com/go-logit/logit) / [GitHub](https://github.com/go-logit/logit) |
| yaml    | go-yaml  | An easy-to-use yaml configuration foundation          | [GitHub](https://gopkg.in/yaml.v2)                                                      |
