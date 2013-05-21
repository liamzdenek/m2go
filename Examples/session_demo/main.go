package main

import "fmt";
import "regexp";
import "../../Mongrel2/";

var ENGINE_CLIENT_UNSECURE = 0;

func main() {
    sh := m2go.NewSessionHandler();

    se := m2go.SessionEngineClientUnsecure{};

    sh.AddEngine(ENGINE_CLIENT_UNSECURE,&se);

    r := m2go.NewRouter();

    r.AddRoute(m2go.Route{Path:regexp.MustCompile(`^/$`),Handler:ActionIndex});
    r.AddRoute(m2go.Route{Path:regexp.MustCompile(`^/login`),Handler:ActionLogin});
    r.AddRoute(m2go.Route{Path:regexp.MustCompile(`^/logout`),Handler:ActionLogout});
    r.NotFound = ErrorNotFound;

    conn := *m2go.NewConnection(r,sh,"82209006-86FF-4982-B5EA-D1E29E55D481", "tcp://127.0.0.1:9997", "tcp://127.0.0.1:9996");
    conn.StartServer();
}

func GetLoginForm(prefix string) string {
    return fmt.Sprintf(
        "%s%s", 
        prefix, 
        "<form action=\"/login\" method=\"POST\">"+
            "Username: <input name=\"username\"><br/>"+
            "Password <input name=\"password\"><br/>"+
            "<input value=\"Login\" type=\"submit\">"+
        "</form>",
    );
}

func ActionIndex(r *m2go.Request) {
    rsp := r.GetResponse();
    
    err,kg := r.GetGroup("sess", ENGINE_CLIENT_UNSECURE);

    if !err && len(kg.Get("username")) > 0 {
        rsp.Body = fmt.Sprintf("<p>You are logged in as %s. <a href=\"/logout\">Logout</a></p>", kg.Get("username") );
        rsp.Dispatch();
    } else {
        rsp.Body = GetLoginForm("");
        rsp.Dispatch();
    }
}

func ActionLogin(r *m2go.Request) {
    request := m2go.Util_parse_str(r.Body);

    rsp := r.GetResponse();
    if request["username"] == nil || request["password"] == nil {
        rsp.Body = "You must POST a username and a password to /login";
        rsp.StatusCode = 400;
        rsp.Status = "Bad Request";
        rsp.Dispatch();
        return;
    }

    username := request["username"][0];
    password := request["password"][0];

    if password == "password" && len(username) > 0 {
        rsp.Body = "<p>Successful Login</p>";
        err, group := r.GetGroup("sess", ENGINE_CLIENT_UNSECURE);
        
        if err {
            rsp.Body = "There was an internal error logging you in";
            rsp.StatusCode = 500;
            rsp.Status = "Internal Error";
            rsp.Dispatch();
            return;
        }

        fmt.Printf("Group: %v\n", group);
        
        group.Set("username", username);
        fmt.Printf("Dispatch\n");
        rsp.Dispatch();
    } else {
        rsp.Body = GetLoginForm("<p>Invalid username or password</p>");
        rsp.Dispatch();
    }

    fmt.Printf("Username: %v, Password: %v\n", username, password); 
}

func ActionLogout(r *m2go.Request) {
    res := r.GetResponse();

    _, group := r.GetGroup("sess", ENGINE_CLIENT_UNSECURE);
    group.Set("username", "");
    
    res.Body = GetLoginForm("");
    res.Dispatch();
}

func ErrorNotFound(r *m2go.Request) {
    res := r.GetResponse();
    res.Body = "The document you are looking for cannot be found\n";
    res.ContentType = "text/plain";
    res.StatusCode = 404;
    res.Status = "Not Found";
    res.Dispatch();
}
