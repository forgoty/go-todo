package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToDoItem(t *testing.T) {
	todo := NewToDoItem()
	assert.NotNil(t, todo)
}
