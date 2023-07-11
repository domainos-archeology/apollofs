package util

import "path"

func SplitPath(p string) []string {
	var parts []string

	for {
		dir, file := path.Split(p)
		if file != "" {
			parts = append([]string{file}, parts...)
		}

		dir = path.Clean(dir) // remove the trailing /

		if dir == "/" {
			break
		}

		p = dir
	}

	return parts
}
