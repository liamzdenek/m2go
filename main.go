package main

import "fmt"
import "reflect"
import "strings"
import zmq "github.com/alecthomas/gozmq"
import "strconv"
import "regexp"
//import "github.com/bitly/go-simplejson"

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

type M2Connection struct
{
    ctx *zmq.Context;
    req *zmq.Socket;
    rsp *zmq.Socket;
    
    sender_id string;
}

func NewM2Connection(sender_id string, req_addr string, rsp_addr string) *M2Connection {
    ctx, _ := zmq.NewContext();
    req, _ := ctx.NewSocket(zmq.PULL);
    rsp, _ := ctx.NewSocket(zmq.PUB);
    req.Connect(req_addr);
    rsp.Connect(rsp_addr);
    req.SetSockOptInt(zmq.RCVTIMEO, 1000);
    rsp.SetSockOptString(zmq.IDENTITY, sender_id);
    
    return &M2Connection{
        ctx:ctx,
        req:req,
        rsp:rsp,
        sender_id:sender_id,
    };
}

func (conn *M2Connection) poll() (*Request) {
    msg, err := conn.req.Recv(0);
    if err == nil {
        parsed := conn.parse(string(msg));
        return parsed;
    }
    return nil;
}

func (conn *M2Connection) parse(msg string) *Request {
    //sender, conn_id, path, rest
    splitdata := strings.SplitN(msg, " ", 4);
    headers, rest := conn.parse_netstring(splitdata[3]);
    body, _ := conn.parse_netstring(rest);
    
    headers = headers[1:len(headers)-1];
    regex, err := regexp.Compile(`"(.*)"\:"(.*)"`);
    
    if err != nil {
        fmt.Printf("Regexp Error: %s", err);
    }

    headerary := make([]Header,1,1);
    
    var count = 0;
    for {
        splitdata := strings.SplitN(headers,",", 2);
        if len(splitdata) == 1 {
            break;
        }
        
        headers = string(splitdata[1]);
        parts := regex.FindStringSubmatch(splitdata[0]);
        fmt.Printf("data: %s parts: %s\n", splitdata[0], len(parts));
        if len(parts) == 3 {
            if len(headerary) >= cap(headerary) {
                newary := make([]Header, cap(headerary)+1);
                copy(newary, headerary);
                headerary = newary;
            }
            headerary[count] = Header{key:string(parts[0]),value:string(parts[1])};
            count++;
        }
    }
     
    return &Request{
        sender:  splitdata[0],
        conn_id: splitdata[1],
        path:    splitdata[2],
        body:    body,
        headers: headerary,
    };
}

func (con *M2Connection) parse_netstring(ns string) (string, string) {
    // length, rest
    splitdata := strings.SplitN(ns, ":", 2);
    datalen,_ := strconv.Atoi(splitdata[0]);
    return splitdata[1][:datalen], splitdata[1][datalen+1:] 
}

type Request struct {
    sender string;
    conn_id string;
    path string;
    body string;
    headers []Header; 
}

type Header struct {
    key string;
    value string;
}
