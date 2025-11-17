package service

import "strings"

func solveImageExtension(extension string) bool {
	e := strings.ToLower(extension)

	switch e {
	case ".jpeg":
		return true
	case ".jpg":
		return true

	default:
		return false
	}
}
