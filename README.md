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

```golang
pakage main;

import "fmt";

func main() {
    conn := *NewM2Connection("82209006-86FF-4982-B4EA-D1E29E55D481", "tcp://127.0.0.1:9997", "tcp://127.0.0.1:9996");
    
    var req *Request;
    for {
        req = conn.poll();
        if req != nil {
            response := Response{};
            response.body = "Hello, World";
            req.reply(response.String());
        }
    }
}
```
