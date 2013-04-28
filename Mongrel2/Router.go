package m2go;

import "fmt";

type Router struct {
     Routes []Route;
     NotFound func(*Request);
};

func (r *Router) Handle(req *Request) {
    for i,_ := range r.Routes {
        if r.Routes[i].Path.MatchString(req.Path) {
            matches := r.Routes[i].Path.FindAllStringSubmatch(req.Path,-1);
            req.URLArgs = matches;
            r.Routes[i].Handler(req);
            return;
        }
    }
    if r.NotFound != nil {
        r.NotFound(req);
    }
}

func (r *Router) AddRoute(route Route) {
    r.Routes = append(r.Routes, route);
}
