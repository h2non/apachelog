# go-apache-log [![Build Status](https://travis-ci.org/h2non/apachelog.png)](https://travis-ci.org/h2non/apachelog) [![GoDoc](https://godoc.org/github.com/h2non/apachelog?status.svg)](https://godoc.org/github.com/h2non/apachelog)

[Go](https://golang.org) `net/http` compatible middleware for [Apache](http://httpd.apache.org/docs/2.2/logs.html) style logging.
Simple. Small. Dependency free.

Originally taken from [`imaginary`](https://github.com/h2non/imaginary) package, now isolated as standalone package for better reusability.

## Installation

```bash
go get gopkg.in/h2non/apachelog.v0
```

## Usage

```go
package main

import (
  "fmt"
  "gopkg.in/h2non/apachelog.v0"
  "net/http"
  "os"
)

func main() {
  mux := http.NewServeMux()

  // Sample route
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World"))
  })

  handler := apachelog.New(mux, os.Stdout)

  fmt.Println("Server listening on port: 3000")
  http.ListenAndServe(":3000", handler)
}
```

## API

Please, see [godoc reference](https://godoc.org/github.com/h2non/apachelog).

## License

MIT - Tomas Aparicio