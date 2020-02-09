package main

import (
	"errors"
	"time"
)

// LastSeenService interface
type LastSeenService interface {
	Update(string) (Visitor, error)
	LastSeen(string) (time.Time, error)
}

// Visitor structure
type Visitor struct {
	identifier string
	lastSeen   time.Time
}

type lastSeenService struct {
}

func (lastSeenService) Update(identifier string) (Visitor, error) {
	if identifier == "" {
		return nilVisitor, errEmpty
	}

	//lookup the visitor

	//update

	//return
	return Visitor{identifier, time.Now()}, nil
}

func (lastSeenService) LastSeen(identifier string) (time.Time, error) {
	if identifier == "" {
		return time.Now(), errEmpty
	}
	return time.Now(), nil
}

var errEmpty = errors.New("No identifier for this visitor")
var nilVisitor = Visitor{"", time.Now()}