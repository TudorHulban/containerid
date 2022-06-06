package linuxid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMachineID(t *testing.T) {
	id, errID := readMachineID()
	require.NoError(t, errID)
	require.Equal(t, []byte{182, 123, 169}, id)
}

func TestRandIntID(t *testing.T) {
	_, errID := randInt()
	require.NoError(t, errID)
}
