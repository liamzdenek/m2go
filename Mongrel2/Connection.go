package m2go

import "strings";
import zmq "github.com/alecthomas/gozmq";
import "regexp";

type Connection struct
{
    Ctx *zmq.Context;
    Req *zmq.Socket;
    Rsp *zmq.Socket;

    SessionHandler *SessionHandler;
    Router *Router;
    SenderId string;
}

func NewConnection(Router *Router, sh *SessionHandler, SenderId string, req_addr string, rsp_addr string) *Connection {
    Ctx, _ := zmq.NewContext();
    Req, _ := Ctx.NewSocket(zmq.PULL);
    Rsp, _ := Ctx.NewSocket(zmq.PUB);
    Req.Connect(req_addr);
    Rsp.Connect(rsp_addr);
    //req.SetSockOptInt(zmq.RCVTIMEO, 1000);
    Rsp.SetSockOptString(zmq.IDENTITY, SenderId);

    return &Connection{
        Ctx:Ctx,
        Req:Req,
        Rsp:Rsp,
        Router:Router,
        SessionHandler:sh,
        SenderId:SenderId,
    };
}

func (conn *Connection) StartServer() {
    for {
        msg, err := conn.Req.Recv(0);
        if err == nil {
            conn.Router.Handle(conn.Parse(string(msg)));
        }
    }
}

func (conn *Connection) Parse(msg string) *Request {
    //sender, conn_id, path, rest
    splitdata := strings.SplitN(msg, " ", 4);
    headers, rest := Util_parse_netstring(splitdata[3]);
    body, _ := Util_parse_netstring(rest);

    // we're not going to break out a JSON parser. DIY. For speed.
    headers = headers[1:len(headers)-1];
    regex,_ := regexp.Compile(`"(.*)"\:"(.*)"`);

    // precalculate header array size. this might accidentally
    // overshoot in size, but whatever.
    headerary := make([]Header,strings.Count(headers,"\":\"")-1);

    var parts, headstring []string;
    var headercount int;
    for {
        headstring = strings.SplitN(headers,",", 2);
        if len(headstring) == 1 {
            break;
        }

        headers = headstring[1];
        // this could be replaced with sscanf. it should perform faster
        parts = regex.FindStringSubmatch(headstring[0]);
        if len(parts) == 3 {
            headerary[headercount] = Header{key:string(parts[1]),value:string(parts[2])};
            headercount++;
        }
    }

    return NewRequest(*conn.SessionHandler, splitdata[0], splitdata[1], splitdata[2], body, conn, headerary);
}


