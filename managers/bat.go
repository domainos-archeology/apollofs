package managers

import "github.com/domainos-archeology/apollofs/fs"

type BATManager struct {
	lvol *fs.LogicalVolume
}

func NewBATManager(lvol *fs.LogicalVolume) *BATManager {
	return &BATManager{
		lvol: lvol,
	}
}

func (bm *BATManager) AllocateBlock() (fs.DAddr, error) {
	return 0, nil
}
