package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRectangle(t *testing.T) {
	t.Run("while initializing a rectangle, a new rectangle must be returned", func(t *testing.T) {
		assert.IsType(t, Rectangle{}, NewRectangle(5, 6))
	})

	t.Run("while initializing a rectangle length and breadth must be present", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewRectangle(6, 4)
		})

	})

	t.Run("while initializing a rectangle length and breadth must be positive", func(t *testing.T) {
		assert.Panics(t, func() {
			NewRectangle(0, 4)
		})

	})
}
