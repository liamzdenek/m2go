package main;

import "fmt";
import "bytes";

type Response struct {
    headers []Header;
    body string;
    statuscode string; // "200"
    status string;     // "OK"
}

func (response Response) String() string {
    var buffer bytes.Buffer;

    if len(response.statuscode) == 0 {
        response.statuscode = "200";
    }

    if len(response.status) == 0 {
        response.status = "OK";
    }

    buffer.WriteString(fmt.Sprintf("HTTP/1.0 %s %s\r\n", response.statuscode, response.status));

    for n, _ := range response.headers {
        buffer.WriteString(fmt.Sprintf("%s: %s\r\n", response.headers[n].key, response.headers[n].value));
    }

    buffer.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n%s",len(response.body),response.body));

    return buffer.String();
    // "HTTP/1.0 200 OK\r\nContent-Length: 3\r\n\r\nlol"
}
