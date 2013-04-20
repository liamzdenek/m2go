package m2go;

import "fmt";
import "bytes";

type Response struct {
    Headers []Header;
    Body string;
    StatusCode string; // "200"
    Status string;     // "OK"
}

func (response Response) String() string {
    var buffer bytes.Buffer;

    if len(response.StatusCode) == 0 {
        response.StatusCode = "200";
    }

    if len(response.Status) == 0 {
        response.Status = "OK";
    }

    buffer.WriteString(fmt.Sprintf("HTTP/1.0 %s %s\r\n", response.StatusCode, response.Status));

    for n, _ := range response.Headers {
        buffer.WriteString(fmt.Sprintf("%s: %s\r\n", response.Headers[n].key, response.Headers[n].value));
    }

    buffer.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n%s",len(response.Body),response.Body));

    return buffer.String();
    // "HTTP/1.0 200 OK\r\nContent-Length: 3\r\n\r\nlol"
}