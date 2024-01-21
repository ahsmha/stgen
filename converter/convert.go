package converter

import (
	"path/filepath"
)

func Convert(path string) error {
	ext := filepath.Ext(path)

	if ext == ".md" || ext == ".mkd" {
		return convertMD(path)
	} else if ext == ".html" || ext == ".xml" {
		return convertHTML(path)
	} else {
		return convertRaw(path)
	}
}

func convertHTML(path string) error {
	return nil
}

func convertMD(path string) error {
	return nil
}

func convertRaw(path string) error {
	return nil
}
