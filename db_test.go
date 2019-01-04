package godb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDB(t *testing.T) {
	db, err := NewDB("")
	require.NoError(t, err)
	assert.NotNil(t, db)
}
