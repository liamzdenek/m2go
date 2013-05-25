package m2go;

type Router struct {
     Routes []*Route;
};

func NewRouter() *Router {
    return &Router{};
}

func (r *Router) Handle(req *Request) {
    for _,route := range r.Routes {
        matches,urlargs := route.Test(req.Path);
        if matches {
            req.URLArgs = urlargs;
            route.Handle(req);
            return;
        }
    }
}

func (r *Router) AddRoute(route Route) {
    r.Routes = append(r.Routes, route);
}
