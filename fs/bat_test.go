package fs

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestBATSizes(t *testing.T) {
	var h BATHeader
	require.Equal(t, uintptr(32), unsafe.Sizeof(h))
}
