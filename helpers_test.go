package linuxid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandIntID(t *testing.T) {
	val, errID := randInt("xxx")
	require.NoError(t, errID)

	fmt.Println(val)
}
