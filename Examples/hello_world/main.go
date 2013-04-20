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

