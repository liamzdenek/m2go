package m2go;

import "fmt";

type Request struct {
    SenderId string;
    ConnId string;
    Path string;
    Body string;
    Conn *Connection;
    Headers []Header;
    URLArgs [][]string;

    SessionHandler *SessionHandler;
    LoadedGroups map[string]*SessionKeyGroup;
}

func NewRequest(sh *SessionHandler, SenderId, ConnId, Path, Body string, Conn *Connection, Headers []Header) *Request {
    return &Request{
        SenderId: SenderId,
        ConnId: ConnId,
        Path: Path,
        Body: Body,
        Conn: Conn,
        Headers: Headers,

        SessionHandler: sh,
        LoadedGroups: make(map[string]*SessionKeyGroup),
    };
}

func (req *Request) Reply(msg string) {
    rsp  := req.Conn.Rsp;
    response := fmt.Sprintf(
        "%s %d:%s, %s",
        req.SenderId,
        len(req.ConnId),
        req.ConnId,
        msg,
    );
    //fmt.Printf("Response: %s\n", response);
    rsp.Send([]byte(response),0);
}

func (req *Request) GetResponse() *Response {
    return &Response{ Request: req, StatusCode:200, Status:"OK" };
}

func (req *Request) GetGroup(key string, engineid int) (bool,*SessionKeyGroup) {
    engine := req.SessionHandler.Engines[engineid];

    err, group := engine.Load(req, key);
    if err {
        return err, nil;
    } else {
        req.LoadedGroups[key] = group;
        group.EngineId = engineid;
        return err, group;
    }
}

func (req *Request) SaveGroups(res *Response) {
    for key,group := range req.LoadedGroups {
        engine := req.SessionHandler.Engines[group.EngineId];
        engine.Save(res, group, key);
    }
}
