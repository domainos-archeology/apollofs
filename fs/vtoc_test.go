package fs

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestVTOCHeaderLayout(t *testing.T) {
	var h VTOCHeader
	require.Equal(t, 100, int(unsafe.Sizeof(h)))
	require.Equal(t, 0x00, int(unsafe.Offsetof(h.Version)))
	require.Equal(t, 0x02, int(unsafe.Offsetof(h.VTOCSizeInBlocks)))
	require.Equal(t, 0x04, int(unsafe.Offsetof(h.VTOCBlocksUsed)))
	require.Equal(t, 0x08, int(unsafe.Offsetof(h.NetworkRootDirVTOCX)))
	require.Equal(t, 0x0c, int(unsafe.Offsetof(h.DiskEntryDirVTOCX)))
	require.Equal(t, 0x10, int(unsafe.Offsetof(h.OSPagingFileVTOCX)))
	require.Equal(t, 0x14, int(unsafe.Offsetof(h.SysbootVTOCX)))
	require.Equal(t, 0x18, int(unsafe.Offsetof(h.VTOCMapData)))
	require.Equal(t, 0x48, int(unsafe.Offsetof(h.Unused)))
}

func TestVTOCBlockLayout(t *testing.T) {
	var b VTOCBlock
	require.Equal(t, 1024, int(unsafe.Sizeof(b)))
	require.Equal(t, 0x00, int(unsafe.Offsetof(b.NextBlockDAddr)))
	require.Equal(t, 0x04, int(unsafe.Offsetof(b.Entries)))
}

func TestVTOCEHeaderLayout(t *testing.T) {
	var h VTOCEHeader
	require.Equal(t, 64, int(unsafe.Sizeof(h)))
	require.Equal(t, 0x00, int(unsafe.Offsetof(h.Info)))
	require.Equal(t, 0x04, int(unsafe.Offsetof(h.ObjectUID)))
	require.Equal(t, 0x0C, int(unsafe.Offsetof(h.ObjectTypeDefUID)))
	require.Equal(t, 0x14, int(unsafe.Offsetof(h.ObjectACLUID)))
	require.Equal(t, 0x1C, int(unsafe.Offsetof(h.CurrentLength)))
	require.Equal(t, 0x20, int(unsafe.Offsetof(h.BlocksUsed)))
	require.Equal(t, 0x24, int(unsafe.Offsetof(h.LastUsedTime)))
	require.Equal(t, 0x28, int(unsafe.Offsetof(h.LastModifiedTime)))
	require.Equal(t, 0x2C, int(unsafe.Offsetof(h.DirectoryUID)))
	require.Equal(t, 0x34, int(unsafe.Offsetof(h.MoreStuff)))
	require.Equal(t, 0x38, int(unsafe.Offsetof(h.ObjectLockKey)))
}

func TestVTOCELayout(t *testing.T) {
	var v VTOCE
	require.Equal(t, 204, int(unsafe.Sizeof(v)))
	require.Equal(t, 0x00, int(unsafe.Offsetof(v.Header)))
	require.Equal(t, 0x40, int(unsafe.Offsetof(v.FileMap0)))
	require.Equal(t, 0xC0, int(unsafe.Offsetof(v.FileMap1Ptr)))
	require.Equal(t, 0xC4, int(unsafe.Offsetof(v.FileMap2Ptr)))
	require.Equal(t, 0xC8, int(unsafe.Offsetof(v.FileMap3Ptr)))
}
