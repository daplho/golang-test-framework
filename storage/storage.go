package storage

import (
	"log"
)

type GetValuer interface {
	GetValue(key string) string
}

type storageProvider struct {
	Gv GetValuer
}

func NewStorageProvider(g GetValuer) *storageProvider {
	s := &storageProvider{}
	if gv, ok := g.(GetValuer); ok {
		s.Gv = gv
		return s
	}

	log.Fatalf("could not convert to GetValuer interface")
	return nil
}
