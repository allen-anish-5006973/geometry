package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRectangle(t *testing.T) {
	t.Run("while initializing a rectangle, a new rectangle must be returned", func(t *testing.T) {
		assert.IsType(t, Rectangle{}, NewRectangle())
	})
}
