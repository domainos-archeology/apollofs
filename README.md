# apollofs

A tool for interacting with Apollo computer filesystem images.

Still early days, with only the `info` command really receiving much attention.  Trying to decipher filesystem documentation in both:

* http://www.bitsavers.org/pdf/apollo/AEGIS_Internals_and_Data_Structures_Jan86.pdf
* http://www.bitsavers.org/pdf/apollo/002398-04_Domain_Engineering_Handbook_Rev4_Jan87.pdf

Usage:
```
% ./apollofs
A tool for interacting with Apollo filesystems

Usage:
  apollofs [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  copyin      Copy a file from the host to the disk image
  copyout     Copy a file from the disk image to the host
  cpboot      Make the disk image bootable (by copying sysboot)
  help        Help about any command
  info        Dump information about filesystem structures
  invol       Initialize a disk image
  list        List files/directories in the disk image (similar to 'ls' in the host)

Flags:
  -d, --diskImage string   Path to disk image (required)
  -h, --help               help for apollofs

Use "apollofs [command] --help" for more information about a command.
```

### dumping labels (physical and logical)
```
% ./apollofs -d dn3500_sr10.4.awd info labels
Disk image: ../dn3500_sr10.4.awd
PV Label:
Version: 1
Name: APOLLO
UID: 10012345.776175af
DriveType: 1284
TotalBlocksInVolume: 329400
BlocksPerTrack: 18
TracksPerCylinder: 15
SectorStart: 5
SectorSize: 2260
PreComp: 5
Logical Volumes:
  LV0: block 1 / 329399 (alt)
...
LV Label:
Version: 1
Name:
ID: 20012345.776175d5
Label Written: Sun, 20 Nov 2044 18:44:46 UTC
Boot Time: Sun, 20 Nov 2044 16:03:05 UTC
Dismounted Time: Sun, 20 Nov 2044 16:03:23 UTC
VTOCHeader:
  Version: 2
  VTOCSizeInBlocks: 3623
  VTOCBlocksUsed: 907
  NetworkRootDirVTOCX: vtoc blk 160110, index 2
  DiskEntryDirVTOCX: vtoc blk 160110, index 1
  OSPagingFileVTOCX: vtoc blk 160111, index 1
  SysbootVTOCX: vtoc blk 160109, index 0
...
%
```

### dumping blocks (referenced by their physical daddr, not relative to a logical volume)
```
% ./apollofs -d ../dn3500_sr10.4.awd info block 160111
Disk image: ../dn3500_sr10.4.awd
BlockHeader:
  object uid: 0000.0202 (vtoc_$uid)
  Page within object: 160110
  Block type: 0
  System type: 0
  Checksum: 0
  Block DAddr: 160111
%
```
