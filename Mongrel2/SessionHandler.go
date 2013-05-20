package m2go;

type SessionHandler struct {
    Engines map[int]SessionEngine;
}

func NewSessionHandler() *SessionHandler {
    sh := SessionHandler{};
    sh.Engines = make(map[int]SessionEngine);
    return &sh;
}

func (sh *SessionHandler) AddEngine(id int, engine SessionEngine) {
    sh.Engines[id] = engine;
}
