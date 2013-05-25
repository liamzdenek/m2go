package m2go;

import "regexp";

type RouteRegexp struct {
    Regexp *regexp.Regexp;
    Handler func(*Request);
};

func NewRouteRegexp(r *regexp.Regexp, handler func(*Request)) *RouteRegexp {
    return &RouteRegexp{Regexp:r, Handler:handler};
}

func (route *RouteRegexp) Test(path string) (bool,[][]string) {
    if route.Regexp.MatchString(path){
        return true, route.Regexp.FindAllStringSubmatch(path,-1);
    } else {
        return false, nil;
    }
}

func (route *RouteRegexp) Handle(r *Request) {
    route.Handler(r);
}
