package m2go;

import "fmt";

type SessionEngine interface {
    Load(*Request,string) (bool,*SessionKeyGroup);
    Save(*Response,*SessionKeyGroup,string) bool;
}

type SessionEngineClientUnsecure struct {}

func (se *SessionEngineClientUnsecure) Load(req *Request, key string) (bool, *SessionKeyGroup) {
    fmt.Printf("%s\n", "Load Client Unsecure");
    kg := NewSessionKeyGroup(key);
    return true, kg;
}

func (se *SessionEngineClientUnsecure) Save(res *Response, group *SessionKeyGroup, key string) bool {
    fmt.Printf("%s\n", "Save Client Unsecure");
    return true;
}
