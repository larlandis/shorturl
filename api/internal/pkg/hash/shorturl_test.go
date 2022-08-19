package hash_test

import (
	"context"
	"testing"

	"github.com/larlandis/shorturl/internal/pkg/hash"
	"github.com/larlandis/shorturl/internal/platform/storage"
)

var (
	s   = hash.New(storage.NewLocal())
	ctx = context.Background()
)

func TestHash_CreateAndGetHash(t *testing.T) {
	data := "input_data"
	short, err := s.CreateNewHash(ctx, data, 8)
	if err != nil {
		t.Errorf(err.Error())
	}
	if short != "0d5DD6rv" {
		t.Errorf("want: %s, got: %s", "0d5DD6rv", short)
	}
	input, err := s.GetFromHash(ctx, short)
	if err != nil {
		t.Errorf(err.Error())
	}
	if input != data {
		t.Errorf("want: %s, got: %s", data, input)
	}
}
