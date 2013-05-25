package m2go;

type RouteLiteral struct {
    Path string;
    Handler func(*Request);
};

func NewRouteLiteral(s string, handler func(*Request)) *RouteLiteral {
    return &RouteLiteral{Path: s, Handler: handler};
}

func (route *RouteLiteral) Test(path string) (bool,[][]string) {
    return (path == route.Path),nil;
}

func (route *RouteLiteral) Handle(r *Request) {
    route.Handler(r);
}
