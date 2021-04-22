package main

import "errors"

//UserSessionData contains user session data
type UserSessionData struct {
	UID int
}

//Session contains session data
type Session struct {
	SessionID string
	UserData  UserSessionData
}

//SessionsStore contains user sessions. map key - session ID
type SessionsStore map[string]Session

//IsSessionAvailable returns true if there is a session with the given ID
func (store *SessionsStore) IsSessionAvailable(sessionID string) bool {
	_, exists := (*store)[sessionID]
	return exists
}

//GetSessionByID gets a copy of the Session at the specified ID
func (store *SessionsStore) GetSessionByID(sessionID string) (Session, error) {
	if session, ok := (*store)[sessionID]; ok {
		return session, nil
	}
	return Session{}, errors.New("session not found")
}
