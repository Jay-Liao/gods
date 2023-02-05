package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	s := NewSet()
	assert.Equal(t, 0, s.Len())
}

func TestSet_Len(t *testing.T) {
	s := NewSet()
	assert.Equal(t, 0, s.Len())
	s.Add(1)
	assert.Equal(t, 1, s.Len())
	s.Remove(1)
	assert.Equal(t, 0, s.Len())
}

func TestSet_Add(t *testing.T) {
	s := NewSet()
	s.Add(1)
	assert.Equal(t, 1, s.Len())
	assert.True(t, s.Contains(1))
}

func TestSet_Remove(t *testing.T) {
	s := NewSet()
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

func TestSet_Contains(t *testing.T) {
	s := NewSet()
	assert.False(t, s.Contains(1))
	s.Add(1)
	assert.True(t, s.Contains(1))
}
