package m2go;

type SessionKeyGroup struct {
    Key   string;
    EngineId int;
    Pairs map[string]string;
    delta []string;
}

func NewSessionKeyGroup(key string) *SessionKeyGroup {
    kg := &SessionKeyGroup{};
    kg.Pairs = make(map[string]string);
    kg.Key = key;
    return kg;
}

func (group *SessionKeyGroup) Set(key string, value string) {
    var found bool = false;
    for _, k := range group.delta {
        if k == key {
            found = true;
        }
    }
    if !found {
        group.delta = append(group.delta, key);
    }
    group.Pairs[key] = value;
}

func (group *SessionKeyGroup) Get(key string) (bool,string) {
    for k,_ := range group.Pairs {
        if key == k {
            return true, group.Pairs[k];
        }
    }
    return false, "";
}

// because delta is read-only
func (group *SessionKeyGroup) GetDelta() []string {
    return group.delta;
}
