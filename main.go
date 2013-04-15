package main

import "fmt";

func main() {
    conn := *NewM2Connection("82209006-86FF-4982-B5EA-D1E29E55D481", "tcp://127.0.0.1:9997", "tcp://127.0.0.1:9996");

    fmt.Printf("Starting\n");
    var req *Request;
    for {
        req = conn.poll();
        if req != nil {
            response := Response{};
            response.body = "Hello, World!";
            fmt.Printf("replying: %s\n", response.String());
            req.reply(response.String());
        }
    }
}

