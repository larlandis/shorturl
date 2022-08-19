package hash

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

type (
	Hash struct {
		storage Storage
	}
	Storage interface {
		SavePair(ctx context.Context, input string, short string) error
		Search(ctx context.Context, short string) (url string, err error)
	}
)

// CreateNewHash creates and saves a new short hash from a given string
func (s Hash) CreateNewHash(ctx context.Context, input string, length uint) (string, error) {
	hash := s.hash(input, length)
	// save pair short/input
	err := s.storage.SavePair(ctx, input, hash)
	fmt.Println(err)
	return hash, err
}

// GetFromHash searches and returns a saved string from a hash
func (s Hash) GetFromHash(ctx context.Context, hash string) (string, error) {
	url, err := s.storage.Search(ctx, hash)
	return url, err
}

// hash creates md5 hash from input string then
// encode hash to base64 string and returns first n chars
func (s Hash) hash(input string, n uint) string {
	hash := md5.Sum([]byte(input))
	short := base64.StdEncoding.EncodeToString(hash[:])[0:n]
	return short
}

func New(storage Storage) *Hash {
	return &Hash{storage: storage}
}
