package fs

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestBlockSizes(t *testing.T) {
	require.Equal(t, uintptr(BlockHeaderSize), unsafe.Sizeof(BlockHeader{}))

	require.Equal(t, uintptr(BlockSize), unsafe.Sizeof(Block{}))
}
