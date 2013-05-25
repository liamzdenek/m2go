package main

import "fmt";
import "bytes";
import "regexp";
import "../../Mongrel2/";

func main() {
    r := m2go.Router{};

    r.AddRoute(m2go.NewRouteLiteral("/", SayHello));
    r.AddRoute(m2go.NewRouteRegexp(regexp.MustCompile(`^/([[:alpha:]]*)$`),SayHelloWithName));
    r.AddRoute(m2go.NewRouteAll(ErrorNotFound));

    conn := *m2go.NewConnection(&r,nil,"82209006-86FF-4982-B5EA-D1E29E55D481", "tcp://127.0.0.1:9997", "tcp://127.0.0.1:9996");
    conn.StartServer();
}

func SayHello(r *m2go.Request) {
    response := r.GetResponse();
    response.Body = "Hello, World!";

    response.Dispatch();
}

func SayHelloWithName(r *m2go.Request) {
    var buffer bytes.Buffer;

    buffer.WriteString(fmt.Sprintf("Hello, %s!", r.URLArgs[0][1]));

    response := r.GetResponse();
    response.Body = buffer.String();
    response.ContentType = "text/plain";

    response.Dispatch();
}

func ErrorNotFound(r *m2go.Request) {
    response := r.GetResponse();
    response.Body = "The document you are looking for cannot be found\n";
    response.ContentType = "text/plain";
    response.StatusCode = 404;
    response.Status = "Not Found";

    response.Dispatch();
}
