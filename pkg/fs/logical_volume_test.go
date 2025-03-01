package fs

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestLVLabelLayout(t *testing.T) {
	// TODO handle lv label version 0

	var l LVLabel
	require.Equal(t, 204, int(unsafe.Sizeof(l)))
}
