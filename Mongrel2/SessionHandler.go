package m2go;

type SessionHandler struct {
    Engines map[int]SessionEngine;
    LoadedGroups map[string]*SessionKeyGroup;
}

func NewSessionHandler() *SessionHandler {
    sh := SessionHandler{};
    sh.LoadedGroups = make(map[string]*SessionKeyGroup);
    sh.Engines = make(map[int]SessionEngine);
    return &sh;
}

func (sh *SessionHandler) AddEngine(id int, engine SessionEngine) {
    sh.Engines[id] = engine;
}

func (sh *SessionHandler) GetGroup(request *Request, key string, engineid int) (bool,*SessionKeyGroup) {
    engine := sh.Engines[engineid];

    err, group := engine.Load(request, key);
    if err {
        return err, nil;
    } else {
        sh.LoadedGroups[key] = group;
        group.EngineId = engineid;
        return err, group;
    }
}

func (sh *SessionHandler) SaveGroups(r *Response) {
    for key,group := range sh.LoadedGroups {
        engine := sh.Engines[group.EngineId];
        engine.Save(r, group, key);
    }
}
