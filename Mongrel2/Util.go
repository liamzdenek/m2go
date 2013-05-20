package m2go;

import "strings";
import "strconv";

func Util_parse_str(s string) map[string][]string {
    r := make(map[string][]string);
   
    if len(s) == 0 {
        return r;
    }

    pairs := strings.SplitN(s,"&",-1);

    for _,val := range pairs {
        keypair := strings.SplitN(val,"=",2);
        r[keypair[0]] = append(r[keypair[0]], keypair[1]);
    }

    return r;
}

func Util_parse_netstring(ns string) (string, string) {
    // length, rest
    splitdata := strings.SplitN(ns, ":", 2);
    datalen,_ := strconv.Atoi(splitdata[0]);
    return splitdata[1][:datalen], splitdata[1][datalen+1:];
}
