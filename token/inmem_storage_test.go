package token_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/vcraescu/go-oblio-api/token"
	"testing"
	"time"
)

func TestInMemStorage_Set(t *testing.T) {
	t.Parallel()

	t.Run("token expired", func(t *testing.T) {
		t.Parallel()

		var (
			storage = token.NewInMemStorage()
			ctx     = context.Background()
			want    = "token"
		)

		err := storage.Set(ctx, want, time.Second/4)
		require.NoError(t, err)

		time.Sleep(time.Second / 2)

		got, err := storage.Get(ctx)
		require.Error(t, err)
		require.Empty(t, got)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		var (
			storage = token.NewInMemStorage()
			ctx     = context.Background()
			want    = "token"
		)

		err := storage.Set(ctx, want, time.Hour)
		require.NoError(t, err)

		got, err := storage.Get(ctx)
		require.NoError(t, err)
		require.Equal(t, want, got)
	})
}
