package stack_test

import (
	"github.com/MickStanciu/go-fn/stack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	s := stack.NewStack[string]()
	require.NotNil(t, s)
	assert.Equal(t, 0, s.GetLength())
}

func TestPush(t *testing.T) {
	s := stack.NewStack[string]()
	require.NotNil(t, s)
	assert.Equal(t, 0, s.GetLength())

	s.Push("hello")
	s.Push("world")
	assert.Equal(t, 2, s.GetLength())
}

func TestPop(t *testing.T) {
	s := stack.NewStack[string]()
	require.NotNil(t, s)

	s.Push("hello")
	s.Push("world")
	assert.Equal(t, 2, s.GetLength())

	itemWorld, err := s.Pop()
	require.NoError(t, err)
	assert.Equal(t, "world", itemWorld)
	assert.Equal(t, 1, s.GetLength())

	itemHello, err := s.Pop()
	require.NoError(t, err)
	assert.Equal(t, "hello", itemHello)
	assert.Equal(t, 0, s.GetLength())
}
