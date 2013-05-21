package m2go;

import "fmt";
import "bytes";

type Response struct {
    Request *Request;
    Headers map[string][]string;
    Body string;
    StatusCode int; // "200"
    Status string;  // "OK"
    ContentType string;
}

func NewResponse(req *Request) *Response {
    return &Response{Request: req, StatusCode:200, Status:"OK", Headers: make(map[string][]string)};
}

// shortcut for `request.Reply(response.String());`
func (rsp Response) Dispatch() {
    s := rsp.String();
    rsp.Request.Reply(s);
}

func (res *Response) AddHeader(key, value string) {
    res.Headers[key] = append(res.Headers[key], value);
}

func (res *Response) String() string {
    var buffer bytes.Buffer;

    res.Request.SaveGroups(res);

    buffer.WriteString(fmt.Sprintf("HTTP/1.0 %d %s\r\n", res.StatusCode, res.Status));

    for k, values := range res.Headers {
        for _,v := range values {
            buffer.WriteString(fmt.Sprintf("%s: %s\r\n", k, v));
        }
    }

    if len(res.ContentType) != 0 {
        buffer.WriteString(fmt.Sprintf("Content-Type: %s\r\n", res.ContentType));
    }

    buffer.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n%s",len(res.Body),res.Body));

    return buffer.String();
    // "HTTP/1.0 200 OK\r\nContent-Length: 3\r\n\r\nlol"
}
