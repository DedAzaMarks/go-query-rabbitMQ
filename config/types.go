package config

import "github.com/google/uuid"

type Request struct {
	From string
	To   string
	Id   uuid.UUID
}

type Respond struct {
	PathLength int
	Path       []string
	Id         uuid.UUID
}
