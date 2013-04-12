package main

import "fmt";
import "reflect";

func main() {
    conn := *NewM2Connection("82209006-86FF-4982-B5EA-D1E29E55D481", "tcp://127.0.0.1:9997", "tcp://127.0.0.1:9996");
    
    fmt.Printf("%s\n", reflect.TypeOf(conn));
    for {
        ret := conn.poll();
        if ret != nil {
            fmt.Printf("%v\n", ret.headers);
        }
    }
}

