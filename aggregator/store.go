package main

import "github.com/dr4g0n7ly/AutoTariff-Service/types"

type MemoryStore struct{}

func (m *MemoryStore) Insert(d types.Distance) error {
	return nil
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}
