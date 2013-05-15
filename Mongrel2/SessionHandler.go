package m2go;

type SessionHandler struct {
    Engines map[int]SessionEngine;
    NeedSaving []SessionKeyGroup;
}

func NewSessionHandler() *SessionHandler {
    sh := SessionHandler{};
    sh.Engines = make(map[int]SessionEngine);
    return &sh;
}

func (handler *SessionHandler) AddEngine(id int, engine SessionEngine) {
    handler.Engines[id] = engine;
}

func (handler *SessionHandler) GetGroup(request *Request, key string, engineid int) (bool,*SessionKeyGroup) {
    engine := handler.Engines[engineid];

    err, group := engine.Load(request, key);
    if(!err)
    {
        sh.NeedsSaving = append(sh.NeedsSaving, group);
    }
    return err, group;
}
