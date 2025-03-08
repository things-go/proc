package password

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SCrypt(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		org := "hahaha"
		cpt := NewScrypt(0)

		dst, err := cpt.GenerateFromPassword(org)
		require.Nil(t, err)
		require.NoError(t, cpt.CompareHashAndPassword(dst, org))
	})

	t.Run("not correct", func(t *testing.T) {
		org := "hahaha"
		cpt := NewScrypt(0)

		dst, err := cpt.GenerateFromPassword(org)
		require.Nil(t, err)
		require.Error(t, cpt.CompareHashAndPassword(dst, "invalid"))
	})
}

func Benchmark_SCrypt_GenerateFromPassword(b *testing.B) {
	cpt := NewScrypt(0)

	for i := 0; i < b.N; i++ {
		_, _ = cpt.GenerateFromPassword("hahaha")
	}
}

func Benchmark_SCrypt_CompareHashAndPassword(b *testing.B) {
	org := "hahaha"
	cpt := NewScrypt(0)
	dst, _ := cpt.GenerateFromPassword(org)

	for i := 0; i < b.N; i++ {
		_ = cpt.CompareHashAndPassword(dst, org)
	}
}
