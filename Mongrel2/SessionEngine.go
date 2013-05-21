package m2go;

import "fmt";
import "strings";
import "github.com/likexian/simplejson";

type SessionEngine interface {
    Load(*Request,string) (bool,*SessionKeyGroup);
    Save(*Response,*SessionKeyGroup,string) bool;
}

type SessionEngineClientUnsecure struct {}

func (se *SessionEngineClientUnsecure) Load(req *Request, key string) (bool, *SessionKeyGroup) {
    fmt.Printf("%s\n", "Load Client Unsecure");

    var cookie string;
    for _,value := range req.Headers["cookie"] {
        parts := strings.SplitN(value, "=", 2);
        if parts[0] == key {
            cookie = parts[1];
            break;
        }
    }

    kg := NewSessionKeyGroup(key);
    if len(cookie) == 0 {
        return false, kg;
    }

    cookie = strings.Replace(cookie, "\\\"", "\"", -1);
    j,err := simplejson.Loads(cookie);

    if err != nil {
        return false, kg;
    } else {
        data,err := j.Map();
        if err == nil {
            for k,v := range data {
                kg.Pairs[k] = v.(string);
            }
        }
        return false, kg;
    }
}

func (se *SessionEngineClientUnsecure) Save(res *Response, group *SessionKeyGroup, key string) bool {
    if len(group.GetDelta()) == 0 {
        return false;
    }
    j,_ := simplejson.Loads("{}"); // there's probably a way to do this without the json parser
    for k,v := range group.Pairs {
        j.Set(k,v);
    }
    data,_ := simplejson.Dumps(j);
    res.AddHeader("Set-Cookie", fmt.Sprintf("%s=%s", key, data));
    return false;
}
