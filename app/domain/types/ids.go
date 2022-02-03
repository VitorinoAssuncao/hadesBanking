package types

import "github.com/google/uuid"

type ExternalID string

type InternalID int

type ContextKey string

func (e ExternalID) ToUUID() uuid.UUID {
	parsed, _ := uuid.Parse(string(e))
	return parsed
}
