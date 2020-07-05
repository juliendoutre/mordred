package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBinary(t *testing.T) {
	t.Run("with binary encoded string", func(t *testing.T) {
		assert.True(t, IsBinary("001101101010101010101111100"))
	})

	t.Run("with non binary encoded string", func(t *testing.T) {
		assert.False(t, IsBinary("001101101010101010101111100a"))
	})

	t.Run("with empty string", func(t *testing.T) {
		assert.False(t, IsBinary(""))
	})
}

func TestIsDecimal(t *testing.T) {
	t.Run("with decimal encoded string", func(t *testing.T) {
		assert.True(t, IsDecimal("57462356873"))
	})

	t.Run("with non binary encoded string", func(t *testing.T) {
		assert.False(t, IsDecimal("57462356873wdf"))
	})

	t.Run("with empty string", func(t *testing.T) {
		assert.False(t, IsDecimal(""))
	})
}

func TestIsHex(t *testing.T) {
	t.Run("with hex encoded string", func(t *testing.T) {
		assert.True(t, IsHex("5746235687fae"))
	})

	t.Run("with non hex encoded string", func(t *testing.T) {
		assert.False(t, IsHex("57462356873wdf"))
	})

	t.Run("with empty string", func(t *testing.T) {
		assert.False(t, IsHex(""))
	})
}

func TestIsBase64(t *testing.T) {
	t.Run("with base64 encoded string", func(t *testing.T) {
		assert.True(t, IsBase64("5746235687fae+/GDCUFHKJCLfjwhsv327849"))
	})

	t.Run("with non base64 encoded string", func(t *testing.T) {
		assert.False(t, IsBase64("!5746235687fae+/GDCUFHKJCLfjwhsv327849"))
	})

	t.Run("with empty string", func(t *testing.T) {
		assert.False(t, IsBase64(""))
	})
}

func TestIsPrintable(t *testing.T) {
	t.Run("with base64 encoded string", func(t *testing.T) {
		assert.True(t, IsPrintable("5746235687fae+/GDCUFHKJCLfjwhsv327849"))
	})

	t.Run("with non base64 encoded string", func(t *testing.T) {
		assert.False(t, IsPrintable("!5746235687fae+/GDCUFHKJCLfjwhsv327849\u0001"))
	})

	t.Run("with empty string", func(t *testing.T) {
		assert.False(t, IsPrintable(""))
	})
}
