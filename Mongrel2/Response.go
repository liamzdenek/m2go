package m2go;

import "fmt";
import "bytes";

type Response struct {
    Request *Request;
    Headers []Header;
    Body string;
    StatusCode int; // "200"
    Status string;  // "OK"
    ContentType string;
}

// shortcut for `request.Reply(response.String());`
func (rsp Response) Dispatch() {
    s := rsp.String();
    rsp.Request.Reply(s);
}

func (res *Response) AddHeader(key, value string) {
    res.Headers = append(res.Headers, Header{key:key, value:value});
}

func (res *Response) String() string {
    var buffer bytes.Buffer;

    res.Request.SaveGroups(res);

    buffer.WriteString(fmt.Sprintf("HTTP/1.0 %d %s\r\n", res.StatusCode, res.Status));

    for n, _ := range res.Headers {
        buffer.WriteString(fmt.Sprintf("%s: %s\r\n", res.Headers[n].key, res.Headers[n].value));
    }

    if len(res.ContentType) != 0 {
        buffer.WriteString(fmt.Sprintf("Content-Type: %s\r\n", res.ContentType));
    }

    buffer.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n%s",len(res.Body),res.Body));

    return buffer.String();
    // "HTTP/1.0 200 OK\r\nContent-Length: 3\r\n\r\nlol"
}
