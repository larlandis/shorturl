package storage

import (
	"context"
	"fmt"
	"sync"
)

type LocalStorage struct {
	m sync.Map
}

func (r *LocalStorage) SavePair(ctx context.Context, input string, short string) error {
	r.m.Store(short, input)
	return nil
}

func (r *LocalStorage) Search(ctx context.Context, short string) (url string, err error) {
	val, ok := r.m.Load(short)
	if !ok {
		return "", fmt.Errorf("value not found")
	}
	return val.(string), nil
}

func NewLocal() *LocalStorage {
	return &LocalStorage{m: sync.Map{}}
}
