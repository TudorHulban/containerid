package linuxid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewID(t *testing.T) {
	id, errNew := NewID()
	require.NoError(t, errNew)

	fmt.Println("ID:", id)
}
