package main

import (
	"errors"
	"log"
	"time"
)

// LastSeenService interface
type LastSeenService interface {
	Update(string) (*Visitor, error)
	LastSeen(string) (time.Time, error)
}

// Visitor structure
type Visitor struct {
	identifier string
	lastSeen   time.Time
}

type lastSeenService struct {
	logger *log.Logger
}

// NewLastSeenService create the service
func NewLastSeenService(logger *log.Logger) LastSeenService {
	return &lastSeenService{logger}
}

func (s *lastSeenService) Update(identifier string) (*Visitor, error) {
	if identifier == "" {
		return nil, errEmpty
	}

	//lookup the visitor

	//update

	//return
	return &Visitor{identifier, time.Now()}, nil
}

func (s *lastSeenService) LastSeen(identifier string) (time.Time, error) {
	if identifier == "" {
		return time.Now(), errEmpty
	}

	s.logger.Println("Last seen")
	return time.Now(), nil
}

var errEmpty = errors.New("No identifier for this visitor")
