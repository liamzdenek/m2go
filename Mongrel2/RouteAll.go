package m2go;

type RouteAll struct {
    Handler func(*Request);
};

func NewRouteAll (handler func(*Request)) *RouteAll {
    return &RouteAll{Handler: handler};
}

func (route *RouteAll) Test(path string) (bool,[][]string) {
    return true,nil;
}

func (route *RouteAll) Handle(r *Request) {
    route.Handler(r);
}
