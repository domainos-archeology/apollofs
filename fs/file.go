package fs

import (
	"github.com/domainos-archeology/apollofs/uid"
)

type File struct {
}

type FileManager struct {
	// Aegis Internals talks about the file manager using the containing
	// directory UID to locate the volume to create the file, but we've only got
	// one volume.
	lvol *LogicalVolume
}

func (fm *FileManager) CreateFile(containingDirectory uid.UID) (*File, error) {
	// The file manager (file_$create) uses the directory UID to determine on
	// which volume to create the file, then creates the object in the following
	// steps:

	// 1. Calls the AST manager (ast_$get_info) to locate the volume that holds
	//    the directory UID.

	// 2. Calls the UID generation routine (uid_$gen) to create a UID for the
	//    object.
	//TODO uid.Generate(fm.lvol.LastUsedTime, fm.lvol.NextNodeID)

	// 3. Fills the VTOCE header for the new rue with default object attributes
	//    and calls the VTOC manager to allocate a VTOCE.

	// 4. The VToe manager (vtoc_'allocate) takes the hashed U1D, locates the
	//    appropriate VTOCE block, and checks for a free VTOCE. If it locates one,
	//    it sets the -in use- field in the VTOCE header and appends the header to
	//    the VTOCE.

	// 5. If the VTOCEs are full, it calls the BAT manager (bat_'allocate) to
	//    allocate an extension block and chains it to the VTOe block. Once it
	//    builds the VTOCE, the routine creates a VTOCX for the object and returns
	//    it to file _ $create, who writes it to the active segment table entry for
	//    the object being created. In turn, file _ 'create returns the generated
	//    UID to its caller.

	return nil, nil
}

func (fm *FileManager) Delete(u uid.UID) error {
	return nil
}
