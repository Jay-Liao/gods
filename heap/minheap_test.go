package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	h := NewMinHeap()
	assert.Equal(t, 0, h.Len())
}

func TestMinHeap_PushVal(t *testing.T) {
	h := NewMinHeap()
	h.PushVal(1)
	assert.Equal(t, 1, h.Len())
	h.PushVal(0)
	assert.Equal(t, 2, h.Len())
}

func TestMinHeap_PopMin(t *testing.T) {
	h := NewMinHeap()
	h.PushVal(1)
	assert.Equal(t, 1, h.Len())
	h.PushVal(0)
	assert.Equal(t, 2, h.Len())
	assert.Equal(t, 0, h.PopMin())
	assert.Equal(t, 1, h.Len())
	assert.Equal(t, 1, h.PopMin())
	assert.Equal(t, 0, h.Len())
}
