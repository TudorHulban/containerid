package linuxid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewIDRandom(t *testing.T) {
	id, errNew := NewIDRandom()
	require.NoError(t, errNew)

	fmt.Println("ID:", id)
}

func TestNewIDIncremental10K(t *testing.T) {
	id1, errNew := NewIDIncremental10K()
	require.NoError(t, errNew)

	id2, _ := NewIDIncremental10K()
	require.Equal(t, uint(1), id2-id1, fmt.Sprintf("id1: %d, id2: %d", id1, id2))
}
