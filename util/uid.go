package util

import "fmt"

func FormatUID(uid uint64) string {
	return fmt.Sprintf("%x", uid)
}
