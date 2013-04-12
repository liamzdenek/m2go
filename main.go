package main

import "fmt";

func main() {
    conn := *NewM2Connection("82209006-86FF-4982-B5EA-D1E29E55D481", "tcp://127.0.0.1:9997", "tcp://127.0.0.1:9996");

    fmt.Printf("Starting\n");
    for {
        req := conn.poll();
        if req != nil {
            //fmt.Printf("%v\n", req.headers[len(req.headers)-1].value);
            req.reply("HTTP/1.0 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 3\r\n\r\nlol");
        }
    }
}

