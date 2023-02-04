package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIntSet(t *testing.T) {
	s := NewIntSet()
	assert.Equal(t, 0, s.Len())
}

func TestIntSet_Len(t *testing.T) {
	s := NewIntSet()
	assert.Equal(t, 0, s.Len())
	s.Add(1)
	assert.Equal(t, 1, s.Len())
	s.Remove(1)
	assert.Equal(t, 0, s.Len())
}

func TestIntSet_Add(t *testing.T) {
	s := NewIntSet()
	s.Add(1)
	assert.Equal(t, 1, s.Len())
	assert.True(t, s.Contains(1))
}

func TestIntSet_Remove(t *testing.T) {
	s := NewIntSet()
	assert.Equal(t, 0, s.Len())
	s.Remove(1)
	assert.Equal(t, 0, s.Len())
	s.Add(1)
	assert.Equal(t, 1, s.Len())
	assert.True(t, s.Contains(1))
	s.Remove(2)
	assert.Equal(t, 1, s.Len())
	assert.True(t, s.Contains(1))
	s.Remove(1)
	assert.Equal(t, 0, s.Len())
	assert.False(t, s.Contains(1))
}

func TestIntSet_Contains(t *testing.T) {
	s := NewIntSet()
	assert.False(t, s.Contains(1))
	s.Add(1)
	assert.True(t, s.Contains(1))
}
