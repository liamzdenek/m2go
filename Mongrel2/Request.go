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
    return &Response{ Request: req };
}

func (req *Request) GetGroup(key string, engineid int) (bool,*SessionKeyGroup) {
    sh := req.Conn.SessionHandler;
    return sh.GetGroup(req, key, engineid);
}
