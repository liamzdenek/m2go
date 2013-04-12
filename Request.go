package main;

import "fmt";

type Request struct {
    sender_id string;
    conn_id string;
    path string;
    body string;
    conn *M2Connection;
    headers []Header; 
}

func (req *Request) reply_http(body string) {
    
}

func (req *Request) reply(msg string) {
    conn := req.conn;
    rsp  := conn.rsp;
    response := fmt.Sprintf(
        "%s %d:%s, %s",
        req.sender_id,
        len(req.conn_id),
        req.conn_id,
        msg,
    );
    //fmt.Printf("Response: %s\n", response);
    rsp.Send([]byte(response),0);
}
