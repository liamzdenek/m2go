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

func (response Response) String() string {
    var buffer bytes.Buffer;

    if response.StatusCode == 0 {
        response.StatusCode = 200;
    }

    if len(response.Status) == 0 {
        response.Status = "OK";
    }

    if nil != response.Request.Conn.SessionHandler {
        response.Request.Conn.SessionHandler.SaveGroups(&response);
    }

    buffer.WriteString(fmt.Sprintf("HTTP/1.0 %d %s\r\n", response.StatusCode, response.Status));

    for n, _ := range response.Headers {
        buffer.WriteString(fmt.Sprintf("%s: %s\r\n", response.Headers[n].key, response.Headers[n].value));
    }

    if len(response.ContentType) != 0 {
        buffer.WriteString(fmt.Sprintf("Content-Type: %s\r\n", response.ContentType));
    }

    buffer.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n%s",len(response.Body),response.Body));

    return buffer.String();
    // "HTTP/1.0 200 OK\r\nContent-Length: 3\r\n\r\nlol"
}
