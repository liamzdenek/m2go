package m2go;

import "regexp";

type Route struct {
    Path   *regexp.Regexp;
    Handler func(*Request);
};
