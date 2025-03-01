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
	var sr9 VTOCBlockSR9
	require.Equal(t, 1024, int(unsafe.Sizeof(sr9)))
	require.Equal(t, 0x00, int(unsafe.Offsetof(sr9.NextBlockDAddr)))
	require.Equal(t, 0x04, int(unsafe.Offsetof(sr9.Entries)))

	var sr10 VTOCBlock
	require.Equal(t, 1012, int(unsafe.Sizeof(sr10)))
	require.Equal(t, 0x00, int(unsafe.Offsetof(sr10.NextBlockDAddr)))
	require.Equal(t, 0x04, int(unsafe.Offsetof(sr10.Entries)))
}

func TestVTOCEHeaderLayout(t *testing.T) {
	var sr9 VTOCEHeaderSR9
	require.Equal(t, 64, int(unsafe.Sizeof(sr9)))
	require.Equal(t, 0x00, int(unsafe.Offsetof(sr9.Version)))
	require.Equal(t, 0x01, int(unsafe.Offsetof(sr9.SystemType)))
	require.Equal(t, 0x02, int(unsafe.Offsetof(sr9.Flags)))
	require.Equal(t, 0x04, int(unsafe.Offsetof(sr9.ObjectUID)))
	require.Equal(t, 0x0C, int(unsafe.Offsetof(sr9.ObjectTypeDefUID)))
	require.Equal(t, 0x14, int(unsafe.Offsetof(sr9.ObjectACLUID)))
	require.Equal(t, 0x1C, int(unsafe.Offsetof(sr9.CurrentLength)))
	require.Equal(t, 0x20, int(unsafe.Offsetof(sr9.BlocksUsed)))
	require.Equal(t, 0x24, int(unsafe.Offsetof(sr9.LastUsedTime)))
	require.Equal(t, 0x28, int(unsafe.Offsetof(sr9.LastModifiedTime)))
	require.Equal(t, 0x2C, int(unsafe.Offsetof(sr9.DirectoryUID)))
	require.Equal(t, 0x34, int(unsafe.Offsetof(sr9.MoreStuff)))
	require.Equal(t, 0x38, int(unsafe.Offsetof(sr9.ObjectLockKey)))

	var sr10 VTOCEHeader
	require.Equal(t, 196, int(unsafe.Sizeof(sr10)))
}

func TestVTOCELayout(t *testing.T) {
	var sr9 VTOCESR9
	require.Equal(t, 204, int(unsafe.Sizeof(sr9)))
	require.Equal(t, 0x00, int(unsafe.Offsetof(sr9.Header)))
	require.Equal(t, 0x40, int(unsafe.Offsetof(sr9.FileMap0)))
	require.Equal(t, 0xC0, int(unsafe.Offsetof(sr9.FileMap1Ptr)))
	require.Equal(t, 0xC4, int(unsafe.Offsetof(sr9.FileMap2Ptr)))
	require.Equal(t, 0xC8, int(unsafe.Offsetof(sr9.FileMap3Ptr)))

	var sr10 VTOCE
	require.Equal(t, 336, int(unsafe.Sizeof(sr10)))
	require.Equal(t, 0x00, int(unsafe.Offsetof(sr10.Header)))
	require.Equal(t, 0xc4, int(unsafe.Offsetof(sr10.FileMap0)))
	require.Equal(t, 0xd8, int(unsafe.Offsetof(sr10.FileMap1Ptr)))
	require.Equal(t, 0xcc, int(unsafe.Offsetof(sr10.FileMap2Ptr)))
	require.Equal(t, 0x118, int(unsafe.Offsetof(sr10.FileMap3Ptr)))
}
