package managers

import (
	"errors"
	"fmt"
	"path"

	"github.com/sirupsen/logrus"

	"github.com/domainos-archeology/apollofs/fs"
	"github.com/domainos-archeology/apollofs/uid"
	"github.com/domainos-archeology/apollofs/util"
)

// if we ever make this a persistent thing, we can add a cache, but for now we
// do every lookup anew.

var errNotFound = errors.New("not found")

type NamingManager struct {
	lvol *fs.LogicalVolume
	file *FileManager
	vtoc *VTOCManager
}

func NewNamingManager(
	lvol *fs.LogicalVolume,
	file *FileManager,
	vtoc *VTOCManager,
) *NamingManager {
	return &NamingManager{
		lvol,
		file,
		vtoc,
	}
}

func (nm *NamingManager) getDiskEntryDirUID() (uid.UID, error) {
	// fmt.Println("getting disk entry dir uid", nm.lvol.Label.VTOCHeader.DiskEntryDirVTOCX)
	entryDirVTOCE, err := nm.vtoc.GetEntry(nm.lvol.Label.VTOCHeader.DiskEntryDirVTOCX)
	if err != nil {
		return uid.Empty, err
	}

	return entryDirVTOCE.Header.ObjectUID, nil
}

func (nm *NamingManager) GetDirEntryUID(dirUID uid.UID, name string) (uid.UID, error) {
	// fmt.Println("looking up", name, "in", dirUID)

	dirVTOCE, err := nm.vtoc.GetEntryForUID(dirUID)
	if err != nil {
		return uid.Empty, err
	}

	if !dirVTOCE.IsDirectory() {
		return uid.Empty, errors.New("not a directory")
	}

	// read the Dir from the first block (will there be more?  I don't think so?)
	dirDAddr := dirVTOCE.FileMap0[0]

	block, err := nm.lvol.ReadBlock(dirDAddr)
	if err != nil {
		return uid.Empty, err
	}

	var dir fs.Dir
	err = block.ReadInto(&dir)
	if err != nil {
		return uid.Empty, err
	}

	// check the linear list first
	for _, entry := range dir.Entries {
		if entry.HasUID() && entry.Name == name {
			return entry.UID, nil
		}
	}

	return uid.Empty, errNotFound
}

func (nm *NamingManager) Resolve(p string) (uid.UID, error) {
	logrus.WithField("path", p).Debug("NamingManager.Resolve")
	if !path.IsAbs(p) {
		return uid.Empty, fmt.Errorf("path must be absolute")
	}

	parts := util.SplitPath(p)
	// loop over parts, looking up the directory for each part. The first lookup is relative to '/'
	curUID, err := nm.getDiskEntryDirUID()
	if err != nil {
		return uid.Empty, err
	}

	for i := 0; i < len(parts); i++ {
		logrus.Debugf("resolving %s contained within %v\n", parts[i], curUID)

		nextUID, err := nm.GetDirEntryUID(curUID, parts[i])
		if err != nil {
			return uid.Empty, err
		}
		curUID = nextUID
	}

	return curUID, nil
}

func (nm *NamingManager) CreateFile(p string) (uid.UID, error) {
	if !path.IsAbs(p) {
		return uid.Empty, fmt.Errorf("file path must be absolute")
	}

	dir, _ /*lastPath*/ := path.Split(p)
	dirUID, err := nm.Resolve(dir)
	if err != nil {
		return uid.Empty, err
	}

	// create a file named lastPath catalogued in dirUID
	u, err := nm.file.CreateFile(dirUID)
	if err != nil {
		return uid.Empty, err
	}

	return u, nil
}

func (nm *NamingManager) CreateDirectory(p string) (*fs.Dir, error) {
	if !path.IsAbs(p) {
		return nil, fmt.Errorf("directory path must be absolute")
	}

	dir, _ /*lastPath*/ := path.Split(p)
	_ /*dirUID*/, err := nm.Resolve(dir)
	if err != nil {
		return nil, err
	}

	// create a directory named lastPath catalogued in dirUID

	return nil, nil
}
