# ð logit-yaml

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/props)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![License](_icons/build.svg)](_icons/build.svg)
[![License](_icons/coverage.svg)](_icons/coverage.svg)

**logit-yaml** æ¯ä¸ä¸ªéç½®æä»¶æ¨¡åï¼å¯ä»¥ç¨æ¥åå»º logit å®ä¾ã

[Read me in English](./README.en.md)

### ð¥ åè½ç¹æ§

* ä½¿ç¨ yaml éç½®åå»º logit ç Loggerã

_åå²çæ¬çç¹æ§è¯·æ¥ç [HISTORY.md](./HISTORY.md)ãæªæ¥çæ¬çæ°ç¹æ§åè®¡åè¯·æ¥ç [FUTURE.md](./FUTURE.md)ã_

### ð ä½¿ç¨æå

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

### ð¥ è´¡ç®è

å¦ææ¨è§å¾ **logit-yaml** ç¼ºå°æ¨éè¦çåè½ï¼è¯·ä¸è¦ç¹è±«ï¼é©¬ä¸åä¸è¿æ¥ï¼åèµ·ä¸ä¸ª _**issue**_ã

### ð¦ logit-yaml ä½¿ç¨çææ¯

| é¡¹ç®    | ä½è       | æè¿°                  | é¾æ¥                                                                                   |
|-------|----------|---------------------|--------------------------------------------------------------------------------------|
| logit | go-logit | ä¸ä¸ªé«æ§è½ãåè½å¼ºå¤§ä¸ææä¸æçæ¥å¿åº | [ç äº](https://gitee.com/go-logit/logit) / [GitHub](https://github.com/go-logit/logit) |
| yaml  | go-yaml  | ä¸ä¸ªææä¸æç yaml éç½®åº    | [GitHub](https://gopkg.in/yaml.v2)                                                   |
