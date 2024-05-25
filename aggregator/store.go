package main

import "github.com/dr4g0n7ly/AutoTariff-Service/types"

type Storer interface {
	Insert(types.Distance) error
}

type MemoryStore struct {
	data map[int]float64
}

func (m *MemoryStore) Insert(d types.Distance) error {
	m.data[d.OBUID] += d.Distance
	return nil
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}
