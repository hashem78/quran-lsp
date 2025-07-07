package server

import (
	"encoding/json"
	"maps"
)

type Callback = func(e StateEvent)

type internalState struct {
	OpenDocuments map[string]string `json:"open_documents,omitempty"`
}

func (s *internalState) copy() *internalState {
	openDocuments := make(map[string]string, len(s.OpenDocuments))
	maps.Copy(openDocuments, s.OpenDocuments)
	return &internalState{
		OpenDocuments: openDocuments,
	}
}

type State struct {
	internalState    []*internalState
	events           []StateEvent
	callbacks        []Callback
	latestStateIndex int
}

func (s *State) latestInternalState() *internalState {
	return s.internalState[s.latestStateIndex]
}

func (s *State) appendState(state *internalState) {
	s.internalState = append(s.internalState, state)
	s.latestStateIndex += 1
}

func (s *State) runCallbacks(e StateEvent) {
	s.events = append(s.events, e)
	for _, callback := range s.callbacks {
		callback(e)
	}
}

func (s *State) SetDocument(uri string, text string) {
	stateCopy := s.latestInternalState().copy()
	stateCopy.OpenDocuments[uri] = text
	s.appendState(stateCopy)
	s.runCallbacks(
		DocumentSetEvent{
			Type: DOCUMENT_SET,
			Uri:  uri,
			Text: text,
		})
}

func (s *State) GetDocument(uri string) (string, bool) {
	text, found := s.latestInternalState().OpenDocuments[uri]
	return text, found
}

func (s *State) AddListener(eventType EventType, callback Callback) {
	fn := func(e StateEvent) {
		if eventType == e.GetType() || eventType == ANY {
			callback(e)
		}
	}
	s.callbacks = append(s.callbacks, fn)
}

type StateTrace struct {
	Event StateEvent     `json:"event,omitempty"`
	State *internalState `json:"state,omitempty"`
}

func (s *State) Dump() string {

	trace := []StateTrace{}
	for i := range s.latestStateIndex + 1 {
		trace = append(trace, StateTrace{
			Event: s.events[i],
			State: s.internalState[i],
		})
	}
	marshalled, _ := json.MarshalIndent(trace, "", "  ")
	return string(marshalled)
}

func NewServerState() *State {
	return &State{
		internalState: []*internalState{{
			OpenDocuments: map[string]string{},
		}},
		callbacks: []Callback{},
		events: []StateEvent{
			ServerStartEvent{},
		},
		latestStateIndex: 0,
	}
}
