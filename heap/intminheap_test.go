package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntMinHeap(t *testing.T) {
	h := NewIntMinHeap()
	assert.Equal(t, 0, h.Len())
}

func TestIntMinHeap_PushVal(t *testing.T) {
	h := NewIntMinHeap()
	h.PushVal(1)
	assert.Equal(t, 1, h.Len())
	h.PushVal(0)
	assert.Equal(t, 2, h.Len())
}

func TestIntMinHeap_PopMin(t *testing.T) {
	h := NewIntMinHeap()
	h.PushVal(1)
	assert.Equal(t, 1, h.Len())
	h.PushVal(0)
	assert.Equal(t, 2, h.Len())
	assert.Equal(t, 0, h.PopMin())
	assert.Equal(t, 1, h.Len())
	assert.Equal(t, 1, h.PopMin())
	assert.Equal(t, 0, h.Len())
}
