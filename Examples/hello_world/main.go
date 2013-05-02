package main

import "fmt";
import "regexp";
import "bytes";
import "../../Mongrel2/";

func main() {
    r := m2go.Router{};

    r.AddRoute(m2go.Route{Path:regexp.MustCompile(`^/$`),Handler:SayHello});
    r.AddRoute(m2go.Route{Path:regexp.MustCompile(`^/([[:alpha:]]*)$`),Handler:SayHelloWithName});
    r.NotFound = ErrorNotFound;

    conn := *m2go.NewConnection(r,"82209006-86FF-4982-B5EA-D1E29E55D481", "tcp://127.0.0.1:9997", "tcp://127.0.0.1:9996");
    conn.StartServer();
}

func SayHello(r *m2go.Request) {
    response := m2go.Response{};
    response.Body = "Hello, World!";
    r.Reply(response.String());
}

func SayHelloWithName(r *m2go.Request) {
    var buffer bytes.Buffer;

    buffer.WriteString(fmt.Sprintf("Hello, %s!", r.URLArgs[0][1]));

    response := m2go.Response{};
    response.Body = buffer.String();
    response.ContentType = "text/plain";

    r.Reply(response.String());
}

func ErrorNotFound(r *m2go.Request) {
    response := m2go.Response{};
    response.Body = "The document you are looking for cannot be found\n";
    response.ContentType = "text/plain";
    response.StatusCode = 404;
    response.Status = "Not Found";
    r.Reply(response.String());
}
