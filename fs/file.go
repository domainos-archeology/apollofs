package fs

type File struct {
}

type FileManager struct {
	// Aegis Internals talks about the file manager using the containing
	// directory UID to locate the volume to create the file, but we've only got
	// one volume.
	lvol *LogicalVolume
}

func (fm *FileManager) CreateFile(containingDirectory UID) (*File, error) {
	return nil, nil
}
