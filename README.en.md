# 📝 logit-yaml

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/props)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![License](_icons/build.svg)](_icons/build.svg)
[![License](_icons/coverage.svg)](_icons/coverage.svg)

**logit-yaml** is a configuration module for logit which can use yaml to create a logit instance.

[阅读中文版的 Read me](./README.md)

### 🥇 Features

* coming soon...

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

### 📖 Guides

```shell
$ go get -u github.com/go-logit/logit-yaml
```

```go
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

	logger.Info("I am created by yaml maker!").String("yaml", configPath).End()
}
```

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects logit-yaml used

| Project | Author | Description | link |
| -----------|--------|-------------|------------------|
| logit | FishGoddess | A high-performance and easy-to-use logging foundation | [Gitee](https://gitee.com/go-logit/logit) / [GitHub](https://github.com/go-logit/logit) |
