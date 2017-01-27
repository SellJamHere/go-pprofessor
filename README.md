# go-pprofessor

A simple package to run [net/http/pprof](https://golang.org/pkg/net/http/pprof/).
https://blog.golang.org/profiling-go-programs

This package runs an http server at the specified port, automatically adding the appropriate net/http/pprof routes.

## Usage:

```
pprofessor.Serve(":6060")
```

Then visit `http://localhost:6060/debug/pprof/`
or run `go tool pprof http://localhost:6060/debug/pprof/profile` (try `top10` at the prompt)
