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

func TestRectanglePerimeter(t *testing.T) {
	t.Run("while length is 10 and breadth is 3 perimeter should be 26", func(t *testing.T) {
		assert.Equal(t, 26.0, NewRectangle(10, 3).Perimeter())
	})

	t.Run("while length is 3 and breadth is 10 perimeter should be 14", func(t *testing.T) {
		assert.Equal(t, 26.0, NewRectangle(3, 10).Perimeter())
	})

	t.Run("while length is 8.2 and breadth is 7.5 perimeter should be 31.4", func(t *testing.T) {
		assert.Equal(t, 31.4, NewRectangle(8.2, 7.5).Perimeter())
	})

	t.Run("while length is 4 and breadth is 4 perimeter should be 16", func(t *testing.T) {
		assert.Equal(t, 16.0, NewRectangle(4, 4).Perimeter())
	})

	t.Run("while length is 5.2 and breadth is 5.2 perimeter should be 20.8", func(t *testing.T) {
		assert.Equal(t, 20.8, NewRectangle(5.2, 5.2).Perimeter())
	})
}

func TestRectangleArea(t *testing.T) {
	t.Run("while length is 1 breadth is 1 area should be 1", func(t *testing.T) {
		assert.Equal(t, 1.0, NewRectangle(1, 1).Area())
	})

	t.Run("while length is 1 breadth is 2 area should be 2", func(t *testing.T) {
		assert.Equal(t, 2.0, NewRectangle(1, 2).Area())
	})

	t.Run("while length is 1 breadth is 1.2 area should be 1.2", func(t *testing.T) {
		assert.Equal(t, 1.2, NewRectangle(1, 1.2).Area())
	})

	t.Run("while length is 2 breadth is 1 area should be 2", func(t *testing.T) {
		assert.Equal(t, 2.0, NewRectangle(2, 1).Area())
	})

	t.Run("while length is 1.5 breadth is 1 area should be 1.5", func(t *testing.T) {
		assert.Equal(t, 1.5, NewRectangle(1.5, 1).Area())
	})

	t.Run("while length is 1.5 breadth is 1.5 area should be 2.25", func(t *testing.T) {
		assert.Equal(t, 2.25, NewRectangle(1.5, 1.5).Area())
	})

}
