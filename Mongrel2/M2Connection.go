package m2go

import "strings";
import zmq "github.com/alecthomas/gozmq";
import "strconv";
import "regexp";

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
    //req.SetSockOptInt(zmq.RCVTIMEO, 1000);
    rsp.SetSockOptString(zmq.IDENTITY, sender_id);

    return &M2Connection{
        ctx:ctx,
        req:req,
        rsp:rsp,
        sender_id:sender_id,
    };
}

func (conn *M2Connection) Poll() (*Request) {
    msg, err := conn.req.Recv(0);
    if err == nil {
        parsed := conn.Parse(string(msg));
        return parsed;
    }
    return nil;
}

func (conn *M2Connection) Parse(msg string) *Request {
    //sender, conn_id, path, rest
    splitdata := strings.SplitN(msg, " ", 4);
    headers, rest := conn.parse_netstring(splitdata[3]);
    body, _ := conn.parse_netstring(rest);

    // we're not going to break out a JSON parser. DIY. For speed.
    headers = headers[1:len(headers)-1];
    regex,_ := regexp.Compile(`"(.*)"\:"(.*)"`);

    // precalculate header array size. this might accidentally
    // overshoot in size, but whatever.
    headerary := make([]Header,strings.Count(headers,"\":\""));

    var parts, headstring []string;
    for {
        headstring = strings.SplitN(headers,",", 2);
        if len(headstring) == 1 {
            break;
        }

        headers = headstring[1];
        // this could be replaced with sscanf. it should perform faster
        parts = regex.FindStringSubmatch(headstring[0]);
        if len(parts) == 3 {
            headerary = append(headerary, Header{key:string(parts[1]),value:string(parts[2])} );
        }
    }

    return &Request{
        sender_id: splitdata[0],
        conn_id:   splitdata[1],
        path:      splitdata[2],
        body:      body,
        conn:      conn,
        headers:   headerary,
    };
}

func (con *M2Connection) parse_netstring(ns string) (string, string) {
    // length, rest
    splitdata := strings.SplitN(ns, ":", 2);
    datalen,_ := strconv.Atoi(splitdata[0]);
    return splitdata[1][:datalen], splitdata[1][datalen+1:];
}
