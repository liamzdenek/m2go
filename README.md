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

```go
package main

import "fmt";
import "../../Mongrel2";

func main() {
    conn := *m2go.NewM2Connection("82209006-86FF-4982-B5EA-D1E29E55D481", "tcp://127.0.0.1:9997", "tcp://127.0.0.1:9996");

    var req *m2go.Request;
    for {
        req = conn.Poll();
        if req != nil {
            response := m2go.Response{};
            response.Body = "Hello, World!";
            fmt.Printf("replying: %s\n", response.String());
            req.Reply(response.String());
        }
    }
}
```
