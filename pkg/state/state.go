package state

import "errors"

type State struct {
	Documents map[string]string
}

func NewState() *State {
	return &State{
		Documents: make(map[string]string),
	}
}
func (s *State) OpenDocument(uri string, content string) {
	s.Documents[uri] = content
}

func (s *State) UpdateDocument(uri string, content string)(error) {
	if _, ok := s.Documents[uri]; !ok {
		return errors.New("document not found")
	}
	s.Documents[uri] = content
	return nil
}