package main

import (
	"./storage"
	"fmt"
)

type realStorageProvider struct {}

func (r *realStorageProvider) GetValue(key string) string {
	return fmt.Sprintf("Real Storage: returning values for key: %s", key)
}

func doWork() {
	s := storage.NewStorageProvider(&realStorageProvider{})
	fmt.Printf("%s\n", s.Gv.GetValue("realkey"))
}

func main() {
	doWork()
}
