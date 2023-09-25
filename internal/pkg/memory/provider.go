package memory

import (
	"time"

	"github.com/so-brian/cache-service/internal/pkg/utility"
)

type memory interface {
	IsExpired() bool
}

type keyValueMemory struct {
	key    string
	value  string
	expire *int64
}

func (memory *keyValueMemory) IsExpired() bool {
	now := utility.GetNowUnix()
	if memory.expire == nil || *memory.expire > now {
		return false
	}

	return true
}

type MemoryProvider interface {
}

type KeyValueMemoryProvider struct {
	memory []keyValueMemory
}

func NewKeyValueMemoryProvider() *KeyValueMemoryProvider {
	return &KeyValueMemoryProvider{memory: make([]keyValueMemory, 0)}
}

func (provider *KeyValueMemoryProvider) Set(key string, value string, expire *time.Time) {
	var duration *int64
	if expire == nil {
		duration = nil
	} else {
		durationValue := expire.Unix()
		duration = &durationValue
	}

	provider.memory = append(provider.memory, keyValueMemory{key: key, value: value, expire: duration})
}

func (provider *KeyValueMemoryProvider) Get(key string) (string, bool) {
	for _, item := range provider.memory {
		if item.key == key {
			now := utility.GetNowUnix()
			if item.expire == nil || *item.expire > now {
				return item.value, true
			}
		}
	}

	return "", false
}
