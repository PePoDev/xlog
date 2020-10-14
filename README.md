# Xlog

E[X]tend [Log]rus is Wrapper function for use logrus as builder design pattern and intergrate to elasticsearch in one line of code.

## Geting Started

```sh
go get github.com/pepodev/xlog
```

### Example

Just import and then use it

```go
package main

import (
    "github.com/pepodev/xlog"
)

func main(){
    xlog.DefaultLogFormatter()

    xlog.Trace("Trace")
    xlog.Debug("Debug")
    xlog.Info("Info")
    xlog.Warn("Warn")
    xlog.Error("Error")
    xlog.Panic("Panic")
    xlog.Fatal("Fatal")
}
```

Use with Elasticserch Hook

```go
package main

import (
    "github.com/pepodev/xlog"
)

func main(){
    xlog.DefaultLogFormatter()

    xlog.ConnectElasticsearch("http://localhost:9200", "username", "password", "host", "index", logrus.DebugLevel)

    xlog.Info("HI Elasticsearch from XLog <3")
}
```

Use builder design pattern

```go
package main

import (
    "github.com/pepodev/xlog"
)

func main(){
    xlog.DefaultLogFormatter()
    xlog.ConnectElasticsearch("http://localhost:9200", "username", "password", "host", "index", logrus.DebugLevel)

    log := xlog.NewLog()

    log.SetField("SomeKey", "SomeValue").Info("HI Elasticsearch from XLog <3")
}
```

## Depenencies

- [Logrus](https://github.com/sirupsen/logrus) (Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger.)

- [eLogrus](https://github.com/sohlich/elogrus) (ElasticSearch Hook for Logrus :walrus:)
