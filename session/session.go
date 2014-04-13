package session

import (
    "container/list"
    "time"
)

var pder = &Provider{list: list.New(), sessions: make(map[string]*list.Element, 0)}

type SessionStore struct {
    sid          string
    timeAccessed time.Time
    value        map[interface{}]interface{}
}

type Session interface {
    Set(key, value interface{}) error
    Get(key interface{}) interface{}
    Delete(key interface{}) error
    SessionID() string
}

func (st *SessionStore) Set(key, value interface{}) error {
    st.value[key] = value
    pder.SessionUpdate(st.sid)
    return nil
}

func (st *SessionStore) Get(key interface{}) interface{} {
    pder.SessionUpdate(st.sid)
    if v, ok := st.value[key]; ok {
        return v
    } else {
        return nil
    }
    return nil
}

func (st *SessionStore) Delete(key interface{}) error {
    delete(st.value, key)
    pder.SessionUpdate(st.sid)
    return nil
}

func (st *SessionStore) SessionID() string {
    return st.sid
}
