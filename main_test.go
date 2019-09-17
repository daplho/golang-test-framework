package main

import (
	"fmt"
	"testing"

	"./storage"
)

type fakeGetValuer struct{}

func (f *fakeGetValuer) GetValue(key string) string {
	return "Fake Storage"
}

func TestDoWork(t *testing.T) {
	s := storage.NewStorageProvider(&fakeGetValuer{})
	fmt.Printf("%s\n", s.Gv.GetValue("fakekey"))
}
