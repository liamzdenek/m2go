m2go - a library to develop Mongrel2 handlers in Golang
=======================================================

* Mongrel2: <http://mongrel2.org>
* m2php, a very similar library to m2go, and the inspiration to write m2go: <http://github.com/winks/m2php>

Requirements
------------

* ZeroMQ 2.1 or later: <http://www.zeromq.org/>
* Golang 1.0 or later: <http://golang.org/>
* ZeroMQ Golang bindings: <http://github.com/alecthomas/gozmq>

Usage
-----

See Examples/hello_world/main.go for the full + latest version. I'm not the most reliable at ensuring that the README is up to date.

```go
package main

import "fmt";
import "regexp";
import "bytes";
import "../../Mongrel2/";

func main() {
    r := m2go.Router{};
    r.AddRoute(m2go.Route{Path:regexp.MustCompile(`^/$`),Handler:SayHello});

    conn := *m2go.NewConnection(r,nil,"82209006-86FF-4982-B5EA-D1E29E55D481", "tcp://127.0.0.1:9997", "tcp://127.0.0.1:9996");
    conn.StartServer();
}

func SayHello(r *m2go.Request) {
    response := r.NewResponse();
    response.Body = "Hello, World!";
    r.Reply(response.String());
}
```
