package kv

import (
	"context"
	"sync"
)

type KvClient struct {
	storage map[string]interface{}
	mu      *sync.RWMutex
}

func NewKvClient(ctx context.Context) (*KvClient, error) {
	return &KvClient{
		storage: make(map[string]interface{}),
		mu:      &sync.RWMutex{},
	}, nil
}

func (kv *KvClient) Get(key string) (interface{}, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	v, ok := kv.storage[key]

	if !ok {
		return "", false
	}

	return v, true
}

func (kv *KvClient) Set(key string, val interface{}) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	kv.storage[key] = val
}

func (kv *KvClient) Delete(key string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	delete(kv.storage, key)
}
