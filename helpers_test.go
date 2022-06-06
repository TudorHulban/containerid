package linuxid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandIntID(t *testing.T) {
	_, errID := randInt()
	require.NoError(t, errID)
}
