package linuxid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandIntID(t *testing.T) {
	val1, errID := randInt()
	require.NoError(t, errID)

	fmt.Println(val1)

	val2, _ := randInt()
	fmt.Println(val2)
}

func TestGetSequenceID(t *testing.T) {
	var id1, id2, id3 string

	n := 10000
	for i := 1; i <= n; i++ {
		seq := getSequenceID()

		if i == 1 {
			id1 = seq
		}

		if i == 2 {
			id2 = seq
		}

		if i == n {
			id3 = seq
		}
	}

	require.Equal(t, "0000", id1, "id1")
	require.Equal(t, "0001", id2, "id2")
	require.Equal(t, "0000", id3, "id3")
}
