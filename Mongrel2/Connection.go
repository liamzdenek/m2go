package m2go

import "strings";
import zmq "github.com/alecthomas/gozmq";
import "strconv";
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
        SenderId: splitdata[0],
        ConnId:   splitdata[1],
        Path:      splitdata[2],
        Body:      body,
        Conn:      conn,
        Headers:   headerary,
    };
}

func (con *Connection) parse_netstring(ns string) (string, string) {
    // length, rest
    splitdata := strings.SplitN(ns, ":", 2);
    datalen,_ := strconv.Atoi(splitdata[0]);
    return splitdata[1][:datalen], splitdata[1][datalen+1:];
}
