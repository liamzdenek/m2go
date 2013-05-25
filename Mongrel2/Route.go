package m2go;

type Route interface {
    Test(string) (bool,[][]string);
    Handle(*Request);
}
