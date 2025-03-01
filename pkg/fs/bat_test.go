package fs

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestBATSizes(t *testing.T) {
	var h BATHeader
	require.Equal(t, 32, int(unsafe.Sizeof(h)))
}
