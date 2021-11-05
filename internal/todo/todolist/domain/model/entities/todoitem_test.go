package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDoItem(t *testing.T) {
	todo := NewToDoItem()
	assert.NotNil(t, todo)
}
