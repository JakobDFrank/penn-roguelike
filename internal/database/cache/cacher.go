package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"
)

// Cacher implements caching functionality
type Cacher interface {
	GetBytes(ctx context.Context, key string) ([]byte, error)                               // GetBytes will check the cache for a given key with a []byte value and return it if it exists
	SetBytes(ctx context.Context, key string, value []byte, expiration time.Duration) error // SetBytes will set the cache with the given key, []byte value, and expiration. An expiration of 0 will not expire.
}

// GetStruct will check the cache for a given key with a T value and return it if it exists
func GetStruct[T any](c Cacher, ctx context.Context, key string) (T, error) {
	val, err := c.GetBytes(ctx, key)

	var lvl T

	if err != nil {
		return lvl, err
	}

	bufVal := bytes.NewBuffer(val)
	decoder := gob.NewDecoder(bufVal)

	if err := decoder.Decode(&lvl); err != nil {
		return lvl, err
	}

	return lvl, nil
}

// SetStruct will set the cache with the given key, T value, and expiration. An expiration of 0 will not expire.
func SetStruct[T any](c Cacher, ctx context.Context, key string, value T, expiration time.Duration) error {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	if err := encoder.Encode(value); err != nil {
		return err
	}

	bufBytes := buf.Bytes()
	return c.SetBytes(ctx, key, bufBytes, expiration)
}
