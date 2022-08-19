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
	data := "data"
	short, err := s.CreateNewHash(ctx, data, 8)
	if err != nil {
		t.Errorf(err.Error())
	}
	if short != "jXd-OF09" {
		t.Errorf("want: %s, got: %s", "jXd-OF09", short)
	}
	input, err := s.GetFromHash(ctx, short)
	if err != nil {
		t.Errorf(err.Error())
	}
	if input != data {
		t.Errorf("want: %s, got: %s", data, input)
	}
}
