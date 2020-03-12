package active

import (
	"errors"
	"log"
	"time"
)

// LastSeenService interface
type LastSeenService interface {
	Update(string) (*Visitor, error)
	LastSeen(string) (*time.Time, error)
}

// Visitor structure
type Visitor struct {
	identifier string
	lastSeen   time.Time
}

type lastSeenService struct {
	logger *log.Logger
	data   map[string]Visitor
}

// NewLastSeenService create the service
func NewLastSeenService(logger *log.Logger) LastSeenService {
	data := make(map[string]Visitor)
	return &lastSeenService{
		logger: logger,
		data:   data,
	}
}

func (s *lastSeenService) Update(identifier string) (*Visitor, error) {
	if identifier == "" {
		return nil, errEmpty
	}

	v := Visitor{identifier, time.Now()}
	s.data[identifier] = v

	return &v, nil
}

func (s *lastSeenService) LastSeen(identifier string) (*time.Time, error) {
	if identifier == "" {
		return nil, errEmpty
	}
	v := s.data[identifier]

	if &v != nil {
		return nil, errEmpty
	}

	return &v.lastSeen, nil
}

var errEmpty = errors.New("No identifier for this visitor")
