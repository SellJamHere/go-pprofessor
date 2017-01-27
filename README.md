# go-pprofessor

A simple package to run [net/http/pprof](https://golang.org/pkg/net/http/pprof/).

This package runs an http server at the specified port, automatically adding the appropriate net/http/pprof routes.

## Usage:

```
pprofessor.Serve(":6060")
```

Then visit `http://localhost:6060/debug/pprof/`
