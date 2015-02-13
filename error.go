package gocouchbaseio

import (
	"fmt"
)

type generalError struct {
	message string
}

func (e generalError) Error() string {
	return e.message
}

type memdError struct {
	code StatusCode
}

func (e memdError) Error() string {
	switch e.code {
	case StatusSuccess:
		return "Success."
	case StatusKeyNotFound:
		return "Key not found."
	case StatusKeyExists:
		return "Key already exists."
	case StatusTooBig:
		return "Document value was too large."
	case StatusNotStored:
		return "The document could not be stored."
	case StatusBadDelta:
		return "An invalid delta was passed."
	case StatusNotMyVBucket:
		return "Operation sent to incorrect server."
	case StatusUnknownCommand:
		return "An unknown command was received."
	case StatusOutOfMemory:
		return "The server is out of memory."
	case StatusTmpFail:
		return "A temporary failure occurred.  Try again later."
	default:
		return fmt.Sprintf("An unknown error occurred (%d).", e.code)
	}
}
func (e memdError) KeyNotFound() bool {
	return e.code == StatusKeyNotFound
}
func (e memdError) KeyExists() bool {
	return e.code == StatusKeyExists
}
func (e memdError) Temporary() bool {
	return e.code == StatusOutOfMemory || e.code == StatusTmpFail
}

type agentError struct {
	message string
}

func (e agentError) Error() string {
	return e.message
}
